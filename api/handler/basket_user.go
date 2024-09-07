package handler

import (
	"context"
	"database/sql"
	"net/http"
	"task_udevs/api/models"
	"task_udevs/config"
	"task_udevs/pkg/utils"

	"github.com/gin-gonic/gin"
)

// CreateCart godoc
// @ID create_cart
// @Router /cart [POST]
// @Summary Create Cart
// @Description Create Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param profile body models.CreateCart true "CreateCartBody"
// @Success 200 {object} Response{data=string} "CartBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateCart(c *gin.Context) {
	var create_cart models.CreateCart

	err := c.ShouldBindJSON(&create_cart)
	if err != nil {
		handleResponse(c, 400, "BadRequest")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	err = h.strg.Cart().Create(ctx, create_cart)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, "ok")
}

// GetByIdCartgodoc
// @ID get_by_id_cart
// @Router /cart/{id} [GET]
// @Summary Get By Id Cart
// @Description Get By Id Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Cart} "GetByIDCartResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIDCart(c *gin.Context) {
	var id = c.Param("id")
	if !utils.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Cart().GetByID(ctx, models.CartPrimaryKey{Id: id})
	if err == sql.ErrNoRows {
		handleResponse(c, http.StatusBadRequest, "no rows in result set")
		return
	}
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// DeleteCart godoc
// @ID delete_cart
// @Router /cart/{id} [DELETE]
// @Summary Delete Cart
// @Description Delete Cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=string} "DeleteCartResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteCart(c *gin.Context) {
	var id = c.Param("id")
	if !utils.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.Cart().Delete(ctx, models.CartPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}
