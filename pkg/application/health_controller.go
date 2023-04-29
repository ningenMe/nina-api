package application

import (
	"context"
	"github.com/bufbuild/connect-go"
	ninav1 "github.com/ningenMe/nina-api/proto/gen_go/v1"
)

type HealthController struct{}

func (s *HealthController) Check(
	ctx context.Context,
	req *connect.Request[ninav1.HealthServiceCheckRequest],
) (*connect.Response[ninav1.HealthServiceCheckResponse], error) {
	return connect.NewResponse[ninav1.HealthServiceCheckResponse](&ninav1.HealthServiceCheckResponse{
		Status: ninav1.HealthServiceCheckResponse_SERVING_STATUS_SERVING,
	}), nil
}
