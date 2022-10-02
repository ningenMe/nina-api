package controller

import (
	"context"
	nina_api_grpc "github.com/ningenMe/mami-interface/mami-generated-server/nina-api-grpc"
	"github.com/ningenme/nina-api/pkg/domainmodel"
	"github.com/ningenme/nina-api/pkg/domainservice"
	"github.com/ningenme/nina-api/pkg/infra"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"time"
)

type GithubContributionController struct {
	nina_api_grpc.UnimplementedGithubContributionServiceServer
}

var contributionRepository = infra.ContributionRepository{}
var service = domainservice.ContributionService{}

func (c *GithubContributionController) Get(ctx context.Context, empty *emptypb.Empty) (*nina_api_grpc.GetGithubContributionResponse, error) {
	list := contributionRepository.GetList()

	viewList := []*nina_api_grpc.Contribution{}
	for _, contribution := range list {

		viewList = append(viewList, &nina_api_grpc.Contribution{
			ContributedAt: contribution.ContributedAt.Format(time.RFC3339),
			Organization:  contribution.Organization,
			Repository:    contribution.Repository,
			User:          contribution.User,
			Status:        contribution.Status,
		})
	}

	return &nina_api_grpc.GetGithubContributionResponse{
		Contributions: viewList,
	}, nil
}

func (c *GithubContributionController) Post(stream nina_api_grpc.GithubContributionService_PostServer) error {

	var list []*domainmodel.Contribution
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		for _, contribution := range req.GetContributions() {
			t, _ := time.Parse(time.RFC3339, contribution.GetContributedAt())
			list = append(list, &domainmodel.Contribution{
				ContributedAt: t,
				Organization:  contribution.GetOrganization(),
				Repository:    contribution.GetRepository(),
				User:          contribution.GetUser(),
				Status:        contribution.GetStatus(),
			})
		}
	}
	contributionRepository.InsertList(list)

	return stream.SendAndClose(&emptypb.Empty{})
}

func (c *GithubContributionController) Delete(ctx context.Context, req *nina_api_grpc.DeleteGithubContributionRequest) (*emptypb.Empty, error) {
	startAt, _ := time.Parse(time.RFC3339, req.GetStartAt())
	endAt, _ := time.Parse(time.RFC3339, req.GetEndAt())
	contributionRepository.Delete(startAt, endAt)
	return &emptypb.Empty{}, nil
}

func (c *GithubContributionController) GetStatistics(ctx context.Context, req *nina_api_grpc.GetStatisticsRequest) (*nina_api_grpc.GetStatisticsResponse, error) {
	s := service.GetStatistics(req.GetUser())
	return s, nil
}
