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
	server.setupUserRoutes(apiRoutes.Group("/users"))
	_ = apiRoutes.Use(middleware.AuthMiddleware(server.tokenMaker, server.service))
	server.setupCourseRoutes(apiRoutes.Group("/courses"))
	server.setupClassRoutes(apiRoutes.Group("/classes"))
	server.setupEnrollmentRoutes(apiRoutes.Group("/enrollments"))
	server.setupSemesterRoutes(apiRoutes.Group("/semesters"))
	server.setuuAttendanceRoutes(apiRoutes.Group("/attendances"))
}

func (server *Server) setupBasicRoutes(routes *gin.RouterGroup) {
	routes.GET("", server.homePage)
}

func (server *Server) setupUserRoutes(routes *gin.RouterGroup) {
	routes.POST("/register", server.createUser)
	routes.POST("/login", server.loginUser)
}

func (server *Server) setupCourseRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createCourse)
	routes.GET("", server.listCourse)
	routes.GET("/:id", server.getCourseById)
	routes.PUT("/:id", server.updateCourse)
	routes.DELETE("/:id", server.deleteCourse)
}

func (server *Server) setupClassRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createClass)
	routes.GET("", server.listClass)
	routes.GET("/:id", server.getClassById)
	routes.PUT("/:id", server.updateClass)
	routes.DELETE("/:id", server.deleteClass)
}

func (server *Server) setupEnrollmentRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createEnrollment)
	routes.GET("", server.listEnrollment)
	routes.GET("/:id", server.getEnrollmentById)
	routes.PUT("/:id", server.updateEnrollment)
	routes.DELETE("/:id", server.deleteEnrollment)
}

func (server *Server) setupSemesterRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createSemester)
	routes.GET("", server.listSemester)
	routes.GET("/:id", server.getSemesterById)
	routes.PUT("/:id", server.updateSemester)
	routes.DELETE("/:id", server.deleteSemester)
}

func (server *Server) setuuAttendanceRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createAttendance)
	routes.GET("", server.listAttendance)
	routes.GET("/:id", server.getAttendanceById)
	routes.PUT("/:id", server.updateAttendance)
	routes.DELETE("/:id", server.deleteAttendance)
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
