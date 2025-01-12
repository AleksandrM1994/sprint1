package endpoints

import (
	"io"
	"net/http"
)

func (c *Controller) SaveURLHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		request, err := io.ReadAll(req.Body)
		if err != nil {
			panic(err)
		}

		url := string(request)
		if url == "" {
			res.WriteHeader(http.StatusBadRequest)
		}
		shortUrl := c.service.SaveURL(url)
		res.WriteHeader(http.StatusCreated)
		res.Header().Set("Content-Type", "text/plain; charset=utf-8")
		_, writeErr := res.Write([]byte(shortUrl))
		if writeErr != nil {
			res.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}
