package server

import (
	"net/http"

	"github.com/Jhnvlglmlbrt/crud-gin/internal/controller"
	"github.com/Jhnvlglmlbrt/crud-gin/utils"
	"github.com/gin-gonic/gin"
)

func NewRouter(tagsController *controller.TagsController) *gin.Engine {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"localhost"})

	router.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "welcome home")
	})

	baseRouter := router.Group("/api")
	tagsRouter := baseRouter.Group("/tags")
	tagsRouter.GET("", utils.ErrorHandler(tagsController.FindAll))
	tagsRouter.GET("/:tagID", utils.ErrorHandler(tagsController.FindByID))
	tagsRouter.POST("", utils.ErrorHandler(tagsController.Create))
	tagsRouter.PATCH("/:tagID", utils.ErrorHandler(tagsController.Update))
	tagsRouter.DELETE("/:tagID", utils.ErrorHandler(tagsController.Delete))

	return router
}

func NewServer(tagsController *controller.TagsController) *Server {
	router := NewRouter(tagsController)

	return &Server{
		Addr:    ":8080",
		Handler: router,
	}
}

type Server struct {
	Addr    string
	Handler *gin.Engine
}

func (s *Server) Start() error {
	return s.Handler.Run(s.Addr)
}
