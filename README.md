# A gRPC Report Generation API

The Project consist of three files the contain server.go that implement the backend service .

## Project Structure

- server: contain server.go which implement the backend service
- client: acts as a client interface and call backend every 10 sec for a user
- reports: contain the scheme definition and protobuf files

## Usage

Download or clone the git repo `gRPC_report_gen`. then run command `go run server/server.go` to start listening for users requests, after that run go command `go run client/client.go` in different terminal tab the client will start sending requests every 10 sec.
