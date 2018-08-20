package user

import (
	. "apiservice/handler"
	"apiservice/model"
	"apiservice/pkg/errno"

	"github.com/gin-gonic/gin"
)

// @Summary Get an user by the user indentifier
// @Description Get an user by username
// @Tags user
// @Accept json
// @Produce json
// @Param username path string true "Username"
// @Success 200 {object} modle.UserModel "{"code":0,"message":"OK","data":{"username":"swagger","passwrod":"$2a$10$E0kwtmtLZbwW/bDQ8qI8e.eHPqhQOW9tvjwpyo/p05f/f4Qvr3OmS"}}"
// @Router /user/{username} [get]
func Get(c *gin.Context) {
	username := c.Param("username")
	// Get the by the `username` from the database.
	user, err := model.GetUser(username)
	if err != nil {
		SendResponse(c, errno.ErrUserNotFound, nil)
		return
	}

	SendResponse(c, nil, user)
}
