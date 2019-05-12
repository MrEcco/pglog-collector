DOCKERNAME=mrecco/pglog-collector
DOCKERTAG=v1.0.0
DOCKERFILE=Dockerfile

run:
	@go run main.go

build:
	@go build -o pglog-collector main.go

docker-build:
	@docker build -f $(DOCKERFILE) -t $(DOCKERNAME):$(DOCKERTAG) .

test:
	@echo Haha! Have no tests. :(
	# @go test .
