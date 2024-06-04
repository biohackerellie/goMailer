FROM alpine:latest
WORKDIR "/app"
COPY gomailer /app/gomailer

ENTRYPOINT ["/app/gomailer"]