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
	DeleteSubscription(ID string) error
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
		return
	}

	response, err := h.subscribeService.GetSubscription(serviceName, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"subscribe": response,
	})

}

func (h *Handler) GetSubscriptionsHandler(ctx *gin.Context) {
	userID := ctx.Query(UserIDQuery)

	if userID == "" {
		ctx.JSON(http.StatusBadRequest, "invalid userID")
		return
	}

	response, err := h.subscribeService.GetSubscriptions(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"subscribes": response,
	})

}

func (h *Handler) NewSubscriptionHandler(ctx *gin.Context) {
	var subscription models.ServiceSubscriptionRequest
	if err := ctx.ShouldBindJSON(&subscription); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.subscribeService.CreateSubscription(subscription); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "created"})
}

func (h *Handler) UpdateSubscriptionHandler(ctx *gin.Context) {
	var subscription models.ServiceSubscriptionRequest
	if err := ctx.ShouldBindJSON(&subscription); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if err := h.subscribeService.UpdateSubscription(subscription); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "updated"})
}

func (h *Handler) DeleteSubscriptionHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, "invalid ID")
		return
	}

	if err := h.subscribeService.DeleteSubscription(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "deleted"})
}

func (h *Handler) GetTotalPriceByFiltersHandler(ctx *gin.Context) {
	var data models.DataFilter
	if err := ctx.ShouldBindQuery(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	response, err := h.subscribeService.GetTotalPriceByFilters(data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, response)
}

