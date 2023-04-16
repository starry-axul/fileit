package file

import (
	"context"
)

type (
	GetAllReq struct {
	}

	CreateReq struct {
		Contect  string
		Type     string
		Private  string
		FileName string
		Token    string
	}

	GetReq struct {
		Token string
	}

	Controller func(ctx context.Context, request interface{}) (interface{}, error)

	// Endpoints struct
	Endpoints struct {
		GetAll Controller
	}
)
