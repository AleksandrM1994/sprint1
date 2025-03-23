package endpoints

import (
	"encoding/json"
	"net/http"
)

type CreateUserRequest struct {
	Login string `json:"login"`
}

type CreateUserResponse struct {
	Password string `json:"password"`
}

func (c *Controller) CreateUser(res http.ResponseWriter, req *http.Request) {
	var request *CreateUserRequest
	errDecode := json.NewDecoder(req.Body).Decode(&request)
	if errDecode != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errDecode.Error(), http.StatusInternalServerError)
		return
	}

	pass, errCreateUser := c.service.CreateUser(request.Login)
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
