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

type Controller struct {
	mami.UnimplementedGithubContributionServer
}
var repository = infra.ContributionRepository{}

func (c *Controller) Get(ctx context.Context, empty *emptypb.Empty) (*mami.GetGithubContributionResponse, error) {
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

func (c *Controller) Post(stream mami.GithubContribution_PostServer) error {

	flag := true
	for flag {
		var list []*domainmodel.Contribution
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				flag = false
				break
			}
			if len(list) >= 3000 {
				break
			}
			t, _ := time.Parse(time.RFC3339,req.Contribution.GetContributedAt())
			list = append(list, &domainmodel.Contribution{
				ContributedAt: t,
				Organization: req.Contribution.GetOrganization(),
				Repository: req.Contribution.GetRepository(),
				User: req.Contribution.GetUser(),
				Status: req.Contribution.GetStatus(),
			})
		}
		repository.InsertList(list)
	}

	return stream.SendAndClose(&emptypb.Empty{})
}

