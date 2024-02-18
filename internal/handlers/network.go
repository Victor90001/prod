package handlers

import (
	"github.com/Victor90001/prod/internal/requests"
	"github.com/Victor90001/prod/internal/services"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

type NetworkHandlers struct {
	svc    *services.NetworkService
	engine *gin.Engine
}

func NewNetworkHandlers(engine *gin.Engine, svc *services.NetworkService) (*NetworkHandlers, error) {
	h := &NetworkHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *NetworkHandlers) initRoute() {
	h.engine.GET("/network", h.GetNetworks)
	h.engine.POST("/network/delete", h.DeleteNetwork)
	h.engine.PUT("/network", h.InsertNetwork)
	h.engine.POST("/network", h.UpdateNetwork)
}

func (h *NetworkHandlers) GetNetworks(c *gin.Context) {
	networks, err := h.svc.GetNetworks()
	logrus.Debug(networks)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get networks error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "item": networks})
}

func (h *NetworkHandlers) DeleteNetwork(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteNetworkRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete network request error", "text": ok})
		return
	}

	if err := h.svc.DeleteNetwork(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete network error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *NetworkHandlers) InsertNetwork(c *gin.Context) {

	req, ok := GetRequest[requests.InsertNetworkRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert network request error", "text": ok})
		return
	}

	if err := h.svc.InsertNetwork(req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert network error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *NetworkHandlers) UpdateNetwork(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateNetworkRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update network request error", "text": ok})
		return
	}

	if err := h.svc.UpdateNetwork(req.Id, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update network error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
