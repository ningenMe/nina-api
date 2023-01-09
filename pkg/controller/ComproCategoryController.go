package controller

import (
	"context"
	nina_api_grpc "github.com/ningenMe/mami-interface/mami-generated-server/nina-api-grpc"
	"github.com/ningenme/nina-api/pkg/domainservice"
	"github.com/ningenme/nina-api/pkg/infra"
	"google.golang.org/protobuf/types/known/emptypb"
)

type ComproCategoryController struct {
	nina_api_grpc.UnimplementedComproCategoryServiceServer
}
var categoryService = domainservice.CategoryService{}
var categoryRepository = infra.CategoryRepository{}

func (c *ComproCategoryController) Post(ctx context.Context, req *nina_api_grpc.PostCategoryRequest) (*emptypb.Empty, error)  {
	categoryService.Post(req)
	return &emptypb.Empty{}, nil
}

func (c *ComproCategoryController) Get(context.Context, *emptypb.Empty) (*nina_api_grpc.GetCategoryResponse, error) {
	list := categoryRepository.GetList()

	var viewList []*nina_api_grpc.Category
	for _, category := range list {

		viewList = append(viewList, &nina_api_grpc.Category{
			CategoryId:          category.CategoryId,
			CategoryDisplayName: category.CategoryDisplayName,
			CategorySystemName:  category.CategorySystemName,
			CategoryOrder:       category.CategoryOrder,
		})
	}

	return &nina_api_grpc.GetCategoryResponse{
		CategoryList: viewList,
	}, nil
}
