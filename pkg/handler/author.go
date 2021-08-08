package handler

import (
	api_go "github.com/KonstantinP85/api-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// @Summary create author
// @Security ApiKeyAuth
// @Tags author
// @Description create author
// @ID create-author
// @Accept json
// @Produce json
// @Param input body api_go.AuthorInput true "author info"
// @Success 200 {integer} integer 1
// @Failure 400/404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/api/author [post]
func (h *Handler) createAuthor(c *gin.Context)  {
	var input api_go.AuthorInput

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest,"invalid request body")
	}
	id, err := h.services.Author.CreateAuthor(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type GetAuthorsResponse struct {
	Result []api_go.Author
}

// @Summary get authors
// @Security ApiKeyAuth
// @Tags author
// @Description get authors
// @ID get-authors
// @Accept json
// @Produce json
// @Success 200 {integer} GetAuthorsResponse
// @Failure 400/404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/api/author [get]
func (h *Handler) getAuthorList(c *gin.Context)  {
	authors, err := h.services.Author.GetAuthors()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetAuthorsResponse{
		Result: authors,
	})
}

// @Summary get author by id
// @Security ApiKeyAuth
// @Tags author
// @Description get author
// @ID get-author
// @Accept json
// @Produce json
// @Param id path int true "Author ID"
// @Success 200 {object} api_go.AuthorResponse
// @Failure 400/404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/api/author/{id} [get]
func (h *Handler) getAuthor(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	author, err := h.services.Author.GetAuthor(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, author)
}

func (h *Handler) updateAuthor(c *gin.Context)  {

}

func (h *Handler) deleteAuthor(c *gin.Context)  {

}