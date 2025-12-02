package routes

import (
	"github.com/emuthianimbithi/pos-service/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterBranchRoutes(rg *gin.RouterGroup, h *handlers.BranchHandler) {
	branches := rg.Group("/branches")
	{
		branches.POST("", h.CreateBranch)
		branches.GET("", h.ListBranches)
		branches.GET("/:id", h.GetBranch)
		branches.PUT("/:id", h.UpdateBranch)
		branches.DELETE("/:id", h.DeleteBranch)
	}
}
