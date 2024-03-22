package routers

import (
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	controller "idstar.com/app/controllers"
	middleware "idstar.com/app/middleware"
	"idstar.com/app/repositories"
	"idstar.com/app/services"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Add Middleware
	logger := middleware.LoggerMiddleware{}
	r.Use(logger.Logger())
	r.Use(middleware.CheckToken())

	// Load instance UserRepository
	rekeningRepository := repositories.NewRekeningRepository()
	detailKaryawanRepository := repositories.NewDetailKaryawanRepository()
	trainingRepository := repositories.NewTrainingRepository()
	karyawanRepository := repositories.NewKaryawanRepository()
	karyawanTrainingRepository := repositories.NewKaryawanTrainingRepository()
	userRepository := repositories.NewUserRepository()

	// Load instance UserService
	rekeningService := services.NewRekeningService(*rekeningRepository, *karyawanRepository)
	trainingService := services.NewTrainingService(*trainingRepository, *karyawanTrainingRepository)
	karyawanService := services.NewKaryawanService(*karyawanRepository, *detailKaryawanRepository, *rekeningRepository, *karyawanTrainingRepository)
	karyawanTrainingService := services.NewKaryawanTrainingService(*karyawanTrainingRepository, *karyawanRepository, *trainingRepository, *detailKaryawanRepository)
	userService := services.NewUserService(*userRepository)
	authService := services.NewAuthenticationService(*userRepository)

	// Load instance UserController
	rekeningController := controller.NewRekeningController(rekeningService)
	trainingController := controller.NewTrainingController(trainingService)
	karyawanController := controller.NewKaryawanController(karyawanService)
	karyawanTrainingController := controller.NewKaryawanTrainingController(karyawanTrainingService)
	userController := controller.NewUserController(userService)
	authController := controller.NewAuthenticationController(authService)
	uploadController := controller.NewUploadController()

	// Create group routing endpoint "/api/v1"
	v1 := r.Group("/api/v1")
	{

		training := v1.Group("/training")
		{
			training.GET("/:id", trainingController.GetTraining)
			training.GET("/list", trainingController.GetAllTraining)
			training.POST("", trainingController.SaveTraining)
			training.PUT("", trainingController.UpdateTraining)
			training.DELETE("/:id", trainingController.DeleteTraining)
		}

		rekening := v1.Group("/rekening")
		{
			rekening.GET("/:id", rekeningController.GetRekening)
			rekening.GET("/list", rekeningController.GetAllRekening)
			rekening.POST("", rekeningController.SaveRekening)
			rekening.PUT("", rekeningController.UpdateRekening)
			rekening.DELETE("/:id", rekeningController.DeleteRekenig)
		}

		karyawan := v1.Group("/karyawan")
		{
			karyawan.GET("/:id", karyawanController.GetKaryawan)
			karyawan.GET("/list", karyawanController.GetAllKaryawan)
			karyawan.POST("", karyawanController.SaveKaryawan)
			karyawan.PUT("", karyawanController.UpdateKaryawan)
			karyawan.DELETE("/:id", karyawanController.DeleteKaryawan)
		}

		karyawanTraining := v1.Group("/karyawanTraining")
		{
			karyawanTraining.GET("/:id", karyawanTrainingController.GetKaryawanTraining)
			karyawanTraining.GET("/list", karyawanTrainingController.GetAllKaryawanTraining)
			karyawanTraining.POST("", karyawanTrainingController.SaveKaryawanTraining)
			karyawanTraining.PUT("", karyawanTrainingController.UpdateKaryawanTraining)
			karyawanTraining.DELETE("/:id", karyawanTrainingController.DeleteKaryawanTraining)
		}

		userRegister := v1.Group("/user-register")
		{
			userRegister.POST("", userController.RegisterUser)
			userRegister.POST("/register-confirm-otp/:otp", userController.ApprovedUser)
			userRegister.POST("/send-otp", userController.SendOtpRegister)
		}

		user := v1.Group("/forget-password")
		{
			user.POST("/change-password", userController.ResetPassword)
			user.POST("/send", userController.SendOtpForgetPassword)
		}

		userLogin := v1.Group("/user-login")
		{
			userLogin.POST("/login", authController.Login)
		}

		v1.GET("user/:id", userController.GetUser)

		v1.POST("/upload", uploadController.UploadFile)
		v1.GET("/showFile/:filename", uploadController.ShowFile)

	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
