package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	repo "bri.co.id/nds/approval-service/pkg/adapter/persistence/inmemory"
	"bri.co.id/nds/approval-service/pkg/approvalpb"
	"google.golang.org/grpc"
)

type server struct {
	r repo.SupervisorIMem
}

func NewServer(r repo.SupervisorIMem) *server {
	return &server{r: r}
}

func (*server) GetSupervisors(ctx context.Context, req *approvalpb.SupervisorRequest) (*approvalpb.SupervisorResponse, error) {

	return nil, nil
}

func main() {
	fmt.Println("Approval gRPC service is starting...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
		os.Exit(1)
	}

	s := grpc.NewServer()
	r := repo.NewSupervisorIMem()
	approvalpb.RegisterSupervisorServiceServer(s, NewServer(*r))
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
		os.Exit(2)
	}
}
