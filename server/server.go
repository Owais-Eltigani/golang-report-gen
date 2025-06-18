package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"

	pb "github.com/report-gen/reports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
	pb.UnimplementedReportServiceServer
}

// ? a simple global var to save all users requests and their reports ids
var usersReports = make(map[string][]string)

// ! implementing the report interface
func (s *Server) GenerateReport(ctx context.Context, reportRequest *pb.ReportRequest) (*pb.ReportResponse, error) {
	if reportRequest.UserId == "" {

		log.Fatal("server: error the userid id missing")
		return nil, status.Errorf(codes.NotFound, "userid is missing")
	}

	// ? randomly generating ids for reports
	reportid := rand.Intn(100)
	log.Printf("report of userid: %s, is: %d", reportRequest.UserId, reportid)

	// adding the request userid and the report id.
	usersReports[reportRequest.UserId] = append(usersReports[reportRequest.UserId], strconv.Itoa(reportid))

	// log.Printf("\n", usersReports[reportRequest.UserId], "\n")
	return &pb.ReportResponse{ReportId: strconv.Itoa(reportid)}, nil
}

//

func (s *Server) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	log.Print("Health check requested")
	return &pb.HealthCheckResponse{Status: "OK"}, nil
}

// =====================
func main() {

	// create a listener for users requests
	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatal("server: error while creating the listener", err)
		return
	}

	// init the grpc server,
	grpcServer := grpc.NewServer()
	s := Server{}

	pb.RegisterReportServiceServer(grpcServer, &s)

	//
	log.Print("server running...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal("server: error while init server", err)
		return
	}

}
