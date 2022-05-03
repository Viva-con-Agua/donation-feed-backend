FROM docker.io/golang:1.18-alpine

# prepare image
RUN apk add --no-cache tini
ENTRYPOINT [ "/sbin/tini", "--" ]

# add application
WORKDIR /app/src
ADD . /app/src/
ENV GOBIN=/usr/local/bin
RUN go mod download &&\
    go install /app/src

# add additional image metadata
ENV APP_PORT=80
EXPOSE 80/tcp
CMD [ "donation-feed-backend" ]
