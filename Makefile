.SILENT:
.PHONY: run
run:
	cd src && go run .

.SILENT:
.PHONY: build
build:
	cd src && go build -o main.out

.PHONY: clean
clean:
	cd src && rm -rf *.out
