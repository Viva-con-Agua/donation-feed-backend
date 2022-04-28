# Donation Feed Backend

> An HTTP server that sends out events whenever a donation occurs

This backend is used by [play4water-overlay](https://github.com/Viva-con-Agua/play4water-overlay) to show
donations in real-time whenever they occur.
This service provides the backend for that.

## Api Specification

A client can request the server to send *donation* events on `/api/donation-events`.

The response consists of [Server-sent events](https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events)
with their data in the following format (defined
in [DonationEvent type](https://github.com/Viva-con-Agua/donation-feed-backend/blob/main/dao/donationEvent.go)):

```
{
  "name": <string or null>,
  "money": {
    "amount": <number>,
    "currency": <string>
  }
}
```
