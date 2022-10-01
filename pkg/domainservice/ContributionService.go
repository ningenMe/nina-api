package domainservice

import (
	nina_api_grpc "github.com/ningenMe/mami-interface/mami-generated-server/nina-api-grpc"
	"github.com/ningenme/nina-api/pkg/domainmodel"
	"github.com/ningenme/nina-api/pkg/infra"
	"time"
)

type ContributionService struct{}

var repository = infra.ContributionRepository{}

func (s *ContributionService) GetStatistics(user string) *nina_api_grpc.GetStatisticsResponse {
	mp := repository.GetSumMap(user)
	year, month, day := time.Now().Date()
	startAt := time.Date(2017, 12, 31, 0, 0, 0, 0, time.Local) //日曜日
	endAt := time.Date(year, month, day, 0, 0, 0, 0, time.Local)

	createdPullRequestStatistics := GetStatusStatistics(startAt, endAt, user, "CREATED_PULL_REQUEST", mp)
	commentedStatistics := GetStatusStatistics(startAt, endAt, user, "COMMENTED", mp)
	approvedStatistics := GetStatusStatistics(startAt, endAt, user, "APPROVED", mp)

	return &nina_api_grpc.GetStatisticsResponse{
		CreatedPullRequestStatistics: createdPullRequestStatistics,
		CommentedStatistics: commentedStatistics,
		ApprovedStatistics: approvedStatistics,
	}
}

func GetStatusStatistics(startAt time.Time, endAt time.Time, user string, status string, mp map[domainmodel.ContributionSumKey]int) *nina_api_grpc.ContributionStatistics  {

	var list []*nina_api_grpc.ContributionSum
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
		list = append(list, &nina_api_grpc.ContributionSum{
			Date: targetWeekDate,
			Sum: int32(sum),
		})
		totalSum += sum
	}

	return &nina_api_grpc.ContributionStatistics{
		Sum: int32(totalSum),
		ContributionSumList: list,
	}
}