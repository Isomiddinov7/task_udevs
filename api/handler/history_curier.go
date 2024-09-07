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

// CreateHistoryCurier godoc
// @ID create_history_curier
// @Router /history-curier [POST]
// @Summary Create HistoryCurier
// @Description Create HistoryCurier
// @Tags HistoryCurier
// @Accept json
// @Produce json
// @Param profile body models.CreateHistoryCurier true "CreateHistoryCurierBody"
// @Success 200 {object} Response{data=models.HistoryCurier} "HistoryCurierBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) CreateHistoryCurier(c *gin.Context) {
	var createHistory models.CreateHistoryCurier

	err := c.ShouldBindJSON(&createHistory)
	if err != nil {
		handleResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()
	err = h.strg.HistoryCurier().Create(
		ctx,
		createHistory,
	)
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusCreated, "ok")
}

// GetByIdHistoryCurier godoc
// @ID get_by_id_history_curier
// @Router /history-curier/{id} [GET]
// @Summary Get By Id HistoryCurier
// @Description Get By Id HistoryCurier
// @Tags HistoryCurier
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.HistoryCurier} "GetByIDHistoryCurierResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIDHistoryCurier(c *gin.Context) {
	var id = c.Param("id")
	if !utils.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.HistoryCurier().GetByID(ctx, models.HistoryCurierPrimaryKey{Id: id})
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

// GetListHistoryCurier godoc
// @ID get_list_history_curier
// @Router /history-curier [GET]
// @Summary Get List HistoryCurier
// @Description Get List HistoryCurier
// @Tags HistoryCurier
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Success 200 {object} Response{data=models.GetListHistoryCurierResponse} "GetListHistoryCurierResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListHistoryCurier(c *gin.Context) {

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

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.HistoryCurier().GetList(ctx, models.GetListHistoryCurierRequest{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		handleResponse(c, http.StatusInternalServerError, err)
		return
	}

	handleResponse(c, http.StatusOK, resp)
}
