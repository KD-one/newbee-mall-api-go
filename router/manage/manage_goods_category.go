package manage

import (
	"github.com/gin-gonic/gin"
	v1 "main.go/api/v1"
	"main.go/middleware"
)

type ManageGoodsCategoryRouter struct {
}

func (r *ManageGoodsCategoryRouter) InitManageGoodsCategoryRouter(Router *gin.RouterGroup) {
	goodsCategoryRouter := Router.Group("v1").Use(middleware.AdminJWTAuth())

	var goodsCategoryApi = v1.ApiGroupApp.ManageApiGroup.ManageGoodsCategoryApi
	{
		goodsCategoryRouter.POST("categories", goodsCategoryApi.CreateCategory)
		goodsCategoryRouter.PUT("categories", goodsCategoryApi.UpdateCategory)
		goodsCategoryRouter.GET("categories", goodsCategoryApi.GetCategoryList)
		goodsCategoryRouter.GET("categories/:id", goodsCategoryApi.GetCategory)
		goodsCategoryRouter.DELETE("categories", goodsCategoryApi.DelCategory)
		// TODO：这个路由有错误：由于 URL 中没有名为 id 的路径参数，函数中直接调用 c.Param("id") 在这个示例中将无法获取到任何值
		// goodsCategoryRouter.GET("categories4Select", goodsCategoryApi.ListForSelect)
		// fix1：
		goodsCategoryRouter.GET("categories4Select/:id", goodsCategoryApi.ListForSelect)
		// fix2：
		// goodsCategoryRouter.GET("categories4Select", goodsCategoryApi.ListForSelect) goodsCategoryApi.ListForSelect函数中使用c.Query("id")获取参数
	}
}
