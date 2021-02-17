FROM golang:1.14

WORKDIR "/app"

CMD ["go", "build", "-o", "/build/imgdiff", "cmd/main.go"]
