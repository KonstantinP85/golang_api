package handler

import (
	"github.com/KonstantinP85/api-go/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api")
	{
		 book := api.Group("/book")
		 {
			 book.GET("/", h.getBookList)
			 book.GET("/:id", h.getBook)
		 }

		 author := api.Group("/author")
		 {
			 author.GET("/", h.getAuthorList)
			 author.GET("/:id", h.getAuthor)
		 }
	}
	admin := router.Group("/admin")
	{
		api := admin.Group("/api")
		{
			book := api.Group("/book")
			{
				book.POST("/", h.createBook)
				book.PUT("/:id", h.updateBook)
				book.DELETE("/:id", h.deleteBook)
			}

			author := api.Group("/author")
			{
				author.POST("/", h.createAuthor)
				author.PUT("/:id", h.updateAuthor)
				author.DELETE("/:id", h.deleteAuthor)
			}

			user := api.Group("/user")
			{
				user.GET("/", h.getUsers)
				user.GET("/:id", h.getUser)
				user.POST("/", h.createUser)
				user.PUT("/:id", h.updateUser)
				user.PATCH("/:id", h.patchUser)
				user.DELETE("/:id", h.deleteUser)
			}
		}
	}

	return router
}
