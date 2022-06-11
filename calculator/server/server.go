package main

import pb "github.com/Aibier/greeting-grpc/calculator/proto"

type Server struct {
	pb.CalculatorServiceServer
}
