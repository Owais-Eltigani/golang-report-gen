package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	report "github.com/report-gen/reports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type User struct {
	Username string
	UserId   string
}

var USERS = []User{
	{
		Username: "Bob",
		UserId:   "1",
	},
	{
		Username: "Alice",
		UserId:   "2",
	},
	{
		Username: "Charlie",
		UserId:   "3",
	},
	{
		Username: "David",
		UserId:   "4",
	},
	{
		Username: "Eve",
		UserId:   "5",
	},
	{
		Username: "Frank",
		UserId:   "6",
	},
	{
		Username: "Grace",
		UserId:   "7",
	},
	{
		Username: "Henry",
		UserId:   "8",
	},
	{
		Username: "Ivy",
		UserId:   "9",
	},
	{
		Username: "Jack",
		UserId:   "10",
	},
}

func main() {
	// Making a connection to server
	conn, err := grpc.NewClient("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("client: error while conn to server", err)
		return
	}
	defer conn.Close()

	// Create client once
	client := report.NewReportServiceClient(conn)

	// create crone call eachh 10 sec.
	log.Println("client calls start here...")
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	callGRPCService(client)

	for {
		select {
		case <-ticker.C:
			callGRPCService(client)
		}
	}
}

// crone logic
func callGRPCService(client report.ReportServiceClient) {

	// request timout of 5 sec
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	randomUserId := USERS[rand.Intn(10)].UserId //? generating a random user id of the users list

	res, err := client.GenerateReport(ctx, &report.ReportRequest{UserId: randomUserId})
	if err != nil {
		log.Printf("client: error in response: %v", err)
		return
	}

	log.Printf("ReportId is: %s", res.ReportId)
}
