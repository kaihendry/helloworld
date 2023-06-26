STACK = helloworld-sam
VERSION = $(shell git rev-parse --abbrev-ref HEAD)-$(shell git rev-parse --short HEAD)
SAM_CLI_TELEMETRY=0

deploy:
	sam build
	sam deploy --no-progressbar --resolve-s3 \
	 --stack-name $(STACK) --parameter-overrides Version=$(VERSION) \
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
