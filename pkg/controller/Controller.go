package controller

import (
	"context"
	"github.com/ningenMe/mami-interface/nina-api-grpc/mami"
	"github.com/ningenme/nina-api/pkg/infra"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Controller struct {
	mami.UnimplementedGithubContributionServer
}

func (c *Controller) Get(ctx context.Context, empty *emptypb.Empty) (*mami.GetGithubContributionResponse, error) {
	repository := infra.ContributionRepository{}
	list := repository.GetList()

	viewList := []*mami.Contribution{}
	for _, contribution := range list {

		viewList = append(viewList, &mami.Contribution{
			ContributedAt: contribution.Time.String(),
			Organization: contribution.Org,
			Repository: contribution.Repo,
			User: contribution.User,
			Status: contribution.Status,
		})
	}

	return &mami.GetGithubContributionResponse{
		Contributions: viewList,
	}, nil
}

//
//type Request struct {
//	ContributionList []Contribution `json:"contributionList"`
//}
//
//type Contribution struct {
//	Time   time.Time `json:"time"`
//	Org    string    `json:"org"`
//	Repo   string    `json:"repo"`
//	User   string    `json:"user"`
//	Status string    `json:"status"`
//}
//
//func (Controller) PostGithubContributionList(c *gin.Context) {
//	//repository := infra.ContributionRepository{}
//	//contributionList := []*infra.Contribution {
//	//	{Time: time.Now(), Org: "org1",Repo: "hoge", User: "aa",Status: "xxx"},
//	//	{Time: time.Now(), Org: "org2",Repo: "fuga", User: "bb",Status: "yyy"},
//	//	{Time: time.Now(), Org: "org3",Repo: "piyo", User: "cc",Status: "zzz"},
//	//}
//	var req Request
//	if err := c.ShouldBindJSON(&req); err != nil {
//		fmt.Println(err)
//		c.JSON(http.StatusBadRequest, gin.H{"message": "error"})
//		return
//	}
//	fmt.Println(req)
//
//	//var contributionList []infra.Contribution
//	//repository.InsertList(contributionList)
//	c.JSON(http.StatusOK, gin.H{
//		"message": "ok",
//	})
//}

