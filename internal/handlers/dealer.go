package handlers

import (
	"github.com/Victor90001/prod/internal/requests"
	"github.com/Victor90001/prod/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type DealerHandlers struct {
	svc    *services.DealerService
	engine *gin.Engine
}

func NewDealerHandlers(engine *gin.Engine, svc *services.DealerService) (*DealerHandlers, error) {
	h := &DealerHandlers{
		svc:    svc,
		engine: engine,
	}
	h.initRoute()
	return h, nil
}

func (h *DealerHandlers) initRoute() {
	h.engine.GET("/dealer", h.GetDealers)           //
	h.engine.POST("/dealer/delete", h.DeleteDealer) //
	h.engine.PUT("/dealer", h.InsertDealer)         //
	h.engine.POST("/dealer", h.UpdateDealer)        //
}

func (h *DealerHandlers) GetDealers(c *gin.Context) {
	//req, ok := GetRequest[requests.GetDealersRequest](c)
	//if !ok {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": "get dealers request error", "text": ok})
	//	return
	//}

	dealers, err := h.svc.GetDealers()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "get dealers error", "text": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "ok", "item": dealers})
}

func (h *DealerHandlers) DeleteDealer(c *gin.Context) {

	req, ok := GetRequest[requests.DeleteDealerRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete dealer request error", "text": ok})
		return
	}

	if err := h.svc.DeleteDealer(req.Id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "delete dealer error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *DealerHandlers) InsertDealer(c *gin.Context) {

	req, ok := GetRequest[requests.InsertDealerRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert dealer request error", "text": ok})
		return
	}

	if err := h.svc.InsertDealer(req.NetworkId, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "insert dealer error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func (h *DealerHandlers) UpdateDealer(c *gin.Context) {

	req, ok := GetRequest[requests.UpdateDealerRequest](c)
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update dealer request error", "text": ok})
		return
	}

	if err := h.svc.UpdateDealer(req.Id, req.Name); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "update dealer error", "text": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
