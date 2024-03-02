package controller

import (
	"github.com/ams-api/internal/middleware"
	"github.com/gin-gonic/gin"
)

func (server *Server) setupRouter() {
	server.router.Use(middleware.CORSMiddleware())
	// server.setupSwagger()
	apiRoutes := server.router.Group("/api/v1")

	server.setupBasicRoutes(apiRoutes)
	//_ = apiRoutes.Use(middleware.AuthMiddleware(server.tokenMaker, server.service, server.cache))
	server.setupUserRoutes(apiRoutes.Group("/users"))
	server.setupDepartmentRoutes(apiRoutes.Group("/departments"))
}

func (server *Server) setupBasicRoutes(routes *gin.RouterGroup) {
	server.router.GET("", server.homePage)
}

func (server *Server) setupUserRoutes(routes *gin.RouterGroup) {
	server.router.POST("/register", server.createUser)
	server.router.POST("/login", server.loginUser)
}

func (server *Server) setupDepartmentRoutes(routes *gin.RouterGroup) {
	server.router.POST("/departments", server.createDepartment)
	server.router.GET("/departments", server.listDepartment)
	server.router.GET("/departments/:id", server.getDepartmentById)
	server.router.PUT("/departments/:id", server.updateDepartment)
	server.router.DELETE("/departments/:id", server.deleteDepartment)
}

// Swagger host path and basepath configuration
// func (server *Server) setupSwagger() {
// 	hostPath := server.config.HOSTPATH
// 	if string_helper.IsEmptyString(hostPath) {
// 		hostPath = "localhost:8080"
// 	}
// 	basePath := server.config.BASEPATH
// 	if string_helper.IsEmptyString(hostPath) {
// 		basePath = "/api/v1"
// 	}
// 	docs.SwaggerInfo.Host = hostPath
// 	docs.SwaggerInfo.BasePath = basePath
// 	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
// }
