package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/ningenMe/nina-api/pkg/application"
	"github.com/ningenMe/nina-api/pkg/infra"
	"github.com/ningenMe/nina-api/proto/gen_go/v1/ninav1connect"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

func main() {
	var err error

	infra.NingenmeMysql, err = sqlx.Open("mysql", infra.GetMysqlConfig("ningenme").FormatDSN())
	if err != nil {
		log.Fatalln(err)
	}
	defer infra.NingenmeMysql.Close()


	//server
	mux := http.NewServeMux()

	{
		nina := &application.NinaController{}
		path, handler := ninav1connect.NewNinaServiceHandler(nina)
		mux.Handle(path, handler)
	}

	corsHandler := cors.New(cors.Options{
		AllowedOrigins:   []string{"https://ningenme.net", "http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}).Handler(h2c.NewHandler(mux, &http2.Server{}))
	err = http.ListenAndServe(":8081", corsHandler)
	if err != nil {
		return
	}
}
