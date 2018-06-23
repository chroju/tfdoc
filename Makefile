GOCMD=go
GOBUILD=gox
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
BINARY_NAME=tfdoc

all: test build
build:
	$(GOBUILD) -os="linux darwin windows" -arch="386 amd64" -output "bin/tfdoc_{{.OS}}_{{.Arch}}/{{.Dir}}"
test:
	$(GOTEST) -v -covermode=count -coverprofile=coverage.out ./...
clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)
run:
	$(GOBUILD) -o $(BINARY_NAME) -v ./...
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop
