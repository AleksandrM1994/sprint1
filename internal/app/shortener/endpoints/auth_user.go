package endpoints

import (
	"encoding/json"
	"net/http"
)

type AuthUserRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (c *Controller) AuthUser(res http.ResponseWriter, req *http.Request) {
	var request *AuthUserRequest
	errDecode := json.NewDecoder(req.Body).Decode(&request)
	if errDecode != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errDecode.Error(), http.StatusInternalServerError)
		return
	}

	user, errGetUserInfo := c.service.AuthenticateUser(request.Login, request.Password)
	if errGetUserInfo != nil {
		makeEndpointError(res, errGetUserInfo)
		return
	}

	http.SetCookie(res, &http.Cookie{
		Name:    c.cfg.AuthUserCookieName,
		Value:   user.Cookie,
		Path:    "/",
		Secure:  false,
		Expires: *user.CookieFinish,
	})

	res.WriteHeader(http.StatusOK)
	_, errWrite := res.Write([]byte("Забирай куку"))
	if errWrite != nil {
		res.WriteHeader(http.StatusInternalServerError)
		http.Error(res, errWrite.Error(), http.StatusInternalServerError)
		return
	}
}
