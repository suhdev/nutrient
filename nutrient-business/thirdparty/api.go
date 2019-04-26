package thirdparty

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pkg/errors"
	"github.com/suhdev/nutrient/proton"
)

const apiURL = "https://api.edamam.com/search?q=%s&app_id=%s&app_key=%s"

// EdmamAPI implements ExternalAPI interface for Edmam recipe search API
type EdmamAPI struct {
	appID  string
	appKey string
}

// NewEdmamAPI creates instantiates
func NewEdmamAPI(id, key string) *EdmamAPI {
	return &EdmamAPI{
		appID:  id,
		appKey: key,
	}
}

// Search issues a search request to Edmam's recipe search API
func (e *EdmamAPI) Search(ctx context.Context, req *proton.QueryRequest) (*proton.Result, error) {
	resp, err := http.Get(fmt.Sprintf(apiURL, req.Q, e.appID, e.appKey))
	if resp != nil && resp.Body != nil {
		defer resp.Body.Close()
	}
	if err != nil {
		return nil, errors.Wrap(err, "failed to request data from edmam's API")
	}
	buf := bytes.NewBuffer([]byte{})
	_, err = io.Copy(buf, resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read Edmam's API resposne")
	}
	result := &result{}
	err = json.Unmarshal(buf.Bytes(), &result)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to unmarshal Edmam's API json response: %s", string(buf.Bytes()))
	}
	presult := &proton.Result{
		Entries: make([]*proton.Recipe, 0),
	}
	if result.Hits != nil && len(result.Hits) > 0 {
		for _, h := range result.Hits {
			if h.Recipe == nil {
				continue
			}
			presult.Entries = append(presult.Entries, &proton.Recipe{
				Calories:    fmt.Sprintf("%f", h.Recipe.Calories),
				Label:       h.Recipe.Label,
				Yield:       h.Recipe.Yield,
				TotalTime:   h.Recipe.TotalTime,
				TotalWeight: h.Recipe.TotalWeight,
				Uri:         h.Recipe.URI,
				Url:         h.Recipe.URL,
			})
		}
	}
	return presult, nil
}
