package handlers

import (
	"fmt"
	"net/http"
	"subs-manager-api/internal/models"

	"github.com/gin-gonic/gin"
)

const (
	ServiceNameQuery = "service-name"
	UserIDQuery      = "user-id"
)

type SubscribeService interface {
	GetSubscriptionFilter(data models.DataFilter) ([]models.ServiceSubscriptionResponse, error)
	GetSubscription(serviceName, userID string) (models.ServiceSubscriptionResponse, error)
	GetSubscriptions(userID string) ([]models.ServiceSubscriptionResponse, error)
	GetTotalPriceByFilters(data models.DataFilter) (models.TotalPriceResponse, error)
	CreateSubscription(subscription models.ServiceSubscriptionRequest) error
	UpdateSubscription(subscription models.ServiceSubscriptionRequest) error
	DeleteSubscription(serviceName, userID string) error
}

type Handler struct {
	subscribeService SubscribeService
}

func NewHandler(subscribeService SubscribeService) *Handler {
	return &Handler{subscribeService: subscribeService}
}

func (h *Handler) GetSubscriptionHandler(ctx *gin.Context) {
	serviceName := ctx.Query(ServiceNameQuery)
	userID := ctx.Query(UserIDQuery)

	if serviceName == "" || userID == "" {
		ctx.JSON(http.StatusBadRequest, fmt.Sprintf("serviceName = %s, userID = %s", serviceName, userID))
	}

	response, err := h.subscribeService.GetSubscription(serviceName, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"subscribe": response,
	})

}

func (h *Handler) GetSubscriptionsHandler(ctx *gin.Context) {
	userID := ctx.Query(UserIDQuery)

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, "invalid userID")
	}

	response, err := h.subscribeService.GetSubscriptions(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{
		"subscribes": response,
	})

}
