package main

import (
	"log"
	"os"
	"university/config"
	"university/database"
	"university/docs"
	_ "university/docs" // Import docs for swagger
	"university/pkg/handler"
	"university/pkg/middleware"
	"university/pkg/repository"
	"university/pkg/service"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger" // echo-swagger middleware
)

// @title University API
// @version 1.0
// @description University Management System API
// @host {{HOST}}
// @BasePath /
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	host := os.Getenv("SWAGGER_HOST")
	if host == "" {
		host = "localhost:8080"
	}

	docs.SwaggerInfo.Host = host

	cfg := config.GetConfig()
	log.Printf("Config DB: host=%s port=%d user=%s dbname=%s url=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.DBName, cfg.DB.URL)
	log.Printf("Connecting to DB: %s", cfg.DB.String())
	conn := database.OpenConnectionPool(cfg.DB.String())
	defer database.CloseConnectionPool(conn)

	migration := database.OpenConnection(cfg.DB.String())
	defer database.CloseConnection(migration)

	sqlDB := database.PGXConnToSQLDB(migration)
	database.GooseMigrate(sqlDB, "./database/migrations")
	log.Println("Database migrated")

	dbWrapper := database.NewDBConnWrapper(sqlDB)
	// Initialize repositories
	userRepo := repository.NewUserRepository(conn)
	studentRepo := repository.NewStudentRepository(conn)
	attendanceRepo := repository.NewAttendanceRepository(conn)
	scheduleRepo := repository.NewScheduleRepository(conn)
	teacherRepo := repository.NewTeacherRepository(conn)
	subjectRepo := repository.NewSubjectRepository(conn)
	groupRepo := repository.NewGroupRepository(conn)

	// Initialize services
	groupService := service.NewGroupService(groupRepo)
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	studentService := service.NewStudentService(studentRepo, userRepo, attendanceRepo, userRepo)
	teacherService := service.NewTeacherService(*teacherRepo, *userRepo, *scheduleRepo)
	scheduleService := service.NewScheduleService(scheduleRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo, studentRepo, subjectRepo)

	// Initialize handlers
	GroupHandler := handler.NewGroupHandler(groupService)
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	studentHandler := handler.NewStudentHandler(studentService)
	teacherHandler := handler.NewTeacherHandler(teacherService, attendanceService)
	scheduleHandler := handler.NewScheduleHandler(scheduleService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService)
	//adminHandler := handler.NewAdminHandler(userService, studentService, teacherService, scheduleService)

	e := echo.New()

	healthHandler := handler.NewHealthHandler(dbWrapper, cfg)
	e.GET("/internal/health", healthHandler.Status)

	e.GET("/swagger/*", echoSwagger.EchoWrapHandler())
	jmtMW := middleware.JWTAuth
	// Auth routes
	auth := e.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	// User routes
	users := e.Group("/users", jmtMW)
	users.GET("/me", userHandler.Me)

	// Student routes
	studentRoutes := e.Group("/student", jmtMW)
	studentRoutes.GET("/:id", studentHandler.GetStudentByID)
	studentRoutes.GET("/:id/attendance", studentHandler.MyAttendance)
	studentRoutes.GET("/all", studentHandler.GetAllStudents)
	studentRoutes.POST("", studentHandler.CreateStudent)
	studentRoutes.PATCH("/:id", studentHandler.UpdateStudent)

	// Teacher routes
	teacherRoutes := e.Group("/teacher", jmtMW)
	teacherRoutes.GET("/:id", teacherHandler.GetTeacherByID)
	teacherRoutes.POST("", teacherHandler.CreateTeacher)
	// Schedule routes
	scheduleRoutes := e.Group("/schedule", jmtMW)
	scheduleRoutes.POST("", scheduleHandler.CreateSchedule)
	scheduleRoutes.GET("/all_class_schedule", scheduleHandler.GetAllSchedules)
	scheduleRoutes.GET("/group/:id", scheduleHandler.GetScheduleByGroupID)

	// Attendance routes
	attendanceRoutes := e.Group("/attendance", jmtMW)
	attendanceRoutes.POST("/subject", attendanceHandler.MarkAttendance)
	attendanceRoutes.GET("/attendanceBySubjectID/:id", attendanceHandler.GetAttendanceBySubjectID)
	attendanceRoutes.GET("/attendanceByStudentID/:id", attendanceHandler.GetAttendanceByStudentID)

	// Group routes
	groupRoutes := e.Group("/groups", jmtMW)
	groupRoutes.GET("", GroupHandler.GetAllStudents)
	groupRoutes.POST("", GroupHandler.CreateGroup)
	// Admin routes

	/*
		adminRoutes := e.Group("/api/admin", middleware.JWTAuth)
		adminRoutes.POST("/students", adminHandler.CreateStudent)
		adminRoutes.POST("/teachers", adminHandler.CreateTeacher)
		adminRoutes.POST("/subjects", adminHandler.CreateSubject)
	*/
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
