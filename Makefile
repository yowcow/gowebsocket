BINARY = gowebsocket

.PHONY: $(BINARY) up test clean

all: $(BINARY)
	godep save -t -v ./handler

$(BINARY):
	go build

up:
	go run main.go

test:
	go test -v ./...

clean:
	rm -rf $(BINARY)
