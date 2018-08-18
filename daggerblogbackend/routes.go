/* Copyright Fengchang Xie 2018 All Rights Reserved */
package main

/**
 * Created by Fengchang on 2018/5/30 12:53.
 */

func initializeRoutes() {

	router.POST("/login", AuthNotRequired(), login)
	router.POST("/logout", AuthRequired(), logout)
	router.POST("/api/v1/image", AuthRequired(), UploadImage)
	router.POST("/api/v1/post", AuthRequired(), CreateNewBlog)
	router.POST("/public/comment-2min-window", OpenCommentWindow4Blog)
	router.POST("/public/comment-stuff", CommentStuff)

	router.PUT("/api/v1/trashblog/:id", AuthRequired(), TrashBlog)
	router.PUT("/api/v1/post/:id", AuthRequired(), EditBlog)
	router.PUT("/api/v1/comment/readcomment/:id", AuthRequired(), ClearUnreadComment)

	router.GET("/public/blog", GetAllBlogsByPage)
	router.GET("/public/post/:id", GetPostDetail)
	router.GET("/public/comment4blog/:id", GetComment4Blog)
	router.GET("/api/v1/user/whoami", AuthRequired(), whoami)
	router.GET("/api/v1/category", AuthRequired(), GetAllCategories)
	router.GET("/api/v1/comment/unreadcnt", AuthRequired(), CountUnreadComment)
	router.GET("/api/v1/comment/list", AuthRequired(), GetCommentList)

	router.DELETE("/api/v1/comment/:id", AuthRequired(), DeleteComment)

}
