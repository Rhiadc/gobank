FROM golang:1.19-bullseye AS builder
# Must set WORKDIR to something other than /go
# or go mod download will fail with:

#6 0.367 $GOPATH/goexoit.mod exists but should not
WORKDIR /usr/src/app

# Copy only files required to install dependencies (better layer caching) COPY go.mod go.sum ./
# Use cache mount to speed up install of existing dependencies
RUN --mount-type=cache, target=/go/pkg/mod \
    --mount-type=cache, target=/root/.cache/go-build \
    go mod download

COPY . .
# Add flags to statically link binary
RUN go build \
    -ldflags="-linkmode external -extldflags -static" \
    -tags netgo \
    -o api-golang 

# Separate build from deploy stage
FROM cgr.dev/chainguard/static@sha256:233f4d5d58644b716bab9f82c1984a089d8b2ed2f239e3fd4b47964ac41b41e9

LABEL version="X.Y.Z"
USER nonroot

#Grab the mini-static binary to use as the entry point
COPY --from=krallin/ubuntu-tini@sha256:fe681211bc1e1c55caa631a379189d728b0239a2bb580ed1c59f513dd2c28f05 \
/usr/bin/tini-static ./tini-static


COPY --from=builder /usr/src/app/api-golang .

EXPOSE 8080

ENTRYPOINT ["./tini-static", "--"]
CMD ["./api-golang"]