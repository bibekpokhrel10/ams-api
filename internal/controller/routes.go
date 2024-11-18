package controller

import (
	docs "github.com/ams-api/docs"
	"github.com/ams-api/internal/middleware"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (server *Server) setupRouter() {
	server.router.Use(middleware.CORSMiddleware())
	server.setupSwagger()
	apiRoutes := server.router.Group("/api/v1")

	server.setupBasicRoutes(apiRoutes)
	server.setupLoginRoutes(apiRoutes)
	_ = apiRoutes.Use(middleware.AuthMiddleware(server.tokenMaker, server.service))
	server.setupUserRoutes(apiRoutes.Group("/users"))
	server.setupInstitutionRoutes(apiRoutes.Group("/institutions"))
	server.setupProgramRoutes(apiRoutes.Group("/programs"))
	server.setupSemesterRoutes(apiRoutes.Group("/semesters"))
	server.setupCourseRoutes(apiRoutes.Group("/courses"))
	server.setupClassRoutes(apiRoutes.Group("/classes"))
	server.setupEnrollmentRoutes(apiRoutes.Group("/classes/enrollments"))
	server.setupProgramEnrollmentRoutes(apiRoutes.Group("/programs/enrollments"))
	server.setupAttendanceRoutes(apiRoutes.Group("/attendances"))
}

func (server *Server) setupBasicRoutes(routes *gin.RouterGroup) {
	routes.GET("", server.homePage)
	routes.GET("/institutions", server.listInstitution)
}

func (server *Server) setupUserRoutes(routes *gin.RouterGroup) {
	routes.GET("/profile", server.getUserFromToken)
	routes.GET("", server.getAllUsers)
	routes.PUT("/:id/activate", server.activateUser)
	routes.PUT("/:id/change-password", server.updateUserPassword)
	routes.PUT("/:id/change-type", server.updateUserType)
	routes.DELETE("/:id", server.deleteUser)
}

func (server *Server) setupLoginRoutes(routes *gin.RouterGroup) {
	routes.POST("/register", server.createUser)
	routes.POST("/login", server.loginUser)
}

func (server *Server) setupInstitutionRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createInstitution)
	routes.GET("/:id", server.getInstitutionById)
	routes.PUT("/:id", server.updateInstitution)
	routes.DELETE("/:id", server.deleteInstitution)
}

func (server *Server) setupProgramRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createProgram)
	routes.GET("", server.listProgram)
	routes.GET("/:id", server.getProgramById)
	routes.PUT("/:id", server.updateProgram)
	routes.DELETE("/:id", server.deleteProgram)
}

func (server *Server) setupSemesterRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createSemester)
	routes.GET("", server.listSemester)
	routes.GET("/:id", server.getSemesterById)
	routes.PUT("/:id", server.updateSemester)
	routes.DELETE("/:id", server.deleteSemester)
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

func (server *Server) setupProgramEnrollmentRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createProgramEnrollment)
	routes.GET("", server.listProgramEnrollment)
	routes.GET("/:id", server.getProgramEnrollmentById)
	routes.PUT("/:id", server.updateProgramEnrollment)
	routes.DELETE("/:id", server.deleteProgramEnrollment)
}

func (server *Server) setupAttendanceRoutes(routes *gin.RouterGroup) {
	routes.POST("", server.createAttendance)
	routes.GET("", server.listAttendance)
	routes.GET("/:id", server.getAttendanceById)
	routes.PUT("/:id", server.updateAttendance)
	routes.DELETE("/:id", server.deleteAttendance)
}

// Swagger host path and basepath configuration
func (server *Server) setupSwagger() {
	hostPath := server.config.HOSTPATH
	if hostPath == "" {
		hostPath = "localhost:8080"
	}
	basePath := server.config.BASEPATH
	if hostPath == "" {
		basePath = "/api/v1"
	}
	docs.SwaggerInfo.Host = hostPath
	docs.SwaggerInfo.BasePath = basePath
	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
