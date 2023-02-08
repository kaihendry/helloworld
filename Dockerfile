FROM docker.io/golang as builder

WORKDIR /app

COPY . .

ARG TARGET_ARCH=amd64

RUN echo Building for ${TARGET_ARCH}
RUN go env && go version
RUN CGO_ENABLED=0 GOOS=linux GOARCH=${TARGET_ARCH} \
    go build -o hello

FROM scratch
COPY --from=builder /app/hello /app/
ENV PORT 8080

ARG VERSION=0.0.0

ENV version ${VERSION}
ENTRYPOINT ["/app/hello"]
