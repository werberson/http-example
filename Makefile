IMAGE := werberson/http-example

image:
	@echo Image: $(IMAGE)
	docker build -f Dockerfile -t $(IMAGE) .

image-release: image
	@echo Image: $(VERSION)
	@echo DOCKER_USERNAME: $(DOCKER_USERNAME)
	@echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker build -f Dockerfile -t $(IMAGE) .
	docker tag $(IMAGE) $(IMAGE):$(VERSION)
	docker push $(IMAGE):$(VERSION)

.PHONY: image image-releases