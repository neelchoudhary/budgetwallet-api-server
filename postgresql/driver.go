package postgresql

import (
	"database/sql"
	"fmt"

	"github.com/neelchoudhary/budgetmanagergrpc/config"

	"github.com/neelchoudhary/budgetmanagergrpc/utils"
	"github.com/golang-migrate/migrate/v4"
	// Required imports
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// ConnectDB connect to postgresql db
func ConnectDB(dbConfig *config.DBConfig) *sql.DB {
	fmt.Println("Connecting to Postgresql DB... " + dbConfig.Host)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.Dbname)
	db, err := sql.Open("postgres", psqlInfo)
	utils.LogIfFatalAndExit(err)

	err = db.Ping()
	utils.LogIfFatalAndExit(err)

	MigrateUp(dbConfig)

	fmt.Println("You connected to your database.")

	return db
}

// MigrateUp migrate up to most recent migration
// Used for development and continuous integration
// migrate create -ext sql -dir db/migrations -seq create_users_table
// migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path db/migrations up
// migrate -database postgres://postgres:password@localhost:5432/postgres?sslmode=disable -path db/migrations down
func MigrateUp(dbConfig *config.DBConfig) {
	fmt.Println("Migrating Up...")
	sourceURL := "file://postgresql/migrations"
	// postgres://user:password@host:port/dbname?query
	localDbURL := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable",
		dbConfig.User, dbConfig.Password, dbConfig.Host, dbConfig.Dbname)
	m, err := migrate.New(sourceURL, localDbURL)
	utils.LogIfFatalAndExit(err)

	if err := m.Up(); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("Finished Migrations. Up to Date.")
}