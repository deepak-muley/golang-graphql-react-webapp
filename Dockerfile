FROM golang:1.14.2 as build

RUN mkdir -p /go/src/golang-graphql-service
WORKDIR /go/src/golang-graphql-service

RUN go env
COPY . .

RUN unset GO111MODULE

RUN go mod download && go mod tidy
RUN go build -o server cmd/main.go
RUN go install -v ./...
RUN ls -ltrah /go/bin/

#-----------------------------
FROM alpine:3.7

COPY --from=build /go/bin/cmd /root/
RUN mv /root/cmd /root/server

EXPOSE 8009
WORKDIR /root/

CMD ["./server"]