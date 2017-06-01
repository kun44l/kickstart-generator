FROM alpine:latest

RUN apk update && apk add curl && rm -rf /var/cache/apk/*

ADD kickstart_generator /kickstart_generator
ADD ks.tmpl /ks.tmpl

EXPOSE 8080

HEALTHCHECK --interval=5s --timeout=3s --retries=3 \
      CMD curl -f http://localhost:8080 || exit 1

ENTRYPOINT ["/kickstart_generator"]
