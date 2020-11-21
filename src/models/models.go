package models

import (
	"database/sql"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

type League struct  {
	ID           int     `json:"id"`
	Name         string  `json:"name"`
	CurrentWeek  int  `json:"current_week"`
	TotalWeek    int  `json:"total_week"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	Actions   string  `json:"actions"`
}
type LeagueCollection struct {
    Leagues []League `json:"items"`
}

type Team struct {
	ID           int     `json:"id"`
	LeagueID     int     `json:"league_id"`
	Name         string  `json:"name"`
	LeagueName   string  `json:"league_name"`
	Update       string  `json:"updated"`
	Actions      string  `json:"actions"`
}

type Score struct {
	ID           int     `json:"id"`
	TeamID       int     `json:"team_id"`
	Won          int     `json:"won"`
	Drawn        int     `json:"drawn"`
	Lost         int     `json:"lost"`
	For          int     `json:"for"`
	Against      int     `json:"against"`
	GoalDiff     int     `json:"goal_diff"`
	Points       int     `json:"points"`
	Rate         int     `json:"rate"`
}

type Match struct {
	ID              int     `json:"id"`
	Week            int     `json:"week"`
	HomeTeamID      int     `json:"home_team_id"`
	AwayTeamID      int     `json:"away_team_id"`
	HomeTeamScore   int     `json:"home_team_score"`
	AwayTeamScore   int     `json:"away_team_score"`
	IsPlayed        int     `json:"is_played"` // 0 if not played, else 1
}

type MatchCollection struct {
    Matches []Match `json:"items"`
}

type ScoreCollection struct {
    Scores []Score `json:"items"`
}

type TeamCollection struct {
    Teams []Team `json:"items"`
}

func GetLeagues(db *sql.DB) LeagueCollection {
    sql := "SELECT * FROM leagues"
    rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
    }
    // make sure to cleanup when the program exits
    defer rows.Close()

    result := LeagueCollection{}
    for rows.Next() {
        league := League{}
        err2 := rows.Scan(&league.ID, &league.Name, &league.CurrentWeek, &league.TotalWeek, &league.CreatedAt, &league.UpdatedAt, &league.Actions)
        // Exit if we get an error
        if err2 != nil {
            panic(err2)
        }
        result.Leagues = append(result.Leagues, league)
    }
    return result
}

func GetTeams(db *sql.DB, league_id int) TeamCollection {
	sql := "SELECT * FROM teams WHERE league_id = " + strconv.Itoa(league_id)
    rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
    }
    // make sure to cleanup when the program exits
	defer rows.Close()

    result := TeamCollection{}
    for rows.Next() {
        team := Team{}
        err2 := rows.Scan(&team.ID, &team.LeagueID, &team.Name, &team.LeagueName, &team.Update, &team.Actions)
        // Exit if we get an error
        if err2 != nil {
            panic(err2)
        }
        result.Teams = append(result.Teams, team)
    }
    return result
}


func GetTeam(db *sql.DB, league_id int, team_id int) Team {
	sql := "SELECT * FROM teams WHERE league_id = " + strconv.Itoa(league_id) + " AND id = " +  strconv.Itoa(team_id) 
    rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
	}
	team := Team{}
    // make sure to cleanup when the program exits
	defer rows.Close()
    for rows.Next() {
        err2 := rows.Scan(&team.ID, &team.LeagueID, &team.Name, &team.LeagueName, &team.Update, &team.Actions)
        // Exit if we get an error
        if err2 != nil {
            panic(err2)
        }
    }
    return team
}

func PutLeague(db *sql.DB, league League) (int64, error) {
    sql := "INSERT INTO leagues(name, current_week, total_week, created_at, updated_at, actions) VALUES(?, ?, ?, ?, ?, ?)"

    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    if err != nil {
        panic(err)
    }
    // Make sure to cleanup after the program exits
    defer stmt.Close()

    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(league.Name, league.CurrentWeek, league.TotalWeek, league.CreatedAt, league.UpdatedAt, league.Actions)
    // Exit if we get an error
    if err2 != nil {
        panic(err2)
    }

    return result.LastInsertId()
}


func PutTeam(db *sql.DB, league_id int, team Team) (int64, error) {
    sql := "INSERT INTO teams(league_id, name, league_name, updated, actions) VALUES(?, ?, ?, ?, ?)"

    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    if err != nil {
        panic(err)
    }
    // Make sure to cleanup after the program exits
    defer stmt.Close()

    // Replace the '?' in our prepared statement with 'name'
    result, err2 := stmt.Exec(team.LeagueID, team.Name, team.LeagueName, team.Update, team.Actions)
    // Exit if we get an error
    if err2 != nil {
        panic(err2)
    }

    return result.LastInsertId()
}

func DeleteLeague(db *sql.DB, id int) (int64, error) {
    sql := "DELETE FROM leagues WHERE id = ?"

    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    if err != nil {
        panic(err)
    }

    // Replace the '?' in our prepared statement with 'id'
    result, err2 := stmt.Exec(id)
    // Exit if we get an error
    if err2 != nil {
        panic(err2)
    }

    return result.RowsAffected()
}

func DeleteTeam(db *sql.DB, league_id int, team_id int) (int64, error) {
    sql := "DELETE FROM leagues WHERE league_id = ? AND team_id = ?"

    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    if err != nil {
        panic(err)
    }

    // Replace the '?' in our prepared statement with 'id'
    result, err2 := stmt.Exec(league_id, team_id)
    // Exit if we get an error
    if err2 != nil {
        panic(err2)
    }

    return result.RowsAffected()
}

/*         MATCH          */
func CreateMatch(team1 Team, team2 Team, week int) (int64, error) {
	sql := "INSERT INTO matches(week, home_team_id, away_team_id, home_team_score, away_team_score, is_played) VALUES(?, ?, ?, ?, ?, ?)"

    // Create a prepared SQL statement
    stmt, err := db.Prepare(sql)
    // Exit if we get an error
    if err != nil {
        panic(err)
    }
    // Make sure to cleanup after the program exits
    defer stmt.Close()

    // Replace the '?' in our prepared statement with 'values'
    result, err2 := stmt.Exec(week, team1.ID, team2.ID, 0, 0, 0)
    // Exit if we get an error
    if err2 != nil {
        panic(err2)
    }

    return result.LastInsertId()
}

func DistributeFixture(db *sql.DB, league_id int){ // I configured the league fixed as 6 weeks since there are 4 teams in leagues. However it can be changed due to the expected outcome
	//rand.Seed(time.Now().UnixNano()) teams or weeks can be determined randomly,
	//for the sake of simplicity it's kind of hardcoded right now
	min := 0
	max := 3
	Teams := GetTeams(db, league_id)
	id1, _ := CreateMatch(1, Teams.Teams[3], Teams.Teams[0])
	id2, _ := CreateMatch(1, Teams.Teams[1], Teams.Teams[0])
	id3, _ := CreateMatch(2, Teams.Teams[2], Teams.Teams[3])
	id4, _ := CreateMatch(2, Teams.Teams[1], Teams.Teams[2])
	id5, _ := CreateMatch(3, Teams.Teams[0], Teams.Teams[3])
	id6, _ := CreateMatch(3, Teams.Teams[2], Teams.Teams[1])
	id7, _ := CreateMatch(4, Teams.Teams[3], Teams.Teams[0])
	id8, _ := CreateMatch(4, Teams.Teams[0], Teams.Teams[1])
	id9, _ := CreateMatch(5, Teams.Teams[1], Teams.Teams[3])
	id10, _ := CreateMatch(5, Teams.Teams[0], Teams.Teams[2])
	id11, _ := CreateMatch(6, Teams.Teams[2], Teams.Teams[0])
	id12, _ := CreateMatch(6, Teams.Teams[3], Teams.Teams[2])
	
}

func PlayOneWeek(db *sql.DB, week int){
	sql := "SELECT * FROM matches WHERE week = " + + strconv.Itoa(week)
	rows, err := db.Query(sql)
    // Exit if the SQL doesn't work for some reason
    if err != nil {
        panic(err)
    }
    // make sure to cleanup when the program exits
	defer rows.Close()

    result := MatchCollection{}
    for rows.Next() {
        match := Match{}
		err2 := rows.Scan(&match.ID, &match.Week, &match.HomeTeamID, &match.AwayTeamID, &match.HomeTeamScore, &match.AwayTeamScore, &match.IsPlayed)
		rand.Seed(time.Now().UnixNano()) // use random seed to determine goals
		goal1 = rand.Intn(4) 
		goal2 = rand.Intn(5) 
		// update matches
		sql = "UPDATE matches SET is_played = 1, home_team_score = ?, away_team_score = ? WHERE home_team_id = "+match.HomeTeamID +" AND away_team_id = "+match.AwayTeamID
		stmt, err3 := db.Prepare(sql)
		if err != nil {
			panic(err)
		}
		defer stmt.Close()
		_, err4 := stmt.Exec(goal1, goal2)
		// new scores
		switch {
		case goal1 > goal2:
			sql1 := "UPDATE scores SET won = won + 1, points = points + 3, for = for + "+goal1+", against = against + "+goal2+" WHERE team_id = "+home_team_id
			sql2 := "UPDATE scores SET lost = lost + 1, for = for + "+goal2+", against = against + "+goal1+"WHERE team_id = "+away_team_id
		case goal2 > goal1:
			sql1 := "UPDATE scores SET won = won + 1, points = points + 3, for = for + "+goal2+", against = against + "+goal1+"WHERE team_id = "+away_team_id
			sql2 := "UPDATE scores SET lost = lost + 1, for = for + "+goal1+", against = against + "+goal2+"WHERE team_id = "+home_team_id
		case goal2 == goal1:
			sql1 := "UPDATE scores SET drawn = drawn + 1, points = points + 1, for = for + "+goal2+", against = against + "+goal1+"WHERE team_id = "+away_team_id
			sql2 := "UPDATE scores SET drawn = drawn + 1, points = points + 1, for = for + "+goal1+", against = against + "+goal2+"WHERE team_id = "+home_team_id
		}
		sql1.Query()
		sql2.Query()

		// Exit if we get an error
        if err2 != nil {
            panic(err2)
		}
}
        result.Teams = append(result.Teams, team)
    }
}
	PlayAll(db))
	PredictLeaders(db))
