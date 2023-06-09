GOCMD = go
GOBUILD = $(GOCMD) build
GOMOD = $(GOCMD) mod
GOTEST = $(GOCMD) test
BINARY_NAME = goadmin
CLI = adm

all: serve

init:
	$(GOMOD) init $(module)

install:
	$(GOMOD) tidy

serve:
	$(GOCMD) run .

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME) -v ./
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME)_mac -v ./
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 $(GOBUILD) -o ./build/$(BINARY_NAME).exe -v ./

generate:
	$(CLI) generate -c adm.ini

test: black-box-test user-acceptance-test

black-box-test: ready-for-data
	$(GOTEST) -v -test.run=TestMainBlackBox
	make clean

user-acceptance-test: ready-for-data
	$(GOTEST) -v -test.run=TestMainUserAcceptance
	make clean

ready-for-data:
	cp admin.db admin_test.db

clean:
	rm admin_test.db

.PHONY: all serve build generate test black-box-test user-acceptance-test ready-for-data clean
