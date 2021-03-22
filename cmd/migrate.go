package cmd

import (
	"database/sql"
	"fmt"
	"os"

	"example.com/app/helpers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/golang-migrate/migrate/database/postgres"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/spf13/cobra"
)

var CmdMigrateMysql = &cobra.Command{
	Use:   "migrate-mysql",
	Short: "migrate database mysql",
	Run: func(cmd *cobra.Command, args []string) {
		upMysql()
	},
}

var CmdRollbackMysql = &cobra.Command{
	Use:   "rollback-mysql",
	Short: "rollback database mysql",
	Run: func(cmd *cobra.Command, args []string) {
		rollbackMysql()
	},
}

func migrateMysql() (*migrate.Migrate, error) {
	db, err := sql.Open("mysql", os.Getenv("MYSQL_DB_USER")+":"+os.Getenv("MYSQL_DB_PASS")+"@("+os.Getenv("MYSQL_DB_HOST")+":"+os.Getenv("MYSQL_DB_PORT")+")/"+os.Getenv("MYSQL_DB_NAME")+"")
	if err != nil {
		helpers.Log("failed to load the database")
		return nil, err
	}
	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		helpers.Log("ping to the database host failed")
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://migration/mysql", os.Getenv("MYSQL_DB_NAME"), driver)
	if err != nil {
		fmt.Println("Err : ", err)
		helpers.Log("failed to prepare migration")
		return nil, err
	}
	return m, nil
}

func upMysql() {
	m, _ := migrateMysql()
	fmt.Println("mysql migration successful")
	m.Up()
}
func rollbackMysql() {
	m, _ := migrateMysql()
	fmt.Println("mysql rollback successful")
	m.Down()
}

var CmdMigratePostgres = &cobra.Command{
	Use:   "migrate-pg",
	Short: "migrate database postgresql",
	Run: func(cmd *cobra.Command, args []string) {
		upPostgres()
	},
}

var CmdRollbackPostgres = &cobra.Command{
	Use:   "rollback-pg",
	Short: "rollback database postgresql",
	Run: func(cmd *cobra.Command, args []string) {
		rollbackPostgres()
	},
}

func migratePostgres() (*migrate.Migrate, error) {
	db, err := sql.Open("postgres", "postgres://"+os.Getenv("PG_DB_USER")+":"+os.Getenv("PG_DB_PASS")+"@"+os.Getenv("PG_DB_HOST")+":"+os.Getenv("PG_DB_PORT")+"/"+os.Getenv("PG_DB_NAME")+"?sslmode=disable")
	if err != nil {
		helpers.Log("failed to load the database")
		return nil, err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		helpers.Log("ping to the database host failed")
		return nil, err
	}
	m, err := migrate.NewWithDatabaseInstance("file://migration/mysql", os.Getenv("PG_DB_NAME"), driver)
	if err != nil {
		fmt.Println("Err : ", err)
		helpers.Log("failed to prepare migration")
		return nil, err
	}
	return m, nil
}

func upPostgres() {
	m, _ := migratePostgres()
	fmt.Println("postgres migration successful")
	m.Up()
}
func rollbackPostgres() {
	m, _ := migratePostgres()
	fmt.Println("postgres rollback successful")
	m.Down()
}
