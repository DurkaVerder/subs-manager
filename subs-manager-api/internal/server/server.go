package server

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	GetSubscriptionHandler(ctx *gin.Context)
	GetSubscriptionsHandler(ctx *gin.Context)
	GetTotalPriceByFiltersHandler(ctx *gin.Context)
	NewSubscriptionHandler(ctx *gin.Context)
	DeleteSubscriptionHandler(ctx *gin.Context)
	UpdateSubscriptionHandler(ctx *gin.Context)
}

type Server struct {
	handler Handler
	r       *gin.Engine
}

func NewServer(handler Handler, r *gin.Engine) *Server {
	return &Server{handler: handler, r: r}
}

func (s *Server) Start() {
	s.initRouter()

	if err := s.r.Run(os.Getenv("PORT")); err != nil {
		panic(err)
	}
}

func (s *Server) initRouter() {
	if s.r == nil {
		s.r = gin.Default()
	}

	s.r.GET("/subscription", s.handler.GetSubscriptionHandler)
	s.r.GET("/subscriptions", s.handler.GetSubscriptionsHandler)
	s.r.GET("/total-price", s.handler.GetTotalPriceByFiltersHandler)
	s.r.POST("/subscription", s.handler.NewSubscriptionHandler)
	s.r.DELETE("/subscription/:id", s.handler.DeleteSubscriptionHandler)
	s.r.PUT("/subscription", s.handler.UpdateSubscriptionHandler)
}
