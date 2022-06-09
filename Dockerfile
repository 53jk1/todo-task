FROM golang:1.18.3 as builder
COPY go.mod go.sum /go/src/github.com/53jk1/todo-task/
WORKDIR /go/src/github.com/53jk1/todo-task/
RUN go mod download
COPY . /go/src/github.com/53jk1/todo-task/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/todo-task github.com/53jk1/todo-task/

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/53jk1/todo-task/build/todo-task /usr/bin/todo-task
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/todo-task"]