# golang alpine 1.16.3
FROM golang:1.16.6-alpine as builder

ARG commit
ARG version

ENV commit=$commit
ENV version=$version

# SSL for HTTPS calls.
RUN apk update && apk add --no-cache ca-certificates && update-ca-certificates

# Create user.
RUN addgroup -S c && adduser -S numeral -G axiom

WORKDIR $GOPATH/src/github.com/javiertlopez/numeral/
COPY . .

RUN go get github.com/javiertlopez/numeral/cmd/container
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s -X main.commit=${commit} -X main.version=${version}" -o /go/bin/main ./cmd/container

# Small image
FROM scratch

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

# Copy our static executable.
COPY --from=builder /go/bin/main /go/bin/main
# Use an unprivileged user.

USER numeral:axiom

# Port on which the service will be exposed.
EXPOSE 8080

ENTRYPOINT ["/go/bin/main"]