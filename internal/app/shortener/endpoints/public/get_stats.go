package public

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

type GetStatsResponse struct {
	URLs  uint32 `json:"urls"`
	Users uint32 `json:"users"`
}

// GetStatsHandler ручка получению статистики
func (c *Controller) GetStatsHandler(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	resGetStats, errGetStats := c.service.GetStats(ctx)
	if errGetStats != nil {
		c.lg.Infow("GetStats error", "error", errGetStats)
		makeEndpointError(res, errGetStats)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	response := GetStatsResponse{
		URLs:  resGetStats.URLs,
		Users: resGetStats.Users,
	}
	body, errMarshal := json.Marshal(response)
	if errMarshal != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errMarshal.Error(), http.StatusInternalServerError)
		return
	}
	_, errWrite := res.Write(body)
	if errWrite != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errWrite.Error(), http.StatusInternalServerError)
		return
	}
}
