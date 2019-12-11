package api

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"hgin/models"
	"hgin/pkg/e"
	"hgin/pkg/util"
	"log"
	"net/http"
)

type auth struct {
	Username string `valid:"Required;MaxSize(50)"`
	Password string `valid:"Required;MaxSize(50)"`
	Yzm      string `valid:"Required;MaxSize(50)"`
}

// @Summary 获取token
// @tags auth
// @Produce  json
// @Param username query string true "用户名(eg:test)"
// @Param password query string true "密码(eg:test123456)"
// @Param yzm formData string true "密码(eg:limq)"
// @Success 200 {string} string "{"code":200,"data":{"token":"xxx"},"msg":"success"}"
// @Router /auth [post]
func GetAuth(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	yzm := c.PostForm("yzm")

	valid := validation.Validation{}
	a := auth{Username: username, Password: password, Yzm: yzm}

	ok, _ := valid.Valid(&a)

	data := make(map[string]interface{})

	code := e.INVALID_PARAMS

	if ok {
		isExist := models.CheckAuth(username, password)
		if isExist {
			token, err := util.GenerateToken(username, password)
			if err != nil {
				code = e.ERROR_AUTH_TOKEN
			} else {
				data["token"] = token
				code = e.SUCCESS
			}
		} else {
			code = e.ERROR_AUTH
		}

	} else {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
