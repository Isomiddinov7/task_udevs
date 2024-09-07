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

// CreateProduct godoc
// @ID create_product
// @Router /product [POST]
// @Summary Create Product
// @Description Create Product
// @Tags Product
// @Accept multipart/form-data
// @Produce json
// @Param name formData file true "Name"
// @Param comment formData string true "Comment"
// @Param price formData string true "Price"
// @Param file formData string true "Upload file"
// @Success 200 {object} Response{data=models.Product} "ProductBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateProduct(c *gin.Context) {
	var createProduct models.CreateProduct

	file, err := c.FormFile("file")
	if err != nil {
		handleResponse(c, 400, gin.H{"error": "Unable to get file"})
		return
	}

	imageURL, err := utils.UploadImage(file)
	if err != nil {
		handleResponse(c, 500, gin.H{"error": "Failed to upload image"})
		return
	}

	createProduct.Name = c.PostForm("name")
	createProduct.Comment = c.PostForm("comment")
	createProduct.Price = c.PostForm("price")
	createProduct.ProductImg = imageURL

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	resp, err := h.strg.Product().Create(ctx, createProduct)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, resp)
}

// GetByIdProduct godoc
// @ID get_by_id_product
// @Router /product/{id} [GET]
// @Summary Get By Id Product
// @Description Get By Id Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Product} "GetByIDProductResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIDProduct(c *gin.Context) {
	var id = c.Param("id")
	if !utils.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Product().GetByID(ctx, models.ProductPrimaryKey{Id: id})
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

// GetListProduct godoc
// @ID get_list_product
// @Router /product [GET]
// @Summary Get List Product
// @Description Get List Product
// @Tags Product
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Param search query string false "search"
// @Success 200 {object} Response{data=models.GetProductListResponse} "GetListProductResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListProduct(c *gin.Context) {

	limit, err := getIntegerOrDefaultValue(c.Query("limit"), 10)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query limit")
		return
	}

	offset, err := getIntegerOrDefaultValue(c.Query("offset"), 0)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query offset")
		return
	}

	search := c.Query("search")
	if err != nil {
		handleResponse(c, http.StatusBadRequest, "invalid query search")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Product().GetList(ctx, models.GetProductListRequest{
		Limit:  limit,
		Offset: offset,
		Search: search,
	})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}

// UpdateProduct godoc
// @ID update_product
// @Router /product/{id} [PUT]
// @Summary  Update Product
// @Description Update Product
// @Tags UpdateProduct
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Param object body models.UpdateProduct true "UpdateProductRequestBody"
// @Success 200 {object} Response{data=models.UpdateProduct} "Product"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) UpdateProduct(c *gin.Context) {

	var updateProduct models.UpdateProduct

	err := c.ShouldBindJSON(&updateProduct)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	var id = c.Param("id")
	if !utils.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	updateProduct.Id = id

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	rowsAffected, err := h.strg.Product().Update(ctx, updateProduct)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	if rowsAffected == 0 {
		handleResponse(c, http.StatusBadRequest, "no rows affected")
		return
	}

	ctx, cancel = context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.Product().GetByID(ctx, models.ProductPrimaryKey{Id: updateProduct.Id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusAccepted, resp)
}

// DeleteProduct godoc
// @ID delete_product
// @Router /product/{id} [DELETE]
// @Summary Delete Product
// @Description Delete Product
// @Tags Product
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.Product} "DeleteProductResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) DeleteProduct(c *gin.Context) {
	var id = c.Param("id")
	if !utils.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	err := h.strg.Product().Delete(ctx, models.ProductPrimaryKey{Id: id})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusNoContent, nil)
}
