package db

type Recipe struct {
	// ID recipe identifier
	ID int64
	// URI recipe URI
	URI string `db:"uri"`
	// URL recipe URL
	URL string `db:"url"`
	// Label recipe label
	Label string `db:"label"`
	// Yield recipe yield
	Yield float64 `db:"yield"`
	// Calories recipe calories
	Calories string `db:"calories"`
	// TotalTime recipe total time to prepare
	TotalTime float64 `db:"total_time"`
	// TotalWeight recipe total weight
	TotalWeight float64 `db:"total_weight"`
}
