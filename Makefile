IMAGE_NAME := server-go


.PHONY: all build run clean

all: build run

build:
	cd web; npm install; npm run build
	go build
	docker build -t $(IMAGE_NAME) .

run:
	docker run -it --rm \
		-p 8080:8080 \
		--env-file .env \
		$(IMAGE_NAME)

clean:
	docker rmi -f $(IMAGE_NAME)
