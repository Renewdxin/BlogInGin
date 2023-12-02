package app

import (
	"github.com/Renewdxin/BlogInGin/pkg/errcode"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 封装了对 Gin 框架的响应操作
type Response struct {
	Ctx *gin.Context
}

// Pager 用于表示分页信息
type Pager struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	TotalRows int `json:"total_rows"`
}

// NewResponse 创建一个新的 Response 实例
func NewResponse(ctx *gin.Context) *Response {
	return &Response{Ctx: ctx}
}

// ToResponse 将数据以 JSON 格式返回给客户端
func (r *Response) ToResponse(data interface{}) {
	if data == nil {
		data = gin.H{}
	}
	r.Ctx.JSON(http.StatusOK, data)
}

// ToResponseList 将列表数据和分页信息以 JSON 格式返回给客户端
func (r *Response) ToResponseList(list interface{}, totalRows int) {
	r.Ctx.JSON(http.StatusOK, gin.H{
		"list": list,
		"pager": Pager{
			Page:      GetPage(r.Ctx),
			PageSize:  GetPageSize(r.Ctx),
			TotalRows: totalRows,
		},
	})
}

// ToErrorResponse 将错误信息以 JSON 格式返回给客户端
func (r *Response) ToErrorResponse(err *errcode.Error) {
	response := gin.H{
		"code": err.Code(),
		"msg":  err.Msg(),
	}
	details := err.Details()
	if len(details) > 0 {
		response["details"] = details
	}
	r.Ctx.JSON(err.StatusCode(), response)
}
