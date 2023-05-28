
# Bitcoin Rate Sender

Simple RESTful web service API with Go and the Gin Web Framework

## Run with Docker

Build the Docker image:

```bash
docker build -t golang-bitcoin-rate-sender .
```

Run the Docker container

```bash
docker run -p 8080:8080 golang-bitcoin-rate-sender
```

Open app by the link

```bash
http://localhost:8080
```

## API Reference

#### Get the current BTC to UAH exchange rate

```http
  GET /rate
```

#### Subscribe an email address to receive the current exchange rate.

```http
  POST /subscribe
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `email`   | `string` | **Required**. Email to subscribe  |

####  Send an email with the current exchange rate to all subscribed email addresses.

```http
  POST /sendEmails
```