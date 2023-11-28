package app

import (
	"BloginGin/global"
	"BloginGin/pkg/convert"
	"github.com/gin-gonic/gin"
)

// GetPage 从 Gin 上下文中获取页码参数，如果未提供或小于等于零，则返回默认值 1。
func GetPage(c *gin.Context) int {
	page := convert.StrTo(c.Query("page")).MustInt()
	if page <= 0 {
		return 1
	}
	return page
}

// GetPageSize 从 Gin 上下文中获取每页大小参数，如果未提供、小于等于零或大于最大允许值，则返回默认值。
func GetPageSize(c *gin.Context) int {
	pageSize := convert.StrTo(c.Query("page_size")).MustInt()
	if pageSize <= 0 {
		return global.AppSetting.DefaultPageSize
	}
	if pageSize > global.AppSetting.MaxPageSize {
		return global.AppSetting.DefaultPageSize
	}
	return pageSize
}

// GetPageOffset 根据给定的页码和每页大小计算偏移量，用于数据库查询。
func GetPageOffset(page, pageSize int) int {
	result := 0
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
