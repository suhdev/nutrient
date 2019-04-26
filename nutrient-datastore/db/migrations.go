package db

import migrate "github.com/rubenv/sql-migrate"

// Migrations are a continuous list of sql migrations
var Migrations = &migrate.MemoryMigrationSource{
	Migrations: []*migrate.Migration{
		{
			Id: "01-recipes",
			Up: []string{
				`CREATE TABLE IF NOT EXISTS recipes (
				id BIGINT AUTO_INCREMENT PRIMARY KEY,
				uri TEXT,
				url TEXT,
				label TEXT,
				yield DOUBLE,
				calories TEXT,
				total_time DOUBLE,
				total_weight DOUBLE
				)`,
				`CREATE INDEX idx_recipe_label ON recipes (label)`,
				`CREATE UNIQUE INDEX idx_uq_uri ON recipes (uri)`},
			Down: []string{"DROP TABLE IF EXISTS recipe"},
		},
	},
}
