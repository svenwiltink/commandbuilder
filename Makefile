all: compile

deb:
	./deb-build.sh

deps:
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

compile:
	go build -o ./build/commandbuilder .

test:
	go test ./...

clean:
	@rm -Rf ./build