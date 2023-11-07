package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromartinsb/gomerit/schemas"
)

// @BasePath /api/v1

// @Summary Show holding
// @Description Show a holding
// @Tags Holdings
// @Accept json
// @Produce json
// @Param id query string true "Holding identification"
// @Success 200 {object} ShowHoldingResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /holding [get]
func ShowHoldingHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "queryParameter").Error())
		return
	}

	holding := schemas.Holding{}

	if err := db.First(&holding, id).Error; err != nil {
		sendError(ctx, http.StatusNotFound, "holding not found")
		return
	}

	sendSuccess(ctx, "show-holding", holding)
}
