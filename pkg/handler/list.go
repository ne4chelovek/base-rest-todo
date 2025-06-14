package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"net/http"
	"strconv"
)

const (
	userID = "userId"
)

// getAllListResponse represents response for get all lists
// @Description Response object containing array of todo lists
type getAllListResponse struct {
	Data []*model.TodoList
}

// @Summary Create todo list
// @Tags lists
// @Security ApiKeyAuth
// @Description create todo list
// @ID create-list
// @Accept json
// @Produce json
// @Param input body model.TodoList true "list info"
// @Success 200 {object} map[string]interface{} "id"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists [post]
func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}
	var input *model.TodoList
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.listService.Create(c.Request.Context(), userId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

// @Summary Get All Lists
// @Tags lists
// @Security ApiKeyAuth
// @Description get all lists
// @ID get-all-lists
// @Accept json
// @Produce json
// @Success 200 {object} getAllListResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists [get]
func (h *Handler) getAllList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	lists, err := h.listService.GetAll(c.Request.Context(), userId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllListResponse{
		Data: lists,
	})

}

// @Summary Get List By Id
// @Tags lists
// @Security ApiKeyAuth
// @Description get list by id
// @ID get-list-by-id
// @Accept json
// @Produce json
// @Param id path int true "list id"
// @Success 200 {object} model.TodoList
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id} [get]
func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	lists, err := h.listService.GetById(c.Request.Context(), userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, lists)
}

// @Summary Update List
// @Tags lists
// @Security ApiKeyAuth
// @Description update list
// @ID update-list
// @Accept json
// @Produce json
// @Param id path int true "list id"
// @Param input body model.UpdateListInput true "update info"
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id} [put]
func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input *model.UpdateListInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.listService.Update(c.Request.Context(), userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.StatusResponse{Status: "OK"})
}

// @Summary Delete List
// @Tags lists
// @Security ApiKeyAuth
// @Description delete list
// @ID delete-list
// @Accept json
// @Produce json
// @Param id path int true "list id"
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id} [delete]
func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.listService.Delete(c.Request.Context(), userId, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, model.StatusResponse{Status: "OK"})
}
