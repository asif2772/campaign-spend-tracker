# Campaign Spend Tracker

## Tech Stack

* Go
* PostgreSQL
* Redis
* Docker Compose

---

## Requirements

* Go 1.22+
* Docker & Docker Compose
* Git

---

## Clone

```bash
git clone https://github.com/asif2772/campaign-spend-tracker.git
cd campaign-spend-tracker
```

---

## Start Dependencies

```bash
docker compose up -d
```

This automatically starts PostgreSQL and Redis and creates the required database table.

---

## Run Application

```bash
go run ./cmd/server
```

Server:

http://localhost:8080

---

## Health Check

```bash
curl http://localhost:8080/healthz
```

Expected:

OK


---

## Create Spend Event

POST

http://localhost:8080/spend


Headers:

Content-Type: application/json
X-Tenant-ID: 10

Body
{
  "campaign_id": 100,
  "amount_micros": 500
}

---

## Get Daily Running Spend

GET

http://localhost:8080/spend/100


Headers:

X-Tenant-ID: 10

---

## Useful Commands

```bash
go fmt ./...
go vet ./...
go build ./...
go test -race ./...
```

---

## Project Structure

cmd/
internal/
    handler/
    model/
    publisher/
    repository/
    service/
    tenant/
pkg/
scripts/


------------------------RUN SIMPLY-----------------------------
git clone https://github.com/asif2772/campaign-spend-tracker.git
cd campaign-spend-tracker

docker compose up -d
go run ./cmd/server