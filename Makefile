REPO=github.com/sean9999/wgvet
SEMVER := $$(git tag --sort=-version:refname | head -n 1)
BRANCH := $$(git branch --show-current)
REF := $$(git describe --dirty --tags --always)
GOPROXY=proxy.golang.org

info:
	@printf "REPO:\t%s\nSEMVER:\t%s\nBRANCH:\t%s\nREF:\t%s\n" $(REPO) $(SEMVER) $(BRANCH) $(REF)

build:
	mkdir -p bin; go build -o bin/wgvet

install:
	go install .

clean:
	go clean
	rm bin/*

tidy:
	go mod tidy

publish:
	GOPROXY=https://${GOPROXY},direct go list -m ${REPO}@${SEMVER}

.PHONY: build tidy install
