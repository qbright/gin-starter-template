package controllers

import (
	"gin-starter-template/modules"
	"gin-starter-template/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u *UserController) Retrieve(c *gin.Context) {

	user := modules.User{}

	err := c.ShouldBindUri(&user)

	if err != nil {
		utils.Elog().Msg(err.Error())
	}

	user.GetById()
	c.JSON(http.StatusOK, &user)
}

func (u *UserController) Create(c *gin.Context) {
	user := modules.User{}
	c.ShouldBindJSON(&user)
	user.Save()
	c.JSON(http.StatusOK, &user)
}
