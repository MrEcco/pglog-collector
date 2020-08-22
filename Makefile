DOCKERNAME=mrecco/pglog-collector
DOCKERTAG=v1.1.0
DOCKERFILE=./Dockerfile

build:
	@docker build -f $(DOCKERFILE) -t $(DOCKERNAME):$(DOCKERTAG) .

push:
	@docker push $(DOCKERNAME):$(DOCKERTAG)

rmi:
	@docker rmi $(DOCKERNAME):$(DOCKERTAG)
