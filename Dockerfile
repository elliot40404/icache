ARG GO_VERSION=1.21.5

FROM golang:${GO_VERSION}

WORKDIR /app

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,source=go.sum,target=go.sum \
    --mount=type=bind,source=go.mod,target=go.mod \
    go mod download -x && \
    apt-get update && \
    apt-get install -y libvips-dev


ARG TARGETARCH 

RUN --mount=type=cache,target=/go/pkg/mod/ \
    --mount=type=bind,target=. \
    GOARCH="$TARGETARCH" GOOS=linux go build -o /bin/server ./cmd/api

ENTRYPOINT ["/bin/server"]
