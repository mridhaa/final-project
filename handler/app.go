package handler

import (
	"h8-movies/database"
	"h8-movies/repository/comment_repository/comment_pg"
	"h8-movies/repository/photo_repository/photo_pg"
	"h8-movies/repository/social_media_repository/social_media_pg"
	"h8-movies/repository/user_repository/user_pg"
	"h8-movies/service"

	"github.com/gin-gonic/gin"
)

func StartApp() {
	var port = "8080"
	database.InitiliazeDatabase()

	db := database.GetDatabaseInstance()

	socialmediaRepo := social_media_pg.NewSocialMediaPG(db)

	socialmediaService := service.NewSocialMediaService(socialmediaRepo)

	socialmediaHandler := NewSocialMediaHandler(socialmediaService)

	commentRepo := comment_pg.NewCommentPG(db)

	commentService := service.NewCommentService(commentRepo)

	commentHandler := NewCommentHandler(commentService)

	photoRepo := photo_pg.NewPhotoPG(db)

	photoService := service.NewPhotoService(photoRepo)

	photoHandler := NewPhotoHandler(photoService)

	userRepo := user_pg.NewUserPG(db)

	userService := service.NewUserService(userRepo)

	userHandler := NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, photoRepo, commentRepo, socialmediaRepo)

	route := gin.Default()

	userRoute := route.Group("/users")
	{
		userRoute.POST("/login", userHandler.Login)
		userRoute.POST("/register", userHandler.Register)
	}

	photoRoute := route.Group("/photos")
	{
		photoRoute.POST("/", authService.Authentication(), photoHandler.CreatePhoto)

		photoRoute.PUT("/:photoId", authService.Authentication(), authService.Authorization(), photoHandler.UpdatePhotoById)
		photoRoute.GET("/:photoId", authService.Authentication(), authService.Authorization(), photoHandler.GetPhotoById)
		photoRoute.DELETE("/:photoId", authService.Authentication(), authService.Authorization(), photoHandler.DeletePhotoById)
		photoRoute.GET("/", authService.Authentication(), photoHandler.GetPhotoAll)
	}

	commentRoute := route.Group("/comments")
	{
		commentRoute.POST("/:photoId", authService.Authentication(), commentHandler.CreateComment)

		commentRoute.PUT("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.UpdateCommentById)
		commentRoute.GET("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.GetCommentById)
		commentRoute.DELETE("/:commentId", authService.Authentication(), authService.AuthorizationComment(), commentHandler.DeleteCommentById)
		commentRoute.GET("/", authService.Authentication(), commentHandler.GetCommentAll)
	}


		socialmediaRoute := route.Group("/socialmedias")
	{
		socialmediaRoute.POST("/", authService.Authentication(), socialmediaHandler.CreateSocialMedia)

		socialmediaRoute.PUT("/:socialmediaId", authService.Authentication(), authService.AuthorizationSocialMedia(), socialmediaHandler.UpdateSocialMediaById)
		socialmediaRoute.GET("/:socialmediaId", authService.Authentication(), authService.AuthorizationSocialMedia(), socialmediaHandler.GetSocialMediaById)
		socialmediaRoute.DELETE("/:socialmediaId", authService.Authentication(), authService.AuthorizationSocialMedia(), socialmediaHandler.DeleteSocialMediaById)
		socialmediaRoute.GET("/", authService.Authentication(), socialmediaHandler.GetSocialMediaAll)
	}


	

	route.Run(":" + port)
}
