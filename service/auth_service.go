package service

import (
	"h8-movies/entity"
	"h8-movies/pkg/errs"
	"h8-movies/pkg/helpers"
	"h8-movies/repository/comment_repository"
	"h8-movies/repository/photo_repository"
	"h8-movies/repository/social_media_repository"
	"h8-movies/repository/user_repository"

	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Authentication() gin.HandlerFunc
	Authorization() gin.HandlerFunc
	AuthorizationComment() gin.HandlerFunc
	AuthorizationSocialMedia() gin.HandlerFunc
}

type authService struct {
	userRepo  user_repository.UserRepository
	photoRepo photo_repository.PhotoRepository
	commentRepo comment_repository.CommentRepository
	socialmediaRepo social_media_repository.SocialMediaRepository
}

func NewAuthService(userRepo user_repository.UserRepository, photoRepo photo_repository.PhotoRepository, commentRepo comment_repository.CommentRepository, socialmediaRepo social_media_repository.SocialMediaRepository) AuthService {
	return &authService{
		userRepo:  userRepo,
		photoRepo: photoRepo,
		commentRepo: commentRepo,
		socialmediaRepo: socialmediaRepo,
	}
}

func (a *authService) Authorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		

		photoId, err := helpers.GetParamId(ctx, "photoId")


		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		photo, err := a.photoRepo.GetPhotoById(photoId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		// if user.Level == entity.Admin {
		// 	ctx.Next()
		// 	return
		// }

		if photo.UserId != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the photo data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		ctx.Next()
	
	}
}


func (a *authService) AuthorizationComment() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		

		commentId, err := helpers.GetParamId(ctx, "commentId")


		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		comment, err := a.commentRepo.GetCommentById(commentId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		// if user.Level == entity.Admin {
		// 	ctx.Next()
		// 	return
		// }

		if comment.UserId != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the comment data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		ctx.Next()
	
	}
}



func (a *authService) AuthorizationSocialMedia() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		user := ctx.MustGet("userData").(entity.User)

		

		socialmediaId, err := helpers.GetParamId(ctx, "socialmediaId")


		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		socialmedia, err := a.socialmediaRepo.GetSocialMediaById(socialmediaId)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		// if user.Level == entity.Admin {
		// 	ctx.Next()
		// 	return
		// }

		if socialmedia.UserId != user.Id {
			unauthorizedErr := errs.NewUnauthorizedError("you are not authorized to modify the social media data")
			ctx.AbortWithStatusJSON(unauthorizedErr.Status(), unauthorizedErr)
			return
		}

		ctx.Next()
	
	}
}



func (a *authService) Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var invalidTokenErr = errs.NewUnauthenticatedError("invalid token")
		bearerToken := ctx.GetHeader("Authorization")  //untuk mendapatkan key authorization di postman

		var user entity.User // User{Id:0, Email: ""}

		err := user.ValidateToken(bearerToken)

		if err != nil {
			ctx.AbortWithStatusJSON(err.Status(), err)
			return
		}

		result, err := a.userRepo.GetUserByEmail(user.Email)

		if err != nil {
			ctx.AbortWithStatusJSON(invalidTokenErr.Status(), invalidTokenErr)
			return
		}

		_ = result

		ctx.Set("userData", user)

		ctx.Next()
	}
}
