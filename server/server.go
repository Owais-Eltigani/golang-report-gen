package main

import (
	"context"
	"log"
	"net"

	pb "github.com/report-gen/reports"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ? implementing the report interface
type Server struct {
	pb.UnimplementedReportServiceServer
}

func (s *Server) GenerateReport(ctx context.Context, reportRequest *pb.ReportRequest) (*pb.ReportResponse, error) {
	if reportRequest.UserId == "" {
		log.Fatal("server: error the userid id missing")

		return nil, status.Errorf(codes.NotFound, "userid is missing")
	}

	log.Printf("report of user: %s, was logged", reportRequest.UserId)
	return &pb.ReportResponse{ReportId: reportRequest.UserId, Status: "success"}, nil
}

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
