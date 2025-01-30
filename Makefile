build:
	mkdir -p bin; go build -o bin/wgvet

install:
	go install .

clean:
	rm bin/*

tidy:
	go mod tidy

.PHONY: build tidy install
