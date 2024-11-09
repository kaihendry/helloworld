STACK = helloworld-sam
export SAM_CLI_TELEMETRY=0

.PHONY: build deploy validate destroy

DOMAINNAME = hellosam.dabase.com
ACMCERTIFICATEARN = arn:aws:acm:eu-west-2:407461997746:certificate/9083a66b-72b6-448d-9bce-6ee2e2e52e36

deploy:
	# Check if stack is in ROLLBACK_COMPLETE state and delete it if true
	aws cloudformation describe-stacks --stack-name $(STACK) --query "Stacks[0].StackStatus" --output text | grep -q "ROLLBACK_COMPLETE" && aws cloudformation delete-stack --stack-name $(STACK) && aws cloudformation wait stack-delete-complete --stack-name $(STACK) || true
	sam build
	sam deploy --no-progressbar --resolve-s3 --stack-name $(STACK) --parameter-overrides DomainName=$(DOMAINNAME) ACMCertificateArn=$(ACMCERTIFICATEARN) --no-confirm-changeset --no-fail-on-empty-changeset --capabilities CAPABILITY_IAM --disable-rollback
	curl -i -H "Host: hellosam.dabase.com" https://$(shell aws cloudformation describe-stacks --stack-name $(STACK) --query "Stacks[0].Outputs[?OutputKey=='CNAME'].OutputValue" --output text)

build-MainFunction:
	CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o ${ARTIFACTS_DIR}/bootstrap

validate:
	aws cloudformation validate-template --template-body file://template.yml

destroy:
	aws cloudformation delete-stack --stack-name $(STACK)

sam-tail-logs:
	sam logs --stack-name $(STACK) --tail

clean:
	rm -rf main gin-bin
