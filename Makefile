run: build
	PORT=8080 ./gofeedql

build: 
	go build

dev:
	CompileDaemon -command="make run"
