BINARY := $(shell basename `pwd`)
VERSION := $(shell cat VERSION)
COMMIT := $(shell git rev-parse HEAD)
COMMIT_SHORT := $(shell git rev-parse --short=8 HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
SOURCEDIR := .
SOURCES := $(shell find $(SOURCEDIR) -name '*.go' -not -path "./vendor/*")
BUILD_TIME := $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")
LDFLAGS := -ldflags "-X main.version=${VERSION} -X main.commit=${COMMIT} -X main.branch=${BRANCH} -X main.buildTime=${BUILD_TIME}"
NOVENDOR := $(shell go list ./... | grep -v -e /vendor/)
COVERAGE := 0 # TODO: raise this

$(BINARY): $(SOURCES) Makefile go.mod go.sum
	go build ${LDFLAGS} -o ${BINARY} .

.PHONY: simplify
simplify:
	gofmt -s -w ${SOURCES}

.PHONY: clean
clean:
	rm -f coverage.html coverage.out ${BINARY}

.PHONY: lint
lint:
	gometalinter \
		--vendor \
		--vendored-linters \
		--config=metalinter.json \
		--deadline=10m \
		./...

.PHONY: test
test: # runs tests
	go test ${NOVENDOR} -race

.PHONY: coverage
coverage: # generates coverage reports
	echo 'mode: atomic' > coverage.out \
		&& echo ${NOVENDOR} \
		| xargs -n1 -I{} sh -c 'go test -race -covermode=atomic -coverprofile=coverage.tmp {} && tail -n +2 coverage.tmp >> coverage.out'
	rm coverage.tmp
	go tool cover -html=coverage.out -o coverage.html
	open coverage.html

.PHONY: coverage_test
coverage_test: # runs coverage, throws error if under threshold
	for i in ${NOVENDOR}; \
	do \
		result=`go test -race -covermode=atomic $$i` \
		&& echo $$result \
		&& pct=`echo $$result | sed -E 's/.*coverage: ([0-9]+).*/\1/'` \
		&& if [ "$$pct" -lt ${COVERAGE} ]; \
		then \
			echo "Insufficient coverage! (Actual: $$pct%, Required: ${COVERAGE}%)"; \
			exit 1; \
		fi; \
	done

.PHONY: sonar
sonar:
	go mod download

	mkdir -p tmp

	(gometalinter \
		--vendor \
		--vendored-linters \
		--config=metalinter.json \
		--deadline=10m \
		--enable-gc \
		--checkstyle \
		./... > tmp/report.xml) || true

	go get github.com/jstemmer/go-junit-report
	go test -v ./... | $(GOPATH)/bin/go-junit-report > tmp/test.xml

	go get github.com/axw/gocov/...
	go get github.com/AlekSi/gocov-xml

	go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/cover.out {{.ImportPath}}"{{end}}' ./... | grep -v -e /vendor/ | xargs -L 1 sh -c
	go list -f '{{if len .TestGoFiles}}"$(GOPATH)/bin/gocov convert {{.Dir}}/cover.out | $(GOPATH)/bin/gocov-xml > {{.Dir}}/coverage.xml"{{end}}' ./... | grep -v -e /vendor/ | xargs -L 1 sh -c

.PHONY: prep_release
prep_release:
	goimports -w $(SOURCES)
	goreturns -w $(SOURCES)
	gofmt -s -w ${SOURCES}
	go mod tidy

.PHONY: version
version:
	for f in "VERSION README.md sonar-project.properties"; do \
		sed -i '' 's/$(shell echo ${old} | sed 's/\./\\./g')/$(new)/g' $$f; \
	done
