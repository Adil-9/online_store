package servicelayer

import (
	dblayer "github.com/Adil-9/online_store/internal/DBlayer"
)

type ServiceHandler struct {
	Authorization
}

func NewSVH(db *dblayer.Query) *ServiceHandler {
	return &ServiceHandler{
		Authorization: newAuthService(db.Authorization),
	}
}
