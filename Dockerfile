FROM golang:1.21 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

# Copy the go source
COPY main.go main.go
COPY proto/ proto/
COPY pkg/ pkg/

# Build server
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$TARGETARCH go build -a -o server main.go


FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/server .
USER 65532:65532

ENTRYPOINT ["/server"]