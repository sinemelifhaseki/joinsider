package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/labstack/echo"
	"main.go/handler"
)

func initDB(filepath string) *sql.DB {
    db, err := sql.Open("sqlite3", filepath)

    // Here we check for any db errors then exit
    if err != nil {
        panic(err)
    }

    // If we don't get any errors but somehow still don't get a db connection
    // we exit as well
    if db == nil {
        panic("db nil")
    }
    return db
}

func migrate(db *sql.DB) {
	sql := `
    CREATE TABLE IF NOT EXISTS leagues(
        id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name VARCHAR NOT NULL,
		current_week INTEGER NOT NULL,
		total_week INTEGER NOT NULL,
		created_at VARCHAR NOT NULL,
		updated_at VARCHAR NOT NULL,
		actions VARCHAR
	);
	
	CREATE TABLE IF NOT EXISTS teams(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		league_id INTEGER NOT NULL,
		name VARCHAR NOT NULL,
		league_name VARCHAR NOT NULL,
		updated VARCHAR NOT NULL,
		actions VARCHAR
    );
    `
    _, err := db.Exec(sql)
    // Exit if something goes wrong with our SQL statement above
    if err != nil {
        panic(err)
    }
}

func main() {
	db := initDB("storage.db")
    migrate(db)
	// Create a new instance of Echo
    e := echo.New()

	// PUT falan da olcak 
    e.GET("/", handler.GetLeagues(db))
    e.GET("/leagues", handler.GetLeagues(db))
	e.GET("/leagues/:league_id/teams", handler.GetTeams(db))
	e.GET("/leagues/:league_id/teams/:team_id", handler.GetTeam(db))
	e.PUT("/addleague", handler.PutLeague(db))
	e.PUT("/leagues/:league_id/add_team", handler.PutTeam(db))
    e.DELETE("/leagues/:league_id/delete", handler.DeleteLeague(db))
    e.DELETE("/leagues/:league_id/delete_team/:team_id", handler.DeleteTeam(db))
	

    // Start as a web server
    e.Start(":8000")
}