package v1

import (
	"BloginGin/global"
	"BloginGin/pkg/app"
	"BloginGin/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Tag struct {
}

type CountTagRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type TagListRequest struct {
	Name  string `form:"name" binding:"max=100"`
	State uint8  `form:"state, default=1" binding:"oneof=0 1"`
}

type CreateTagRequest struct {
	Name      string `form:"name" binding:"required,min=3,max=100"`
	State     uint8  `form:"state, default=1" binding:"oneof=0 1"`
	CreatedBy string `form:"created_by" binding:"required,min=3,max=100"`
}

type UpdateTagRequest struct {
	ID         uint32 `form:"id" binding:"required,gte=1"`
	Name       string `form:"name" binding:"min=3,max=100"`
	State      uint8  `form:"state, default=1" binding:"oneof=0 1"`
	ModifiedBy string `form:"modified_by" binding:"required,min=3,max=100"`
}

type DeleteTagRequest struct {
	ID uint32 `form:"id" binding:"required,gte=1"`
}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {

}

// List @Summary
// @Produce json
// @Param name query string false "tag name" maxlength(100)
// @Param state query int false "state" Enums(0, 1) default(1)
// @Param page query int false "page"
// @Param page_size query int false "pageSize"
// @Success 200 {object} model.TagSwagger "success"
// @Failure 400 {object} errcode.Error
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := struct {
		Name  string `form:"name" binding:"max=100"`
		State uint8  `form:"state,default=1" binding:"oneof=0 1"`
	}{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf("app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// Create @Summary Create tags
// @Produce json
// @Param name body string true "tag body" minlength(3) maxlength(100)
// @Param state body int false "state" Enums(0, 1) default(1)
// @Param created_by body string true "founder" minlength(3) maxlength(100)
// @Success 200 {object} model.Tag "success"
// @Failure 400 {object} errcode.Error "request error"
// @Failure 500 {object} errcode.Error "Internal Server Error"
// @Router /api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {

}

func (t Tag) Update(c *gin.Context) {

}

func (t Tag) Delete(c *gin.Context) {

}
