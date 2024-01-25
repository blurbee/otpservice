

server:
	cd server
	GO_POST_PROCESS_FILE="gofmt -w"; openapi-generator generate -g go-gin-server  -p packageName=server -i ../api/openapi.yaml
	cd ..

build:
	go build -o bin/ ...

dep: api util store comms

api util store comms:
	cd $@ && go mod tidy && go mod vendor && cd ..

clean:
	rm -rf bin/
