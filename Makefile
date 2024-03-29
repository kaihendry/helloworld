STACK = helloworld-sam
VERSION = $(shell git rev-parse --abbrev-ref HEAD)-$(shell git rev-parse --short HEAD)
SAM_CLI_TELEMETRY=0

DOMAINNAME = hellosam.dabase.com # https://ap-southeast-1.console.aws.amazon.com/apigateway/main/publish/domain-names?api=b0p0urf4eb&domain=hellosam.dabase.com&region=ap-southeast-1
ACMCERTIFICATEARN = arn:aws:acm:ap-southeast-1:407461997746:certificate/87b0fd84-fb44-4782-b7eb-d9c7f8714908

deploy:
	sam build
	sam deploy --no-progressbar --resolve-s3 \
	 --stack-name $(STACK) --parameter-overrides DomainName=$(DOMAINNAME) ACMCertificateArn=$(ACMCERTIFICATEARN) Version=$(VERSION) \
	 --no-confirm-changeset --no-fail-on-empty-changeset --capabilities CAPABILITY_IAM

build-Hello:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${ARTIFACTS_DIR}/bootstrap

validate:
	aws cloudformation validate-template --template-body file://template.yml

destroy:
	aws cloudformation delete-stack --stack-name $(STACK)

sam-tail-logs:
	sam logs --stack-name $(STACK) --tail

sync:
	sam sync --watch --stack-name $(STACK)

sam-list-endpoints:
	sam list stack-outputs --stack-name $(STACK)
