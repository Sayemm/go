package user

import (
	"ecommerce/util"
	"encoding/json"
	"net/http"
)

type ReqLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var req ReqLogin
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&req)

	if err != nil {
		http.Error(w, "Give me valid json", http.StatusBadRequest)
		return
	}

	usr, err := h.svc.Find(req.Email, req.Password)

	if err != nil {
		util.SendError(w, http.StatusUnauthorized, "Unauthorized User")
		return
	}
	if usr == nil {
		util.SendError(w, http.StatusBadRequest, "Invalid Credential")
		return
	}

	accessToken, err := util.CreateJwt(h.cnf.JwtSecretKey, util.Payload{
		Sub:         usr.ID,
		FirstName:   usr.FirstName,
		LastName:    usr.LastName,
		Email:       usr.Email,
		IsShopOwner: usr.IsShopOwner,
	})
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	util.SendData(w, http.StatusCreated, accessToken)
}
