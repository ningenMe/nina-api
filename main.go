package main

import (
	"fmt"
	"github.com/ningenMe/mami-interface/nina-api-grpc/mami"
	"github.com/ningenme/nina-api/pkg/controller"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	mami.RegisterGithubContributionServer(s, &controller.Controller{})
	fmt.Println("server start")
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
