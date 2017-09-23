VERSION=`cat VERSION`
BUILD=`git symbolic-ref HEAD 2> /dev/null | cut -b 12-`-`git log --pretty=format:%h -1`
PACKAGES = "./..."

# Setup the -ldflags option for go build here, interpolate the variable
# values
LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

install:
	go install $(LDFLAGS) -v $(PACKAGES)

.PHONY: build
build:
	go build $(LDFLAGS) -v $(PACKAGES)

.PHONY: version
version:
	@echo $(VERSION)-$(BUILD)

.PHONY: clean
clean:
	go clean
	rm -rf coverage-all.out

# Testing

.PHONY: test
test:
	go test -v $(PACKAGES)

.PHONY: cover
cover:
	go test -cover $(PACKAGES)

.PHONY: cover-html
cover-html:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(shell go list ./...),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	rm -rf coverage.out
	go tool cover -html=coverage-all.out

# Lint

lint:
	gometalinter --tests .

# Dependencies

deps:
	go get -u github.com/PuerkitoBio/goquery

dev-deps:
	go get -u github.com/alecthomas/gometalinter
	gometalinter --install

# Documentation

.PHONY: docs
docs: docs-clean
	 cd docs-src && hugo

.PHONY: docs-clean
docs-clean:
	 rm -rf docs/* 

.PHONY: docs-serve
docs-serve:
	cd docs-src && hugo server -D 