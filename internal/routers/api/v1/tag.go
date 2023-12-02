package v1

import (
	"BloginGin/global"
	"BloginGin/internal/service"
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

// List handles the request to retrieve a list of tags.
// @Summary Retrieve a list of tags
// @Produce json
// @Param name query string false "Tag name" maxlength(100)
// @Param state query int false "State" Enums(0, 1) default(1)
// @Param page query int false "Page number"
// @Param page_size query int false "Number of items per page"
// @Success 200 {object} model.TagSwagger "Success"
// @Failure 400 {object} errcode.Error
// @Router /api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)

	// Bind and validate request parameters
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c.Request.Context(), "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	// Count the total number of tags based on request parameters
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf(c.Request.Context(), "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	// Retrieve the list of tags based on request parameters and pagination
	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf(c.Request.Context(), "svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	// Respond with the list of tags and total count
	response.ToResponseList(tags, totalRows)
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
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)

	// Bind and validate request parameters
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c.Request.Context(), "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf(c.Request.Context(), "svc.CreateTag err : %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c.Request.Context(), "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf(c.Request.Context(), "svc UpdateTagErr : %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Errorf(c.Request.Context(), "app.BindAndValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf(c.Request.Context(), "svc DeleteTagErr : %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
