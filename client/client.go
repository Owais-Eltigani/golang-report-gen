package main

import (
	"context"
	"log"

	report "github.com/report-gen/reports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	// making a connection to server
	conn, err := grpc.NewClient("localhost:5000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {

		log.Fatal("client: error while conn to server", err)
		return
	}
	//
	defer conn.Close()

	clinet := report.NewReportServiceClient(conn)
	res, err := clinet.GenerateReport(context.Background(), &report.ReportRequest{UserId: "5"})

	if err != nil {
		log.Fatal("client: error in response", err)

		return
	}

	log.Print("response is: ", res.Status)

}
