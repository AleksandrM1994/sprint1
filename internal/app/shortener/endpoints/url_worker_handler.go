package endpoints

import (
	"io"
	"net/http"
)

func (c *Controller) UrlWorkerHandlerHandler(res http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		request, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		url := string(request)
		shortUrl := c.service.SaveURL(url)
		res.WriteHeader(http.StatusCreated)
		res.Header().Set("Content-Type", "text/plain")
		_, _ = res.Write([]byte(shortUrl))
	case http.MethodGet:
		var id string
		for k, v := range req.URL.Query() {
			if k == "id" {
				id = v[0]
			}
		}

		originalURL := c.service.GetOriginalURL(id)
		if originalURL != "" {
			res.Header().Set("Location", originalURL)
			res.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}
