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
		panic(err)
	}

	saveURLsBatchRequest := &SaveURLsBatchRequest{}
	errUnmarshal := json.Unmarshal(request, saveURLsBatchRequest)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}

	urls := make([]URLInBatch, 0, len(saveURLsBatchRequest.URLs))
	if len(saveURLsBatchRequest.URLs) != 0 {
		for _, url := range saveURLsBatchRequest.URLs {
			shortUrl, errSaveURL := c.service.SaveURL(url.OriginalURL)
			if errSaveURL != nil {
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
			panic(errMarshal)
		}
		_, errWrite := res.Write(body)
		if errWrite != nil {
			panic(errWrite)
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
