FROM golang:latest

WORKDIR /go/src/p02

Copy . .

# EXPOSE 8080
CMD ["go", "run", "main.go"]
