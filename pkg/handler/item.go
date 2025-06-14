package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ne4chelovek/base-rest-todo/internal/model"
	"net/http"
	"strconv"
)

// @Summary Create todo item
// @Tags items
// @Security ApiKeyAuth
// @Description create todo item
// @ID create-item
// @Accept json
// @Produce json
// @Param id path int true "list id"
// @Param input body model.TodoItem true "item info"
// @Success 200 {object} map[string]interface{} "id"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id}/items [post]
func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input *model.TodoItem
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := h.itemService.Create(c.Request.Context(), userId, listId, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id": id,
	})

}

// @Summary Get All Items
// @Tags items
// @Security ApiKeyAuth
// @Description get all items
// @ID get-all-items
// @Accept json
// @Produce json
// @Param id path int true "list id"
// @Success 200 {object} map[string]interface{} "items"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/lists/{id}/items [get]
func (h *Handler) getAllItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	items, err := h.itemService.GetAllItem(c.Request.Context(), userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"items": items,
	})
}

// @Summary Get Item By Id
// @Tags items
// @Security ApiKeyAuth
// @Description get item by id
// @ID get-item-by-id
// @Accept json
// @Produce json
// @Param id path int true "item id"
// @Success 200 {object} map[string]interface{} "item"
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [get]
func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	item, err := h.itemService.GetById(c.Request.Context(), userId, itemId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"items": item,
	})

}

// @Summary Update Item
// @Tags items
// @Security ApiKeyAuth
// @Description update item
// @ID update-item
// @Accept json
// @Produce json
// @Param id path int true "item id"
// @Param input body model.UpdateItemInput true "update info"
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [put]
func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	var input *model.UpdateItemInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	err = h.itemService.Update(c.Request.Context(), userId, id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.StatusResponse{Status: "OK"})
}

// @Summary Delete Item
// @Tags items
// @Security ApiKeyAuth
// @Description delete item
// @ID delete-item
// @Accept json
// @Produce json
// @Param id path int true "item id"
// @Success 200 {object} model.StatusResponse
// @Failure 400 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Router /api/items/{id} [delete]
func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id")
		return
	}

	err = h.itemService.Delete(c.Request.Context(), userId, listId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, model.StatusResponse{Status: "OK"})

}
