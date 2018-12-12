FROM golang:1.11.2 as builder
WORKDIR /go/src/github.com/ynigoreyes/sandbox/api/Calculator
RUN go get -u google.golang.org/grpc
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

FROM alpine:3.6
LABEL maintainer="@ynigoreyes"
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /go/src/github.com/ynigoreyes/sandbox/api/Calculator .
CMD [ "./Calculator" ]

