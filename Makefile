-include .env
export

BINARY_NAME ?= slurp
GOOS        ?= freebsd
GOARCH      ?= amd64
RSYNC       := rsync -azvP --delete --include='$(BINARY_NAME)' --include='run.sh' --include='files/***' --exclude='*'

.PHONY: build deploy script restart clean

build:
	GOOS=$(GOOS) GOARCH=$(GOARCH) go build -o $(BINARY_NAME)

deploy: build script
	$(RSYNC) ./ $(SSH_HOST):$(REMOTE_PATH)/
	$(MAKE) restart

script:
	@echo '#!/bin/sh' > run.sh
	@echo 'set -eou pipefail' >> run.sh
	@echo 'exec /home/protected/$(BINARY_NAME) serve' >> run.sh
	@chmod +x run.sh
	@echo 'Created run.sh'

restart:
	ssh $(SSH_HOST) '$(RESTART_CMD)'

clean:
	rm -f $(BINARY_NAME) run.sh