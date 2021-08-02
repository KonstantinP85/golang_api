package handler

import (
	api_go "github.com/KonstantinP85/api-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createBook(c *gin.Context)  {
	
}

type GetBooksResponse struct {
	Result []api_go.Book
}

func (h *Handler) getBookList(c *gin.Context)  {
	books, err := h.services.Book.GetBooks()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetBooksResponse{
		Result: books,
	})
}

func (h *Handler) getBook(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	book, err := h.services.Book.GetBook(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, book)
}

func (h *Handler) updateBook(c *gin.Context)  {

}

func (h *Handler) deleteBook(c *gin.Context)  {

}
