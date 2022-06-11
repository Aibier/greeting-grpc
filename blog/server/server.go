package main

import pb "github.com/Aibier/greeting-grpc/blog/proto"

type Server struct {
	pb.BlogServiceServer
}
