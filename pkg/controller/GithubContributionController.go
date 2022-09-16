package controller

import (
	"context"
	"github.com/ningenMe/mami-interface/nina-api-grpc/mami"
	"github.com/ningenme/nina-api/pkg/domainmodel"
	"github.com/ningenme/nina-api/pkg/infra"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"time"
)

type GithubContributionController struct {
	mami.UnimplementedGithubContributionServiceServer
}
var repository = infra.ContributionRepository{}

func (c *GithubContributionController) Get(ctx context.Context, empty *emptypb.Empty) (*mami.GetGithubContributionResponse, error) {
	list := repository.GetList()

	viewList := []*mami.Contribution{}
	for _, contribution := range list {

		viewList = append(viewList, &mami.Contribution{
			ContributedAt: contribution.ContributedAt.Format(time.RFC3339),
			Organization: contribution.Organization,
			Repository: contribution.Repository,
			User: contribution.User,
			Status: contribution.Status,
		})
	}

	return &mami.GetGithubContributionResponse{
		Contributions: viewList,
	}, nil
}

func (c *GithubContributionController) Post(stream mami.GithubContributionService_PostServer) error {

	var list []*domainmodel.Contribution
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		}
		for _, contribution := range req.GetContributions() {
			t, _ := time.Parse(time.RFC3339,contribution.GetContributedAt())
			list = append(list, &domainmodel.Contribution{
				ContributedAt: t,
				Organization: contribution.GetOrganization(),
				Repository: contribution.GetRepository(),
				User: contribution.GetUser(),
				Status: contribution.GetStatus(),
			})
		}
	}
	repository.InsertList(list)

	return stream.SendAndClose(&emptypb.Empty{})
}

func (c *GithubContributionController) Delete(ctx context.Context, req *mami.DeleteGithubContributionRequest) (*emptypb.Empty, error) {
	startAt, _ := time.Parse(time.RFC3339,req.GetStartAt())
	endAt, _ := time.Parse(time.RFC3339,req.GetEndAt())
	repository.Delete(startAt, endAt)
	return &emptypb.Empty{}, nil
}