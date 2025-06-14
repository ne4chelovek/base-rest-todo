package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/ne4chelovek/base-rest-todo/internal/service"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "github.com/ne4chelovek/base-rest-todo/pkg/docs"
)

type Handler struct {
	authServices service.Authorization
	itemService  service.TodoItem
	listService  service.TodoList
	tokenService service.Token
}

func NewHandler(authServices service.Authorization, listService service.TodoList, itemServices service.TodoItem, tokenService service.Token) *Handler {
	return &Handler{authServices: authServices, listService: listService, itemService: itemServices, tokenService: tokenService}
}

func (h *Handler) InitRouts() *gin.Engine {
	router := gin.New()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.userIdentity)
	{
		list := api.Group("/lists")
		{
			list.POST("/", h.createList)
			list.GET("/", h.getAllList)
			list.GET("/:id", h.getListById)
			list.PUT("/:id", h.updateList)
			list.DELETE("/:id", h.deleteList)

			items := list.Group(":id/items")
			{
				items.POST("/", h.createItem)
				items.GET("/", h.getAllItem)

			}
		}
		items := api.Group("items")
		{
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}
	return router
}
