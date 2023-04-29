package application

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/ningenme/nina-api/pkg/domainmodel"
	"github.com/ningenme/nina-api/pkg/infra"
	ninav1 "github.com/ningenme/nina-api/proto/gen_go/v1"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type NinaController struct{}

var contributionRepository = infra.ContributionRepository{}

func (s *NinaController) ContributionGet(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[ninav1.ContributionGetResponse], error) {
	list := contributionRepository.GetList()

	viewList := []*ninav1.Contribution{}
	for _, contribution := range list {

		viewList = append(viewList, &ninav1.Contribution{
			ContributedAt: contribution.ContributedAt.Format(time.RFC3339),
			Organization:  contribution.Organization,
			Repository:    contribution.Repository,
			User:          contribution.User,
			Status:        contribution.Status,
		})
	}

	return connect.NewResponse[ninav1.ContributionGetResponse](
		&ninav1.ContributionGetResponse{
			ContributionList: viewList,
		},
	), nil
}
func (s *NinaController) ContributionPost(
	ctx context.Context,
	req ninav1.ContributionPostRequest,
) (*connect.Response[emptypb.Empty], error) {

	//var list []*domainmodel.Contribution
	//for {
	//	req, err := stream.Recv()
	//	if err == io.EOF {
	//		break
	//	}
	//	for _, contribution := range req.GetContributions() {
	//		t, _ := time.Parse(time.RFC3339, contribution.GetContributedAt())
	//		list = append(list, &domainmodel.Contribution{
	//			ContributedAt: t,
	//			Organization:  contribution.GetOrganization(),
	//			Repository:    contribution.GetRepository(),
	//			User:          contribution.GetUser(),
	//			Status:        contribution.GetStatus(),
	//		})
	//	}
	//}
	//contributionRepository.InsertList(list)
	//
	//return stream.SendAndClose(&emptypb.Empty{})

	return connect.NewResponse[emptypb.Empty](
		&emptypb.Empty{},
	), nil
}
func (s *NinaController) ContributionDelete(
	ctx context.Context,
	req *connect.Request[ninav1.ContributionDeleteRequest],
) (*connect.Response[emptypb.Empty], error) {
	startAt, _ := time.Parse(time.RFC3339, req.Msg.StartAt)
	endAt, _ := time.Parse(time.RFC3339, req.Msg.EndAt)
	contributionRepository.Delete(startAt, endAt)

	return connect.NewResponse[emptypb.Empty](
		&emptypb.Empty{},
	), nil
}

func (s *NinaController) ContributionStatisticsGet(
	ctx context.Context,
	req *connect.Request[ninav1.ContributionStatisticsGetRequest],
) (*connect.Response[ninav1.ContributionStatisticsGetResponse], error) {
	user := req.Msg.User
	mp := contributionRepository.GetSumMap(user)
	year, month, day := time.Now().Date()
	startAt := time.Date(2017, 12, 31, 0, 0, 0, 0, time.Local) //日曜日
	endAt := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	createdPullRequestStatistics := GetStatusStatistics(startAt, endAt, user, "CREATED_PULL_REQUEST", mp)
	commentedStatistics := GetStatusStatistics(startAt, endAt, user, "COMMENTED", mp)
	approvedStatistics := GetStatusStatistics(startAt, endAt, user, "APPROVED", mp)

	return connect.NewResponse[ninav1.ContributionStatisticsGetResponse](
		&ninav1.ContributionStatisticsGetResponse{
			CreatedPullRequestStatistics: createdPullRequestStatistics,
			CommentedStatistics: commentedStatistics,
			ApprovedStatistics: approvedStatistics,
		},
	), nil
}


func GetStatusStatistics(startAt time.Time, endAt time.Time, user string, status string, mp map[domainmodel.ContributionSumKey]int) *ninav1.ContributionStatistics  {

	var list []*ninav1.ContributionSum
	totalSum := 0
	for targetWeekAt := startAt; targetWeekAt.Before(endAt); targetWeekAt = targetWeekAt.AddDate(0,0,7) {

		sum := 0
		for i := 0; i < 7; i++ {
			targetDayDate := targetWeekAt.AddDate(0,0,i).Format("2006-01-02")
			key := domainmodel.ContributionSumKey{
				Date: targetDayDate,
				User: user,
				Status: status,
			}
			if _, ok := mp[key]; ok {
				sum += mp[key]
			}

		}

		targetWeekDate := targetWeekAt.Format("2006-01-02")
		list = append(list, &ninav1.ContributionSum{
			Date: targetWeekDate,
			Sum: int32(sum),
		})
		totalSum += sum
	}

	return &ninav1.ContributionStatistics{
		Sum: int32(totalSum),
		ContributionSumList: list,
	}
}