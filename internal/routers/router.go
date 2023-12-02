package routers

import (
	_ "BloginGin/docs"
	"BloginGin/global"
	"BloginGin/internal/middleware"
	"BloginGin/internal/routers/api"
	v1 "BloginGin/internal/routers/api/v1"
	"BloginGin/pkg/limiter"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"
)

var methodLimiters = limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
	Key:          "/auth",
	FillInterval: time.Second,
	Capacity:     10,
	Quantum:      10,
})

func NewRouter() *gin.Engine {
	r := gin.New()
	if global.ServerSetting.RunMode == "debug" {
		r.Use(gin.Logger())
		r.Use(gin.Recovery())
	} else {
		r.Use(middleware.AccessLog())
		r.Use(middleware.Recovery())
	}

	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(60 * time.Second))
	r.Use(middleware.Translations())
	// 注册一个针对Swagger的路由
	// http://localhost:9090/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	article := v1.NewArticle()
	tag := v1.NewTag()

	upload := NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.GET("/auth", api.GetAuth)

	apiv1 := r.Group("/api/v1")

	apiv1.Use(middleware.JWT())
	{
		apiv1.POST("/tags", tag.Create)
		apiv1.DELETE("/tags/:id", tag.Delete)
		apiv1.PUT("/tags/:id", tag.Update)
		apiv1.PATCH("/tags/:id/state", tag.Update)
		apiv1.GET("/tags", tag.List)

		apiv1.POST("/articles", article.Create)
		apiv1.DELETE("/articles/:id", article.Delete)
		apiv1.PUT("/articles/:id", article.Update)
		apiv1.PATCH("/article/:id/state", article.Update)
		apiv1.GET("/articles/:id", article.Get)
		apiv1.GET("/articles", article.List)
	}
	return r
}
