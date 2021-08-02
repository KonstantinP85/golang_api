package handler

import (
	"github.com/KonstantinP85/api-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) createUser(c *gin.Context)  {

}

type GetUsersResponse struct {
	Result []api_go.User
}

// @Summary get users
// @Security ApiKeyAuth
// @Tags user
// @Description get all user
// @ID get-all-user
// @Accept json
// @Produce json
// @Success 200 {object} GetUsersResponse
// @Failure 400/404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/api/user [get]
func (h *Handler) getUsers(c *gin.Context)  {
	users, err := h.services.User.GetUsers()
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, GetUsersResponse{
		Result: users,
	})
}

// @Summary get user by id
// @Security ApiKeyAuth
// @Tags user
// @Description get user
// @ID get-user
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} api_go.User
// @Failure 400/404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /admin/api/user/{id} [get]
func (h *Handler) getUser(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	user, err := h.services.User.GetUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *Handler) updateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	var input api_go.UpdateUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.User.UpdateUser(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resultResponse{
		Result: "Ok",
	})
}

func (h *Handler) patchUser(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}
	var input api_go.PatchUserInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	err = h.services.User.Patch(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

func (h *Handler) deleteUser(c *gin.Context)  {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "Invalid id param")
		return
	}

	err = h.services.User.DeleteUser(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, resultResponse{
		Result: "Ok",
	})
}