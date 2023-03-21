package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type HealthController struct{}

func (h *HealthController) Status(c *gin.Context) {
	session := sessions.Default(c)

	var count int
	v := session.Get("count")

	if v == nil {
		count = 0
	} else {
		count = v.(int)
		count++
	}
	session.Set("count", count)
	session.Save()

	c.String(http.StatusOK, "Working!"+fmt.Sprint(count))
}
