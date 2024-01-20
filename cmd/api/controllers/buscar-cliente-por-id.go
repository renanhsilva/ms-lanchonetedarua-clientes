package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	use_cases "github.com/tbrito/ms-lanchonetedarua-clientes/internal/application/use-cases/buscar-cliente-por-id"
	"net/http"
)

func BuscarClientePorId(ctx *gin.Context, useCase *use_cases.BuscarClientePorIdUseCase, id uuid.UUID) {

	cliente, err := useCase.Execute(id)

	if err != nil {
		ctx.AbortWithStatusJSON(
			http.StatusNotFound,
			gin.H{
				"success": false,
			})
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{
			"sucess":  true,
			"cliente": cliente,
		})
}
