FROM golang:1.21 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .
COPY go.sum .
COPY go.mod .

ENV CGO_ENABLED=0

RUN go build -o bin .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/server ./
EXPOSE 8080/tcp
ENTRYPOINT ["./server"]
