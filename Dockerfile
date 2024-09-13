FROM golang:1.23.1-alpine as base
WORKDIR /root/

RUN apk add git

COPY go.mod go.sum ./
RUN go mod download

COPY cmd cmd
COPY pkg pkg

# Build service in seperate stage.
FROM base as builder
RUN go build -o openslides-projector-service cmd/projectord/main.go


# Test build.
FROM base as testing

RUN apk add build-base

CMD go vet ./... && go test -test.short ./...


# Development build.
FROM base as development

RUN ["go", "install", "github.com/githubnemo/CompileDaemon@latest"]
EXPOSE 9050

COPY entrypoint.sh ./
ENTRYPOINT ["./entrypoint.sh"]
CMD CompileDaemon -log-prefix=false -build="go build -o openslides-projector-service cmd/projectord/main.go" -command="./openslides-projector-service"


# Productive build
FROM alpine:3

LABEL org.opencontainers.image.title="OpenSlides Projector Service"
LABEL org.opencontainers.image.description="The Projector Service is a http endpoint that serves projectors in Openslides."
LABEL org.opencontainers.image.licenses="MIT"
LABEL org.opencontainers.image.source="https://github.com/OpenSlides/openslides-projector-service"

COPY entrypoint.sh ./
COPY --from=builder /root/openslides-projector-service .
EXPOSE 9050
ENTRYPOINT ["./entrypoint.sh"]
CMD exec ./openslides-projector-service
