APP_NAME = go-alert
VM_DIR = ~/Desktop/OracleVM

build:
	go mod download
	go build -o ${APP_NAME} main.go
	./${APP_NAME}

zip:
	tar -zcvf ${VM_DIR}/${APP_NAME}.tar.gz .

all: zip build

clean:
	go clean
ifneq ($(wildcard ${APP_NAME}),)
	rm ${APP_NAME} && rm ${APP_NAME}.* 
	
endif