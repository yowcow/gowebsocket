BINARY = gowebsocket

.PHONY: $(BINARY) up test clean

all: $(BINARY)
	godep save

$(BINARY):
	go build

up:
	go run main.go

test:
	go test -v ./handler

clean:
	rm -rf $(BINARY)
