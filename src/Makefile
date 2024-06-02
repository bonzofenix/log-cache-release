DEPLOYMENT_NAME ?= cf

build-syslog-server:
		@echo "# building $*"
		pushd cmd/syslog-server && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ../../build/sylog-server && popd

push-binary-syslog-server:
	$(eval LOG_CACHE_VM := $(shell bosh -d ${DEPLOYMENT_NAME} vms --json | jq -r '.Tables | .[] | .Rows  | .[] | select(.instance|test("log-cache")) | .instance' ))
	@echo "LOG_CACHE VM: $(LOG_CACHE_VM)"
	bosh -d $(DEPLOYMENT_NAME) scp build/syslog-server $(LOG_CACHE_VM):/tmp/.


	bosh -d $(DEPLOYMENT_NAME) ssh $(LOG_CACHE_VM) -c "sudo mv /tmp/syslog-server /var/vcap/jobs/log-cache/syslog-server"
	bosh -d $(DEPLOYMENT_NAME) ssh $(LOG_CACHE_VM) -c "sudo chown vcap:vcap /var/vcap/jobs/log-cache/syslog-server"
	bosh -d $(DEPLOYMENT_NAME) ssh $(LOG_CACHE_VM) -c "sudo chmod +x /var/vcap/jobs/log-cache/syslog-server"
	bosh -d $(DEPLOYMENT_NAME) ssh $(LOG_CACHE_VM) -c "sudo monit restart log-cache-syslog-server" || true

  @echo "Check monit result at:"
	@echo "bosh -d $(DEPLOYMENT_NAME) ssh $(LOG_CACHE_VM)"
