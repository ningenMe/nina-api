package controller

import (
	"context"
	"github.com/ningenMe/mami-interface/nina-api-grpc/mami"
	"github.com/ningenme/nina-api/pkg/domainmodel"
	"github.com/ningenme/nina-api/pkg/infra"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"sort"
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

func (c *GithubContributionController) GetSummary(ctx context.Context, req *mami.GetGithubContributionSummaryRequest) (*mami.GetGithubContributionSummaryResponse, error) {
	summaryList := repository.GetSummaryList(req.GetUser())
	sort.Slice(summaryList, func(i, j int) bool { return summaryList[i].Date < summaryList[j].Date })

	var pullRequest []*mami.ContributionSummary
	var comment     []*mami.ContributionSummary
	var approve     []*mami.ContributionSummary

	for _, summary := range summaryList {
		cs := mami.ContributionSummary{
			Date:  summary.Date,
			Count: int32(summary.Count),
		}

		switch summary.Status {
		case "CREATED_PULL_REQUEST":
			pullRequest = append(pullRequest, &cs)
		case "COMMENTED":
			comment = append(comment, &cs)
		case "APPROVED":
			approve = append(approve, &cs)
		}
	}

	return &mami.GetGithubContributionSummaryResponse{
		PullRequestSummaries: pullRequest,
		CommentedSummaries: comment,
		ApprovedSummaries: approve,
	}, nil
}
