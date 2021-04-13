FROM golang:1.16-alpine AS base
WORKDIR /plants
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

FROM base AS dev
ENV MODE=development
CMD ["go", "run", "main.go"]

FROM base as ci
RUN apk add build-base
RUN go get golang.org/x/tools/cmd/goimports
RUN wget -O- -nv https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

FROM base AS prod-build
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOARCH=amd64 \
    GOOS=linux \
    MODE=production
RUN go build -ldflags="-s -w" -o /bin/plants main.go

FROM alpine:latest AS certs
RUN apk --no-cache --update add ca-certificates

FROM scratch AS prod
COPY --from=prod-build /bin/plants ./
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT [ "/plants" ]
