package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	nina_api_grpc "github.com/ningenMe/mami-interface/mami-generated-server/nina-api-grpc"
	"github.com/ningenme/nina-api/pkg/controller"
	"github.com/ningenme/nina-api/pkg/infra"
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
	infra.NingenmeMysql, err = sqlx.Open("mysql", infra.GetMysqlConfig("ningenme").FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	defer infra.NingenmeMysql.Close()

	infra.ComproMysql, err = sqlx.Open("mysql", infra.GetMysqlConfig("compro").FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	defer infra.NingenmeMysql.Close()

	s := grpc.NewServer()
	reflection.Register(s)

	{
		nina_api_grpc.RegisterGithubContributionServiceServer(s, &controller.GithubContributionController{})
		nina_api_grpc.RegisterBlogServiceServer(s, &controller.BlogController{})
		nina_api_grpc.RegisterHealthServiceServer(s, &controller.HealthController{})
		nina_api_grpc.RegisterComproCategoryServiceServer(s, &controller.ComproCategoryController{})
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
