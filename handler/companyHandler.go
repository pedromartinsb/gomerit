package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromartinsb/gomerit/schemas"
)

// @BasePath /api/v1

// @Summary Create company
// @Description Create a new company
// @Tags Companies
// @Accept json
// @Produce json
// @Param request body CreateCompanyRequest true "Request body"
// @Success 200 {object} CreateCompanyResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /company [post]
func CreateCompanyHandler(ctx *gin.Context) {
	request := &CreateCompanyRequest{}

	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	company := schemas.Company{
		Name:      request.Name,
		HoldingID: request.HoldingID,
	}

	if err := db.Create(&company).Error; err != nil {
		logger.Errorf("error creating company: %v", err.Error())
		sendError(ctx, http.StatusInternalServerError, "error creating company on database")
		return
	}

	sendSuccess(ctx, "create-company", company)
}
