.SILENT:
.PHONY: run
run:
	cd src && go run .

.SILENT:
.PHONY: build
build:
	cd src && go build -o lg

.PHONY: clean
clean:
	cd src && rm -rf lg
