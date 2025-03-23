package endpoints

import "net/http"

func (c *Controller) PingHandler(res http.ResponseWriter, req *http.Request) {
	err := c.service.Ping()
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusOK)
}
