# Campaign Spend Tracker

## Requirements

* Go 1.22+
* Docker
* Git

---

## Clone Repository

git clone https://github.com/asif2772/campaign-spend-tracker.git

---

## Start Dependencies

docker compose up -d

This starts:

* PostgreSQL
* Redis

---

## Verify Containers

docker ps

---

## Run Application

go run ./cmd/server

The server starts on:

http://localhost:8080

---

## Health Check

curl http://localhost:8080/healthz

Expected:

OK

---

## Create Spend Event (postman or other platform)

url: http://localhost:8080/spend
method: POST

header: {
Content-Type: application/json
X-Tenant-ID: 10
} 

body: data {
    "campaign_id":10,
    "amount_micros":50
}

---

## Get Daily Running Spend

url: http://localhost:8080/spend/100
method: GET
header: {
    X-Tenant-ID: 10
    }

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

---

## Features

* Layered architecture
* PostgreSQL persistence
* Redis running daily totals
* Tenant isolation
* Async event publisher
* Health endpoint

---

## Useful Commands

go fmt ./...
go vet ./...
go build ./...
go test -race ./...

