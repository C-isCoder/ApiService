package user

import (
	"apiservice/handler"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Delete delete an user by the user indentifier.
func Delete(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("id"))
	if err := model.DeleteUser(uint64(userId)); err != nil {
		handler.SendResponse(c, errno.ErrDtabaser, nil)
		return
	}

	handler.SendResponse(c, nil, nil)
}
