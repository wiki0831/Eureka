
default:
	go build

run:
	./eureka -port 3000 -envfile config/debug.env
