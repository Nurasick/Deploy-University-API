package main

import (
	"context"
	"university/config"
	_ "university/docs" // Import docs for swagger
	"university/helpers/db"
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
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {

	cfg := config.Load()

	db.Connect(cfg.DBUrl)
	defer db.Conn.Close(context.Background())

	// Initialize repositories
	userRepo := repository.NewUserRepository(db.Conn)
	studentRepo := repository.NewStudentRepository(db.Conn)
	scheduleRepo := repository.NewScheduleRepository(db.Conn)
	attendanceRepo := repository.NewAttendanceRepository(db.Conn)
	teacherRepo := repository.NewTeacherRepository(db.Conn)
	subjectRepo := repository.NewSubjectRepository(db.Conn)

	// Initialize services
	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	studentService := service.NewStudentService(studentRepo, userRepo, attendanceRepo)
	teacherService := service.NewTeacherService(*teacherRepo, *userRepo, *scheduleRepo)
	scheduleService := service.NewScheduleService(scheduleRepo)
	attendanceService := service.NewAttendanceService(attendanceRepo, studentRepo, subjectRepo)

	// Initialize handlers
	authHandler := handler.NewAuthHandler(authService)
	userHandler := handler.NewUserHandler(userService)
	studentHandler := handler.NewStudentHandler(studentService)
	teacherHandler := handler.NewTeacherHandler(teacherService, attendanceService)
	scheduleHandler := handler.NewScheduleHandler(scheduleService)
	attendanceHandler := handler.NewAttendanceHandler(attendanceService)
	//adminHandler := handler.NewAdminHandler(userService, studentService, teacherService, scheduleService)

	e := echo.New()

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

	// Admin routes

	/*
		adminRoutes := e.Group("/api/admin", middleware.JWTAuth)
		adminRoutes.POST("/students", adminHandler.CreateStudent)
		adminRoutes.POST("/teachers", adminHandler.CreateTeacher)
		adminRoutes.POST("/subjects", adminHandler.CreateSubject)
	*/
	e.Logger.Fatal(e.Start(":" + cfg.Port))
}
