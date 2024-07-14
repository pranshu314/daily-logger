.SILENT:
.PHONY: run
run:
	go run src/main.go

.SILENT:
.PHONY: build
build:
	cd src && go build -o main.out

.PHONY: clean
clean:
	cd src && rm -rf *.out
