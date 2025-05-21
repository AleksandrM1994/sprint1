package endpoints

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sprint1/internal/app/shortener/endpoints/middleware"
	"github.com/sprint1/internal/app/shortener/service"
)

// GetUserURLsResponse ответ по ручке по получению урлов пользователя
type GetUserURLsResponse struct {
	OriginalURL string `json:"original_url,omitempty"`
	ShortURL    string `json:"short_url,omitempty"`
}

// GetUserURLs ручка по получению урлов пользователя
func (c *Controller) GetUserURLs(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	userIDValue := ctx.Value(middleware.UserID)
	userID, ok := userIDValue.(string)
	if !ok {
		userID = ""
	}

	urls, errGetUserURLs := c.service.GetUserURLs(ctx, userID)
	if errGetUserURLs != nil {
		c.lg.Infow("GetUserURLs error", "error", errGetUserURLs)
		makeEndpointError(res, errGetUserURLs)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusOK)
	response := mappingGetUserURLsResponse(urls)
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

// mappingGetUserURLsResponse маппинг ответа по ручке по получению урлов пользователя
func mappingGetUserURLsResponse(urls []*service.UserURLs) []*GetUserURLsResponse {
	res := make([]*GetUserURLsResponse, 0)
	for _, url := range urls {
		res = append(res, &GetUserURLsResponse{
			OriginalURL: url.OriginalURL,
			ShortURL:    url.ShortURL,
		})
	}
	return res
}
