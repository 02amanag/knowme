dep:
	go get github.com/gin-gonic/gin

build:
	go build

run :	build
	./p-02

builddocker:
	docker build -t resume .
