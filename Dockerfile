FROM golang:1.24-bullseye AS build-base

WORKDIR /app 

COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
	--mount=type=cache,target=/root/.cache/go-build \
	go mod download

FROM build-base AS build-production

RUN useradd -u 1001 nonroot

COPY . .

RUN go build \
	-ldflags="-linkmode external -extldflags -static" \
	-tags netgo \
	-o todo-api ./cmd/todo-api/

FROM golang:1.24-bullseye

ENV GIN_MODE=release

WORKDIR /

COPY --from=build-production /etc/passwd /etc/passwd

COPY --from=build-production /app/todo-api todo-api

RUN mkdir -p /task-data && chown -R 1001:1001 /task-data

COPY ./docs /app/docs

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

USER nonroot

EXPOSE 8080

ENTRYPOINT ["/entrypoint.sh"]
