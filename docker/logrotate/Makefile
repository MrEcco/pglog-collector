IMAGE_NAME=mrecco/logrotate
IMAGE_TAG=v1.0.0
DOCKERFILE=./Dockerfile

build:
	@docker build -f $(DOCKERFILE) -t $(IMAGE_NAME):$(IMAGE_TAG) .

push:
	@docker push $(IMAGE_NAME):$(IMAGE_TAG)

rmi:
	@docker rmi $(IMAGE_NAME):$(IMAGE_TAG)
