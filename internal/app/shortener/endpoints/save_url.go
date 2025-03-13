package endpoints

import (
	"io"
	"net/http"
)

func (c *Controller) SaveURLHandler(res http.ResponseWriter, req *http.Request) {
	request, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	url := string(request)
	if url != "" {
		shortUrl, errSaveURL := c.service.SaveURL(url)
		if errSaveURL != nil {
			http.Error(res, errSaveURL.Error(), http.StatusInternalServerError)
		}
		if shortUrl != "" {
			res.WriteHeader(http.StatusCreated)
			res.Header().Set("Content-Type", "text/plain")
			_, writeErr := res.Write([]byte(c.cfg.BaseShortURL + "/" + shortUrl))
			if writeErr != nil {
				res.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
