# build stage
FROM golang:latest AS builder

WORKDIR /go/src/coupon-service

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY cmd ./cmd
COPY pkg ./pkg

RUN CGO_ENABLED=0 go build -o bin/code-review cmd/coupon_service/main.go

FROM scratch
EXPOSE 8080
COPY --from=builder /go/src/coupon-service/bin .
ENTRYPOINT ["./code-review"]

