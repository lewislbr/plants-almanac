FROM golang:1.16-alpine AS base
WORKDIR /users
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

FROM base AS dev
ENV MODE=development
CMD ["go", "run", "main.go"]

FROM base as ci
RUN apk add build-base
RUN go get golang.org/x/tools/cmd/goimports
RUN go get honnef.co/go/tools/cmd/staticcheck

FROM base AS prod-build
ENV CGO_ENABLED=0 \
    GO111MODULE=on \
    GOARCH=amd64 \
    GOOS=linux \
    MODE=production
RUN go build -ldflags="-s -w" -o /bin/users main.go

FROM alpine:latest AS certs
RUN apk --no-cache --update add ca-certificates

FROM scratch AS prod
COPY --from=prod-build /bin/users ./
COPY --from=certs /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT [ "/users" ]
