IMAGE_NAME := voyagerwy130/gpu-management:1.0
CONTAINER_NAME := ctn_bot

.PHONY: build
build:
	@docker build . -t $(IMAGE_NAME)

.PHONY: run
run:
	@docker run -d --gpus=all --name=$(CONTAINER_NAME) $(IMAGE_NAME) 

.PHONY: stop
stop:
	@docker rm -f $(CONTAINER_NAME)
