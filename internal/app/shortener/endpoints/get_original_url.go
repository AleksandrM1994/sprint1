package endpoints

import (
	"net/http"
)

func (c *Controller) GetOriginalURLHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		id := req.URL.Path[len("/url/"):]

		originalURL := c.service.GetOriginalURL(id)
		if originalURL != "" {
			res.Header().Add("Location", originalURL)
			res.WriteHeader(http.StatusTemporaryRedirect)
		} else {
			res.WriteHeader(http.StatusBadRequest)
		}
	} else {
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}
