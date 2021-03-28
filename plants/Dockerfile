FROM golang:1-alpine AS base
WORKDIR /go/src/plants
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

FROM base AS dev
ENV MODE=development
CMD ["go", "run", "main.go"]

FROM base as ci
RUN apk add build-base

FROM base AS prod-build
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOARCH=amd64 \
    GOOS=linux \
    MODE=production
RUN go build -ldflags="-s -w" -o /go/bin/plants main.go

FROM alpine:latest AS certs
RUN apk --no-cache add ca-certificates

FROM scratch AS prod
COPY --from=prod-build /go/bin/plants ./
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT [ "/plants" ]
