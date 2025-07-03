package public

import (
	"context"
	"encoding/json"
	"net/http"
	"time"
)

// CreateUserRequest запрос для ручки по созданию пользователя
type CreateUserRequest struct {
	Login string `json:"login"`
}

// CreateUserResponse ответ для ручки по созданию пользователя
type CreateUserResponse struct {
	Password string `json:"password"`
}

// CreateUser ручка по созданию пользователя
func (c *Controller) CreateUser(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), 30*time.Second)
	defer cancel()

	var request *CreateUserRequest
	errDecode := json.NewDecoder(req.Body).Decode(&request)
	if errDecode != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errDecode.Error(), http.StatusInternalServerError)
		return
	}

	pass, errCreateUser := c.service.CreateUser(ctx, request.Login)
	if errCreateUser != nil {
		makeEndpointError(res, errCreateUser)
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	response := &CreateUserResponse{
		Password: pass,
	}
	body, errMarshal := json.Marshal(response)
	if errMarshal != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errMarshal.Error(), http.StatusInternalServerError)
		return
	}
	_, errWrite := res.Write(body)
	if errWrite != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errWrite.Error(), http.StatusInternalServerError)
		return
	}
}
