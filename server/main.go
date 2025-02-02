package main

import (
	"log"
	"math/rand"
	"net"
	"time"

	pb "github.com/raphaelrreis/vehicle-tracking"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedVehicleTrackingServiceServer
}

func (s *server) StreamLocation(req *pb.Empty, stream pb.VehicleTrackingService_StreamLocationServer) error {
	vehicleIDs := []string{"VEHICLE_1", "VEHICLE_2", "VEHICLE_3"}
	rand.Seed(time.Now().UnixNano())

	for {
		for _, vehicleID := range vehicleIDs {
			location := &pb.VehicleLocation{
				VehicleId: vehicleID,
				Latitude:  rand.Float64()*180 - 90,  // Latitude entre -90 e 90
				Longitude: rand.Float64()*360 - 180, // Longitude entre -180 e 180
				Timestamp: time.Now().Unix(),
			}

			if err := stream.Send(location); err != nil {
				return err
			}
		}
		time.Sleep(2 * time.Second) // Intervalo de atualização
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Falha ao ouvir: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterVehicleTrackingServiceServer(grpcServer, &server{})

	log.Println("Servidor gRPC iniciado na porta 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Falha ao servir: %v", err)
	}
}
