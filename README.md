# Go Webhook Processor

This project is a simple HTTP server written in Go that listens for incoming JSON payloads, transforms the data, and forwards the transformed data to a specified webhook URL.

## Features


- HTTP server listening on port `8080`
- JSON payload processing

- Data transformation from `InputData` to `OutputData`
- Asynchronous request handling

- Forwarding transformed data to a webhook URL

## Getting Started

### Prerequisites

Before running this server, make sure you have Go installed on your machine. [Download Go](https://golang.org/dl/).

### Installation

To set up this project locally, clone the repository using:

```bash
git clone https://github.com/Anandvp92/go-webhook-processor.git
cd go-webhook-processor\go-webhook
go run main.go


