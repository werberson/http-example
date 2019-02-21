IMAGE := werberson/http-example

test:
	true

image:
	docker build -t $(IMAGE) .

image-release:
	echo "$(DOCKER_PASSWORD)" | docker login -u "$(DOCKER_USERNAME)" --password-stdin
	docker tag $(IMAGE) $(IMAGE):latest
#	docker tag $(IMAGE) $(IMAGE):
#	docker push $(USER_NAME)/$(SITE_NAME):latest
#	docker push $(USER_NAME)/$(SITE_NAME):$(SHA)

.PHONY: image push-image test