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
	
	CREATE TABLE IF NOT EXISTS scores(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		team_id INTEGER NOT NULL REFERENCES teams(id),
		won INTEGER NOT NULL,
		drawn INTEGER NOT NULL,
		lost INTEGER NOT NULL,
		for INTEGER NOT NULL,
		against INTEGER NOT NULL,
		goal_diff INTEGER NOT NULL,
		points INTEGER NOT NULL,
		rate INTEGER NOT NULL
    );

	CREATE TABLE IF NOT EXISTS matches(
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		week INTEGER NOT NULL,
		home_team_id INTEGER NOT NULL REFERENCES teams(id),
		away_team_id INTEGER NOT NULL REFERENCES teams(id),
		home_team_score INTEGER NOT NULL,
		away_team_score INTEGER NOT NULL,
		is_played INTEGER NOT NULL
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
    e.DELETE("/leagues/:league_id", handler.DeleteLeague(db))
	e.DELETE("/leagues/:league_id/teams/:team_id", handler.DeleteTeam(db))
	// MATCH 
	e.POST("leagues/${leagueId}/distribute-fixture", handler.DistributeFixture(db))
	e.POST("leagues/${leagueId}/play-one-week", handler.PlayOneWeek(db))
	e.POST("leagues/${leagueId}/play-all", handler.PlayAll(db))
	e.POST("leagues/${leagueId}/predict-leaders", handler.PredictLeaders(db))
	

    // Start as a web server
    e.Start(":8000")
}