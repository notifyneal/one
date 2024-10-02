PROJECT := project
SERVICE_NAME := go-starter-svc
STAGE := dev
PROJECT_NUMBER := 12345
PROJECT_ID := $(PROJECT)-$(STAGE)-$(PROJECT_NUMBER)
REGION := us-east5 
REPO_NAME := $(SERVICE_NAME)
IMAGE_NAME := $(REGION)-docker.pkg.dev/$(PROJECT_ID)/$(REPO_NAME)/$(SERVICE_NAME)
ACCOUNT_ID := 123456789
SERVICE_ACCOUNT := $(ACCOUNT_ID)-compute@developer.gserviceaccount.com
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "0.1.0")


.PHONY: all build run clean

all: build run

build:
	cd web; npm install; npm run build
	GOARCH=amd64 GOOS=linux go build
	docker build -t $(IMAGE_NAME):$(VERSION) -t $(IMAGE_NAME):latest .
	docker push $(IMAGE_NAME)

run:
	docker run -it --rm \
		-p 8080:8080 \
		--env-file .env \
		$(IMAGE_NAME)

clean:
	docker rmi -f $(IMAGE_NAME)
