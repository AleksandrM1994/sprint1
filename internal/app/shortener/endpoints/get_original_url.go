package endpoints

import (
	"net/http"
)

func (c *Controller) GetOriginalURLHandler(res http.ResponseWriter, req *http.Request) {
	id := req.URL.Path[len("/"):]
	originalURL, err := c.service.GetOriginalURL(id)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
	}
	if originalURL != "" {
		res.Header().Add("Location", originalURL)
		res.WriteHeader(http.StatusTemporaryRedirect)
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
