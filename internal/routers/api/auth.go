package api

import (
	"github.com/Renewdxin/BlogInGin/global"
	"github.com/Renewdxin/BlogInGin/internal/service"
	"github.com/Renewdxin/BlogInGin/pkg/app"
	"github.com/Renewdxin/BlogInGin/pkg/errcode"
	"github.com/gin-gonic/gin"
)

func GetAuth(c *gin.Context) {
	param := service.AuthRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c.Request.Context(), "app BindAndValid errs: %v", errs)
		response.ToErrorResponse(errcode.UnauthorizedAuthNotExist)
		return
	}

	token, err := app.GenerateToken(param.AppKey, param.AppSecret)
	if err != nil {
		global.Logger.Errorf(c.Request.Context(), "app GenerateToken err: %v", err)
		response.ToErrorResponse(errcode.UnauthorizedTokenGenerate)
		return
	}

	response.ToResponse(gin.H{
		"token": token,
	})
}
