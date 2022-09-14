package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ningenMe/mami-interface/nina-api-grpc/mami"
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


	infra.NingenmeMysql, err = sqlx.Open("mysql",  infra.NingenmeMysqlConfig.GetConfig())
	if err != nil {
		log.Fatalln(err)
	}
	defer infra.NingenmeMysql.Close()


	s := grpc.NewServer()
	reflection.Register(s)
	mami.RegisterGithubContributionServer(s, &controller.GithubContributionController{})
	mami.RegisterHealthServer(s, &controller.HealthController{})
	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
