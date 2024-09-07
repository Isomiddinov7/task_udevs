package handler

import (
	"context"
	"net/http"
	"task_udevs/api/models"
	"task_udevs/config"

	"github.com/gin-gonic/gin"
)

// CreateAdditionProduct godoc
// @ID create_addition_product
// @Router /addition-product [POST]
// @Summary Create AdditionProduct
// @Description Create AdditionProduct
// @Tags AdditionProduct
// @Accept json
// @Produce json
// @Param profile body models.CreateAdditionProduct true "CreateAdditionProductBody"
// @Success 200 {object} Response{data=string} "ProductBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateAdditionProduct(c *gin.Context) {
	var create_addition models.CreateAdditionProduct

	err := c.ShouldBindJSON(&create_addition)
	if err != nil {
		handleResponse(c, 400, "BadRequest")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	resp, err := h.strg.AdditionProduct().Create(ctx, create_addition)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetAdditionProductByID godoc
// @ID get_addtion_product_by_id
// @Router /addition-product/{id} [GET]
// @Summary Get AdditionProduct  By ID
// @Description Get AdditionProduct  By ID
// @Tags Product
// @Accept json
// @Produce json
// @Param product_id path string true "product_id"
// @Success 200 {object} Response{data=models.GetAdditionProductByIdResponse} "AdditionProductBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetAdditionProductByID(c *gin.Context) {

	product_id := c.Param("id")
	resp, err := h.strg.AdditionProduct().GetByID(
		context.Background(),
		models.GetAdditionProductById{
			ProductId: product_id,
		},
	)

	if err != nil {
		handleResponse(c, 500, "GRPCError"+err.Error())
		return
	}

	handleResponse(c, http.StatusOK, resp)
}
