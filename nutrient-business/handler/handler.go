package handler

import (
	"context"
	"fmt"

	"github.com/pkg/errors"
	"github.com/suhdev/nutrient/proton"
)

// BusinessHandler implements the business service handler
type BusinessHandler struct {
	ds  proton.DataStoreService
	api ExternalAPI
}

// NewBusinessHandler creates a new BusinessHandler instance given the data store service
func NewBusinessHandler(ds proton.DataStoreService, api ExternalAPI) *BusinessHandler {
	return &BusinessHandler{
		ds:  ds,
		api: api,
	}
}

// Search performs search on the database first using the underlying datastore service, if no results
// were found, the search will issue a request to third party API, and will store the returned results
// using the data store service
func (b *BusinessHandler) Search(ctx context.Context, req *proton.QueryRequest, res *proton.Result) error {
	resp, err := b.ds.Search(ctx, req)
	fmt.Println(err)
	if err != nil || len(resp.Entries) == 0 {
		resp, err = b.api.Search(ctx, req)
		if err != nil {
			return errors.Wrap(err, "unable to query third party API")
		}
		res.Entries = resp.Entries
		if len(resp.Entries) > 0 {
			_, err := b.ds.Add(ctx, resp)
			if err != nil {
				return errors.Wrap(err, "failed to add data into the database")
			}
		}
	}
	return errors.Wrap(err, "failed to perform search")
}
