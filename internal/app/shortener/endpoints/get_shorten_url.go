package endpoints

import (
	"encoding/json"
	"io"
	"net/http"
)

type GetShortenURLRequest struct {
	URL string `json:"url"`
}

type GetShortenURLResponse struct {
	Result string `json:"result"`
}

func (c *Controller) GetShortenURLHandler(res http.ResponseWriter, req *http.Request) {
	request, err := io.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	getShortenURLRequest := &GetShortenURLRequest{}
	errUnmarshal := json.Unmarshal(request, getShortenURLRequest)
	if errUnmarshal != nil {
		panic(errUnmarshal)
	}

	if getShortenURLRequest.URL != "" {
		shortUrl := c.service.SaveURL(getShortenURLRequest.URL)
		if shortUrl != "" {
			res.WriteHeader(http.StatusCreated)
			res.Header().Add("Content-Type", "application/json")
			getShortenURLResponse := GetShortenURLResponse{
				Result: c.cfg.BaseShortURL + "/" + shortUrl,
			}
			body, errMarshal := json.Marshal(getShortenURLResponse)
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
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
