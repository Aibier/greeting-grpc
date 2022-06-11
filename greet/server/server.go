package main

import pb "github.com/Aibier/greeting-grpc/greet/proto"

type Server struct {
	pb.GreetServiceServer
}
