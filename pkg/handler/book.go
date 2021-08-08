package handler

import (
	api_go "github.com/KonstantinP85/api-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary create book
// @Security ApiKeyAuth
// @Tags book
// @Description create book
// @ID create-book
// @Accept json
// @Produce json
// @Param input body api_go.BookInput true "book info"
// @Success 200 {integer} integer 1
// @Failure 400/404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/api/book [post]
func (h *Handler) createBook(c *gin.Context)  {
	var input api_go.BookInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest,"invalid request body")
	}
	id, err := h.services.Book.CreateBook(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type GetBooksResponse struct {
	Result []api_go.BookResponse
}

// @Summary get books
// @Security ApiKeyAuth
// @Tags book
// @Description get books
// @ID get-books
// @Accept json
// @Produce json
// @Param authorId query int false "Author id"
// @Success 200 {object} GetBooksResponse
// @Failure 400/404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/api/book [get]
func (h *Handler) getBookList(c *gin.Context)  {
	authorId, err := strconv.Atoi(c.Request.URL.Query().Get("authorId"))
	if err != nil {
		authorId = 0
	}

	books, err := h.services.Book.GetBooks(authorId)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetBooksResponse{
		Result: books,
	})
}

// @Summary get book by id
// @Security ApiKeyAuth
// @Tags book
// @Description get book
// @ID get-book
// @Accept json
// @Produce json
// @Param id path int true "Book ID"
// @Success 200 {object} api_go.BookResponse
// @Failure 400/404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/api/book/{id} [get]
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
