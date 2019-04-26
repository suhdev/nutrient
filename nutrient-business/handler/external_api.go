package handler

import (
	"context"

	"github.com/suhdev/nutrient/proton"
)

type ExternalAPI interface {
	Search(context.Context, *proton.QueryRequest) (*proton.Result, error)
}
