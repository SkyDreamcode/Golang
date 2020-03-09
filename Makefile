#compiler
GOCMD := go
GORUN := $(GOCMD) run

BUILD := `date +%FT%T%z`
PACKAGES := `go list ./... | grep -v /vendor/`
VETPACKAGES := `go list ./... | grep -v /vendor/ | grep -v /example/`
#source file path  
GOFILES := `find . -name "*.go" -type f -not -path "./vendor/*"`

#generate target file
BINARY := main

default:
	go build -o $(BINARY) -tags=jsoniter

list:
	echo $(PACKAGES)
	echo $(VETPACKAGES)
	echo $(GOFILES)

fmt:
	gofmt -s -w $(GOFILES)

fmt-check:
	diff=$$(gofmt -s -d $(GOFILES));\
	if[ -n "$$diff" ];then\
	  echo "please run 'make fmt' and commit the result:";\
	  echo "$$(diff)";\
	  exit 1;\
	fi;

install:
	govendor sync -v

test:
	go test -cpu=1,2,4 -v -tags integration ./...

vet:
	go vet $(VETPACKAGES)

docker:
	docker build -t 


build:
	$(GOCMD) build -o $(BINARY) -v
	./$(BINARY)

run:
	$(GORUN) $(GOFILES)

clean:
	if [ -f $(BINARY) ];then rm -f ${BINARY}; fi

.PHONY: default fmt fmt-check install test vet docker clean build run


