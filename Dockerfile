FROM golang:1.19-buster

WORKDIR /go/src/go-grpc-api-gateway/

COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -o ./out/go-grpc-api-gateway ./cmd/

EXPOSE 50052 50052

CMD ["./out/go-grpc-api-gateway"]