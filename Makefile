REPO=$$(git remote -v | head -n 1 | cut -f 2 | cut -d ' ' -f 1)
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

pkgsite:
	if [ -z "$$(command -v pkgsite)" ]; then go install golang.org/x/pkgsite/cmd/pkgsite@latest; fi

docs: pkgsite
	pkgsite -open .

publish:
	GOPROXY=https://${GOPROXY},direct go list -m ${REPO}@${SEMVER}

.PHONY: build tidy install
