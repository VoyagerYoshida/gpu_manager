IMAGE_NAME := go/bot/sample
CONTAINER_NAME := ctn_bot

.PHONY: build
build:
	@docker build . -t $(IMAGE_NAME)

.PHONY: run
run:
	@docker run -d --name $(CONTAINER_NAME) $(IMAGE_NAME) 

.PHONY: stop
stop:
	@docker rm -f $(CONTAINER_NAME)
