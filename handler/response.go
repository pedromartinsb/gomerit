package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pedromartinsb/gomerit/schemas"
)

func sendError(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(code, gin.H{
		"message":   msg,
		"errorCode": code,
	})
}

func sendSuccess(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("operation from handler: %s successfull", op),
		"data":    data,
	})
}

type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

// Opening

type CreateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type DeleteOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ShowOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

type ListOpeningsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.OpeningResponse `json:"data"`
}

type UpdateOpeningResponse struct {
	Message string                  `json:"message"`
	Data    schemas.OpeningResponse `json:"data"`
}

// Holding

type CreateHoldingResponse struct {
	Message string                  `json:"message"`
	Data    schemas.HoldingResponse `json:"data"`
}

type ShowHoldingResponse struct {
	Message string                  `json:"message"`
	Data    schemas.HoldingResponse `json:"data"`
}

type ListHoldingsResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.HoldingResponse `json:"data"`
}

type UpdateHoldingResponse struct {
	Message string                  `json:"message"`
	Data    schemas.HoldingResponse `json:"data"`
}

type DeleteHoldingResponse struct {
	Message string                  `json:"message"`
	Data    schemas.HoldingResponse `json:"data"`
}

// Opening

type CreateCompanyResponse struct {
	Message string                  `json:"message"`
	Data    schemas.CompanyResponse `json:"data"`
}

type ListCompaniesResponse struct {
	Message string                    `json:"message"`
	Data    []schemas.CompanyResponse `json:"data"`
}
