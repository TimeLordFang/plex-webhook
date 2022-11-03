FROM golang:1.19-alpine AS build-env
RUN apk add ca-certificates upx
ADD . /plex-webhook
RUN cd /plex-webhook \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags "-s -w" -o plex-webhook plex-webhook.go \
    && upx -9 goddns

FROM scratch
COPY --from=build-env /plex-webhook/plex-webhook /
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
ENTRYPOINT ["/plex-webhook"]