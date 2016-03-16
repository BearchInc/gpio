all: deploy

build:
	GOOS=linux GOARCH=arm go build gpio.go

deploy: build
	scp gpio pi@raspberrypi.local:~/
