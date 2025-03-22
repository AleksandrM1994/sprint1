package endpoints

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	custom_errs "github.com/sprint1/internal/app/shortener/errors"
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
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}

	getShortenURLRequest := &GetShortenURLRequest{}
	errUnmarshal := json.Unmarshal(request, getShortenURLRequest)
	if errUnmarshal != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errUnmarshal.Error(), http.StatusInternalServerError)
		return
	}

	if getShortenURLRequest.URL != "" {
		shortURL, errSaveURL := c.service.SaveURL(getShortenURLRequest.URL)
		switch {
		case errors.Is(errSaveURL, custom_errs.ErrUniqueViolation) && shortURL != "":
			res.Header().Set("Content-Type", "application/json")
			getShortenURLResponse := GetShortenURLResponse{
				Result: c.cfg.BaseShortURL + "/" + shortURL,
			}
			body, errMarshal := json.Marshal(getShortenURLResponse)
			if errMarshal != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, errMarshal.Error(), http.StatusInternalServerError)
				return
			}
			res.WriteHeader(http.StatusConflict)
			_, errWrite := res.Write(body)
			if errWrite != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, errWrite.Error(), http.StatusInternalServerError)
			}
		case errSaveURL != nil && shortURL == "":
			res.WriteHeader(http.StatusInternalServerError)
			http.Error(res, errSaveURL.Error(), http.StatusInternalServerError)
		case shortURL != "" && errSaveURL == nil:
			res.Header().Set("Content-Type", "application/json")
			getShortenURLResponse := GetShortenURLResponse{
				Result: c.cfg.BaseShortURL + "/" + shortURL,
			}
			body, errMarshal := json.Marshal(getShortenURLResponse)
			if errMarshal != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, errMarshal.Error(), http.StatusInternalServerError)
				return
			}
			res.WriteHeader(http.StatusCreated)
			_, errWrite := res.Write(body)
			if errWrite != nil {
				res.WriteHeader(http.StatusInternalServerError)
				http.Error(res, errWrite.Error(), http.StatusInternalServerError)
				return
			}
		case shortURL == "":
			res.WriteHeader(http.StatusBadRequest)
		default:
			res.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		res.WriteHeader(http.StatusBadRequest)
	}
}
