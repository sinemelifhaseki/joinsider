package handler

import (
	"database/sql"
    "net/http"
    "strconv"
    "main.go/models"
    "fmt"
    "github.com/labstack/echo"
)

type H map[string]interface{}

// GetLeagues endpoint
func GetLeagues(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        return c.JSON(http.StatusOK, models.GetLeagues(db))
    }
}

// GetLeague endpoint
func GetTeams(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        league_id, _ := strconv.Atoi(c.Param("league_id"))
        teams := models.GetTeams(db, league_id)
        return c.JSON(http.StatusOK, H{
            "gotTeamsOfLeague": teams.Teams[0].LeagueID,
        })        
    }
}

// GetLeague endpoint
func GetTeam(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
		league_id, _ := strconv.Atoi(c.Param("league_id"))
        team_id, _ := strconv.Atoi(c.Param("team_id"))
        team := models.GetTeam(db, league_id, team_id)
        return c.JSON(http.StatusOK, H{
            "got": team.LeagueID,
            "got2": team.ID,
        })
        
    }
}



// PutTeam endpoint
func PutTeam(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
		league_id, _ := strconv.Atoi(c.Param("league_id"))
        var team models.Team
        c.Bind(&team)
        id, err := models.PutTeam(db, league_id, team)
        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "created": id,
            })
        } else {
            return err
        }
	}
}

// PutLeague endpoint
func PutLeague(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        var league models.League
        c.Bind(&league)
        fmt.Println(league.Name)
        id, err := models.PutLeague(db, league)
        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "created": id,
            })
        } else {
            return err
        }
	}
}

// DeleteLeague endpoint
func DeleteLeague(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        id, _ := strconv.Atoi(c.Param("league_id"))
        _, err := models.DeleteLeague(db, id)
        if err == nil {
            return c.JSON(http.StatusOK, H{
                "deleted": id,
            })
        } else {
            return err
        }
    }
}

// DeleteTeam endpoint
func DeleteTeam(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        league_id, _ := strconv.Atoi(c.Param("league_id"))
		team_id, _ := strconv.Atoi(c.Param("team_id"))
        _, err := models.DeleteTeam(db, league_id, team_id)
        if err == nil {
            return c.JSON(http.StatusOK, H{
                "del": league_id,
                "del2": team_id,
            })
        } else {
            return err
        }
    }
}

/*         MATCH       */
func DistributeFixture(db *sql.DB) echo.HandlerFunc {
    return func(c echo.Context) error {
        league_id, _ := strconv.Atoi(c.Param("league_id"))
        err := models.DistributeFixture(db, league_id)
        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "okStatus": 0,
            })
        } else {
            return err
        }
	}
}

func PlayOneWeek(db *sql.DB, week int) echo.HandlerFunc {
    return func(c echo.Context) error {
        league_id, _ := strconv.Atoi(c.Param("league_id"))
        err := models.PlayOneWeek(db, week, league_id)
        if err == nil {
            return c.JSON(http.StatusCreated, H{
                "okStatus": 0,
            })
        } else {
            return err
        }
	}
}

