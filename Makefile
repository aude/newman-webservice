IMAGE   ?= aude/newman-webservice
VERSION ?= latest
TAG     ?= $(IMAGE):$(VERSION)

.PHONY: all
all: docker-build

.PHONY: clean
clean: docker-clean

.PHONY: docker-build
docker-build:
	docker build -t $(TAG) .

.PHONY: docker-run
docker-run:
	docker run --rm -p 8080:8080 $(TAG)

.PHONY: docker-clean
docker-clean:
	docker rmi $(TAG)
