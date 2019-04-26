package handler

import (
	"context"
	"fmt"
	"sync"

	"github.com/gocraft/dbr"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/pkg/errors"
	"github.com/suhdev/nutrient-datastore/db"
	"github.com/suhdev/nutrient/proton"
)

// DataStore implements the datastore handler interface
type DataStore struct {
	conn *dbr.Connection
	mu   sync.Mutex
}

// NewDataStore creates a new data store
func NewDataStore(conn *dbr.Connection) *DataStore {
	return &DataStore{
		conn: conn,
	}
}

// Search searches the sql database for the given query
// and returns the result if any results are found
// alternatively, it returns an error
func (d *DataStore) Search(ctx context.Context, req *proton.QueryRequest, res *proton.Result) error {
	sess := d.conn.NewSession(nil)
	var recipes []*db.Recipe
	sess.Select("*").
		From("recipes").
		Where("label LIKE ?", fmt.Sprintf("%%%s%%", req.Q)).
		Load(&recipes)
	if len(recipes) == 0 {
		return errors.New("no results found")
	}
	res.Entries = make([]*proton.Recipe, len(recipes))
	for i, recipe := range recipes {
		res.Entries[i] = &proton.Recipe{
			Label: recipe.Label,
			Yield: recipe.Yield,
			Url:   recipe.URL,
			Uri:   recipe.URI,
		}
	}
	return nil
}

// Add adds given set of results into the database
func (d *DataStore) Add(ctx context.Context, res *proton.Result, e *empty.Empty) error {
	fmt.Println(res)
	// ensure that writes happen sequentially to avoid db deadlock
	d.mu.Lock()
	defer d.mu.Unlock()
	sess := d.conn.NewSession(nil)
	tx, err := sess.Begin()
	if err != nil {
		return errors.Wrap(err, "could not create insertion transaction")
	}
	defer tx.RollbackUnlessCommitted()
	if len(res.Entries) > 0 {
		for _, entry := range res.Entries {
			_, err := tx.InsertBySql("INSERT IGNORE INTO recipes (label, yield, calories, url, uri) VALUES (?,?,?,?,?)",
				entry.Label, entry.Yield, entry.Calories, entry.Url, entry.Uri).
				ExecContext(ctx)
			if err != nil {
				return errors.Wrapf(err, "failed to insert recipe %s", entry.Label)
			}
		}
		tx.Commit()
	}
	return nil
}
