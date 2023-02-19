package routes

import (
	"github.com/J-Nokwal/ihar_backend/pkg/controllers"
	"github.com/gin-gonic/gin"
)

var RegisterBookStoreRoutes = func(router *gin.Engine) {
	//user apis
	router.POST("/user", controllers.CreateUser)
	router.GET("/user/:id", controllers.GetUser)
	router.PATCH("/user", controllers.PatchUser)

	router.POST("/post", controllers.CreatePost)
	router.GET("/post/:id/:byUser", controllers.GetPostForUser)
	router.GET("/post/all/OfUser/:ofUser/:byUser", controllers.GetAllPostOfUserByUserId)
	router.GET("/post/all/:byUser", controllers.GetAllPostByUserId)
	router.GET("/post/byPage/:pageId/:byUser", controllers.GetPostByPageIdByUser)

	router.PATCH("/post", controllers.PatchPost)
	router.DELETE("/post/:id", controllers.DeletePost)

	router.POST("/comment", controllers.CreateComment)
	router.GET("/comment/all/:postId", controllers.GetAllCommentFromPost)
	router.DELETE("/comment/:commentId", controllers.DeleteComment)
	router.PATCH("/comment", controllers.PatchComment)

	router.POST("/like", controllers.TriggerLike)
	router.GET("/like/byPostId/:postId", controllers.GetUsersByPostLikes)

	router.GET("/search/:searchQuery/:byUser", controllers.GetSearchQueryResults)

	router.GET("/isServerOnline", controllers.IsServerOnline)

	router.POST("/report", controllers.ReportUserOrPost)
}
