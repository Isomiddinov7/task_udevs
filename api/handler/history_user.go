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

// GetByIdHistoryUser godoc
// @ID get_by_id_history_user
// @Router /history-user/{id} [GET]
// @Summary Get By Id HistoryUser
// @Description Get By Id HistoryUser
// @Tags User
// @Accept json
// @Produce json
// @Param id path string true "id"
// @Success 200 {object} Response{data=models.HistoryUser} "GetByIDHistoryUserResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetByIDHistoryUser(c *gin.Context) {
	var id = c.Param("id")
	if !utils.IsValidUUID(id) {
		handleResponse(c, http.StatusBadRequest, "id is not uuid")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.CtxTimeout)
	defer cancel()

	resp, err := h.strg.HistoryUser().GetByID(ctx, models.HistoryUserPrimaryKey{Id: id})
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

// GetListHistoryUser godoc
// @ID get_list_history_user
// @Router /history-user [GET]
// @Summary Get List HistoryUser
// @Description Get List HistoryUser
// @Tags User
// @Accept json
// @Produce json
// @Param limit query number false "limit"
// @Param offset query number false "offset"
// @Param search query string false "search"
// @Success 200 {object} Response{data=models.GetHistoryUserListResponse} "GetHistoryUserListResponseBody"
// @Response 400 {object} Response{data=string} "Invalid Argument"
// @Failure 500 {object} Response{data=string} "Server Error"
func (h *Handler) GetListHistoryUser(c *gin.Context) {

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

	resp, err := h.strg.HistoryUser().GetList(ctx, models.GetHistoryUserListRequest{
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
