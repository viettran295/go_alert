APP_NAME = go-alert

build-local:
	go mod download
	go build -o ${APP_NAME} main.go

build-image:
	docker build -t ${APP_NAME} .

run-local: build-local
	./${APP_NAME}

run-image: build-image
	docker run --name ${APP_NAME} -it ${APP_NAME}

all: run-local

clean:
	go clean
ifneq ($(wildcard ${APP_NAME}),)
	rm ${APP_NAME} && rm ${APP_NAME}.* 
	
endif

ifeq ($(shell docker ps -a --format "{{.Names}}"), ${APP_NAME})
	docker stop ${APP_NAME} && docker rm ${APP_NAME}
	docker rmi ${APP_NAME}
endif

save-image:
	docker save -o ${APP_NAME}.tar ${APP_NAME}

test: 
