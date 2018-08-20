package user

import (
	. "apiservice/handler"
	"apiservice/model"
	"apiservice/pkg/errno"
	"apiservice/util"

	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

// @Summary Add new user to the database
// @Description Add a new user
// @Tags user
// @Accept  json
// @Produce  json
// @Param user body user.CreateRequest true "Create a new user"
// @Success 200 {object} user.CreateResponse "{"code":0,"message":"OK","data":{"username":"swagger"}}"
// @Router /user [post]
func Create(c *gin.Context) {
	log.Info("User Create funciton called.", lager.Data{"X-Requesst-Id": util.GetReqID(c)})
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		SendResponse(c, errno.ErrBind, nil)
		return
	}

	u := model.UserModel{Username: r.Username, Password: r.Password}

	// Valiate the data.
	if err := u.Validate(); err != nil {
		SendResponse(c, errno.ErrValidation, nil)
		return
	}

	// Encrypt the user passowrd.
	if err := u.Encrypt(); err != nil {
		SendResponse(c, errno.ErrEncrypt, nil)
		return
	}

	// Insert the user to the database.
	if err := u.Create(); err != nil {
		log.Infof("create user error: %s", err)
		SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	rsp := CreateResponse{Username: r.Username}

	// Show the user information.
	SendResponse(c, nil, rsp)
}

func (r *CreateRequest) checkparam() error {
	if r.Username == "" {
		return errno.New(errno.ErrValidation, nil).Add("username is empty.")
	}

	if r.Password == "" {
		return errno.New(errno.ErrValidation, nil).Add("password is empty.")
	}

	return nil
}
