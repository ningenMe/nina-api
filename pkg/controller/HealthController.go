package controller

import (
	"context"
	"fmt"
	nina_api_grpc "github.com/ningenMe/mami-interface/mami-generated-server/nina-api-grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthController struct {
	nina_api_grpc.UnimplementedHealthServiceServer
}

func (c *HealthController) Get(ctx context.Context, empty *emptypb.Empty) (*nina_api_grpc.GetHealthResponse, error) {
	fmt.Println("health")
	return &nina_api_grpc.GetHealthResponse{
		Message: "ok",
	}, nil
}

