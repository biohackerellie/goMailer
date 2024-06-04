FROM alpine:latest
WORKDIR "/app"
COPY goMailer /app/goMailer

ENTRYPOINT ["/app/goMailer"]