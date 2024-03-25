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
	server.setupCourseRoutes(apiRoutes.Group("/courses"))
	server.setupClassRoutes(apiRoutes.Group("/classes"))
	server.setupEnrollmentRoutes(apiRoutes.Group("/enrollments"))
	server.setupSemesterRoutes(apiRoutes.Group("/semesters"))
	server.setuuAttendanceRoutes(apiRoutes.Group("/attendances"))
}

func (server *Server) setupBasicRoutes(routes *gin.RouterGroup) {
	server.router.GET("", server.homePage)
}

func (server *Server) setupUserRoutes(routes *gin.RouterGroup) {
	server.router.POST("/register", server.createUser)
	server.router.POST("/login", server.loginUser)
}

func (server *Server) setupDepartmentRoutes(routes *gin.RouterGroup) {
	server.router.POST("/", server.createDepartment)
	server.router.GET("/", server.listDepartment)
	server.router.GET("/:id", server.getDepartmentById)
	server.router.PUT("/:id", server.updateDepartment)
	server.router.DELETE("/:id", server.deleteDepartment)
}

func (server *Server) setupCourseRoutes(routes *gin.RouterGroup) {
	server.router.POST("/", server.createCourse)
	server.router.GET("/", server.listCourse)
	server.router.GET("/:id", server.getCourseById)
	server.router.PUT("/:id", server.updateCourse)
	server.router.DELETE("/:id", server.deleteCourse)
}

func (server *Server) setupClassRoutes(routes *gin.RouterGroup) {
	server.router.POST("/", server.createClass)
	server.router.GET("/", server.listClass)
	server.router.GET("/:id", server.getClassById)
	server.router.PUT("/:id", server.updateClass)
	server.router.DELETE("/:id", server.deleteClass)
}

func (server *Server) setupEnrollmentRoutes(routes *gin.RouterGroup) {
	server.router.POST("/", server.createEnrollment)
	server.router.GET("/", server.listEnrollment)
	server.router.GET("/:id", server.getEnrollmentById)
	server.router.PUT("/:id", server.updateEnrollment)
	server.router.DELETE("/:id", server.deleteEnrollment)
}

func (server *Server) setupSemesterRoutes(routes *gin.RouterGroup) {
	server.router.POST("/", server.createSemester)
	server.router.GET("/", server.listSemester)
	server.router.GET("/:id", server.getSemesterById)
	server.router.PUT("/:id", server.updateSemester)
	server.router.DELETE("/:id", server.deleteSemester)
}

func (server *Server) setuuAttendanceRoutes(routes *gin.RouterGroup) {
	server.router.POST("/", server.createAttendance)
	server.router.GET("/", server.listAttendance)
	server.router.GET("/:id", server.getAttendanceById)
	server.router.PUT("/:id", server.updateAttendance)
	server.router.DELETE("/:id", server.deleteAttendance)
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
