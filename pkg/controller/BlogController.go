package controller

import (
	nina_api_grpc "github.com/ningenMe/mami-interface/mami-generated-server/nina-api-grpc"
	"github.com/ningenMe/nina-api/pkg/infra"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type BlogController struct {
	nina_api_grpc.UnimplementedBlogServiceServer
}

var blogRepository = infra.BlogRepository{}

func (c *BlogController) Get(empty *emptypb.Empty, stream nina_api_grpc.BlogService_GetServer) error {
	list := blogRepository.GetList()
	n := len(list)
	i := 0
	for {
		tmp := list[i%n]
		blog := nina_api_grpc.Blog{
			Url:   tmp.Url,
			Date:  tmp.Date,
			Type:  tmp.Type,
			Title: tmp.Title,
		}
		if err := stream.Send(&nina_api_grpc.GetBlogResponse{Blog: &blog}); err != nil {
			return err
		}
		time.Sleep(time.Second * 20)
		i += 1
		i %= n
	}
}
