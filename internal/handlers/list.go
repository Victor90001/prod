package handlers

import (
	"github.com/Victor90001/prod/internal/requests"
	"github.com/Victor90001/prod/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListHandlers struct {
	svc    *services.ListService
	engine *gin.Engine
}

func NewListHandlers(engine *gin.Engine, svc *services.ListService) (*ListHandlers, error) {
	h := &ListHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *ListHandlers) initRoute() {
	h.engine.GET("/list", h.GetLists)           //
	h.engine.POST("/list/delete", h.DeleteList) //
	h.engine.PUT("/list", h.InsertList)         //
	h.engine.POST("/list", h.UpdateList)        //
}

func (h *ListHandlers) GetLists(c *gin.Context) {
	lists, err := h.svc.GetLists()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get lists error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "item": lists})
}

func (h *ListHandlers) DeleteList(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteListRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete list request error", "text": ok})
		return
	}

	if err := h.svc.DeleteList(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete list error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *ListHandlers) InsertList(c *gin.Context) {

	req, ok := GetRequest[requests.InsertListRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert list request error", "text": ok})
		return
	}
	req.Id = 0
	if err := h.svc.InsertList(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert list error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *ListHandlers) UpdateList(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateListRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update list request error", "text": ok})
		return
	}
	req.DealerId = 0
	if err := h.svc.UpdateList(req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update list error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
