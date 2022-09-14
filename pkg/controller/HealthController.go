package controller

import (
	"context"
	"github.com/ningenMe/mami-interface/nina-api-grpc/mami"
	"google.golang.org/protobuf/types/known/emptypb"
)

type HealthController struct {
	mami.UnimplementedHealthServiceServer
}

func (c *HealthController) Get(ctx context.Context, empty *emptypb.Empty) (*mami.GetHealthResponse, error) {
	return &mami.GetHealthResponse{
		Message: "ok",
	}, nil
}

