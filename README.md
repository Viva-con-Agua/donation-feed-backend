# Donation Feed Backend

> An HTTP server that sends out events whenever a donation occurs

This backend is used by [play4water-overlay](https://github.com/Viva-con-Agua/play4water-overlay) to show
donations in real-time whenever they occur.
This service provides the backend for that.

## Api Specification

A client can request the server to send it events by doing `GET /api/donation-events`.

The response consists of [Server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
with their data in the format defined by the [DonationEvent type](https://github.com/Viva-con-Agua/donation-feed-backend/blob/main/dao/donationEvent.go).

## How to Run (locally)

The *donation-feed-backend* supports local testing using a running mongo server but without NATS.
This works by the service periodically publishing dummy events and then processing them as if they would have come from
NATS.

To do this, execute the following commands:
```shell
docker run -it --rm -p 27017:27017 docker.io/mongo
```
and in a separate shell:
```shell
go run server.go -port 8080 -start-dummy-emitter
```
