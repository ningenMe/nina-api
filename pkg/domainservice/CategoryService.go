package domainservice

import (
	nina_api_grpc "github.com/ningenMe/mami-interface/mami-generated-server/nina-api-grpc"
	"github.com/ningenme/nina-api/pkg/domainmodel"
	"github.com/ningenme/nina-api/pkg/infra"
)

type CategoryService struct{}

var categoryRepository = infra.CategoryRepository{}

func (s *CategoryService) Post(req *nina_api_grpc.PostCategoryRequest) {
	categoryId := req.GetCategoryId()
	if req.GetCategory() != nil {
		categoryId = req.GetCategoryId()
		if categoryId == "" {
			categoryId = domainmodel.GetNewCategoryId()
		}
		categoryRepository.Upsert(
			&domainmodel.Category{
				CategoryId:          categoryId,
				CategoryDisplayName: req.GetCategory().GetCategoryDisplayName(),
				CategorySystemName:  req.GetCategory().GetCategorySystemName(),
				CategoryOrder:       req.GetCategory().GetCategoryOrder(),
			})
	} else {
		categoryRepository.Delete(categoryId)
	}
}
