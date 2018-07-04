package user

import (
	"strconv"

	. "apiservice/handler"
	"apiservice/model"
	"apiservice/pkg/errno"

	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user indentifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	SendResponse(c, nil, nil)
}
