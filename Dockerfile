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
	-o todo-api

FROM scratch

ENV GIN_MODE=release

WORKDIR /

COPY --from=build-production /etc/passwd /etc/passwd

COPY --from=build-production /app/todo-api todo-api

USER nonroot

EXPOSE 8080

CMD ["/todo-api"]
