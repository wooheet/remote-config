FROM golang:alpine AS builder

ENV GOPROXY=https://proxy.golang.org \
  GO111MODULE=on \
  CGO_ENABLED=0 \
  GOOS=linux \
  GOARCH=amd64

RUN go get github.com/cespare/reflex

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o cmd/metrics-api/main .

FROM scratch

COPY --from=builder /build/main .

ENTRYPOINT ["/main"]