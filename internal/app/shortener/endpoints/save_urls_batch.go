package endpoints

import (
	"encoding/json"
	"io"
	"net/http"
)

type SaveURLsBatchRequest struct {
	URLs []URLInBatch `json:"urls"`
}

type SaveURLsBatchResponse struct {
	URLs []URLInBatch `json:"urls"`
}

type URLInBatch struct {
	CorrelationId string `json:"correlation_id"`
	OriginalURL   string `json:"original_url,omitempty"`
	ShortURL      string `json:"short_url,omitempty"`
}

func (c *Controller) SaveURLsBatch(res http.ResponseWriter, req *http.Request) {
	request, err := io.ReadAll(req.Body)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, err.Error(), http.StatusInternalServerError)
	}

	saveURLsBatchRequest := &SaveURLsBatchRequest{}
	errUnmarshal := json.Unmarshal(request, saveURLsBatchRequest)
	if errUnmarshal != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errUnmarshal.Error(), http.StatusInternalServerError)
	}

	urls := make([]URLInBatch, 0, len(saveURLsBatchRequest.URLs))
	if len(saveURLsBatchRequest.URLs) != 0 {
		for _, url := range saveURLsBatchRequest.URLs {
			shortUrl, errSaveURL := c.service.SaveURL(url.OriginalURL)
			if errSaveURL != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, errSaveURL.Error(), http.StatusInternalServerError)
			}
			if shortUrl != "" {
				url.OriginalURL = ""
				url.ShortURL = shortUrl
				urls = append(urls, url)
			} else {
				res.WriteHeader(http.StatusBadRequest)
			}
		}
		res.WriteHeader(http.StatusCreated)
		res.Header().Add("Content-Type", "application/json")
		saveURLsBatchResponse := SaveURLsBatchResponse{
			URLs: urls,
		}
		body, errMarshal := json.Marshal(saveURLsBatchResponse)
		if errMarshal != nil {
			res.WriteHeader(http.StatusInternalServerError)
			http.Error(res, errMarshal.Error(), http.StatusInternalServerError)
		}
		_, errWrite := res.Write(body)
		if errWrite != nil {
			res.WriteHeader(http.StatusInternalServerError)
			http.Error(res, errWrite.Error(), http.StatusInternalServerError)
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
