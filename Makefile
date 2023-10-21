ROOT_DIR    = $(shell pwd)
NAMESPACE   = "default"
DEPLOY_NAME = "prometheus-notifier"
DOCKER_NAME = "prometheus-notifier"
PREFIX      = "lingcoder"

include ./hack/hack.mk

.PHONY: build-image

build-image:
	docker build -t $(PREFIX)/$(DOCKER_NAME) .
