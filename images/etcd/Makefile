REGISTRY?=icr.io/ai-services-private
IMAGE=etcd
TAG?=3.6.5

CONTAINER_BUILDER?=podman
build:
	$(CONTAINER_BUILDER) build --build-arg ETCD_VERSION=${TAG} -t $(REGISTRY)/$(IMAGE):$(TAG) .
.PHONY: build

push:
	$(CONTAINER_BUILDER) push $(REGISTRY)/$(IMAGE):$(TAG)
.PHONY: push
