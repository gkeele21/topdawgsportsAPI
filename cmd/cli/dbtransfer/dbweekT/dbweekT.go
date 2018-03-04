package main

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"topdawgsportsAPI/pkg/database"
	"topdawgsportsAPI/pkg/database/dbweek"
	"strings"
)

var db *sql.DB

func main() {
	// grab all players from the existing database
	db, err := sql.Open("mysql", "webuser:lakers55@tcp(127.0.0.1:3306)/topdawg?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT sw.SeasonWeekID, sw.SeasonID, sw.WeekNo, sw.StartDate, sw.EndDate, sw.Status, sw.WeekType, fsw.FSSeasonWeekID FROM SeasonWeek sw INNER JOIN FSSeasonWeek fsw ON fsw.SeasonWeekID = sw.SeasonWeekID")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var seasonweekid, seasonid, fsseasonweekid int64
		var weekno, status, weektype string
		var startdate, enddate database.NullTime

		if err := rows.Scan(&seasonweekid, &seasonid, &weekno, &startdate, &enddate, &status, &weektype, &fsseasonweekid); err != nil {
			log.Fatal(err)
		}

		fmt.Printf("WeekID : [%d], SeasonID : [%d], WeekName : [%s], Start : [%s], End : [%s], Status : [%s], WeekType : [%s]\n", seasonweekid, seasonid, weekno, startdate, enddate, status, weektype)

		if status == "COMPLETED" {
			status = "final"
		} else if status == "CURRENT" {
			status = "active"
		}

		week := dbweek.Week{
			WeekID:    fsseasonweekid,
			SeasonID:  seasonid,
			WeekName:  weekno,
			StartDate: startdate,
			EndDate:   enddate,
			Status:    strings.ToLower(status),
			WeekType:  strings.ToLower(weektype),
		}

		fmt.Printf("Week : %v\n", week)
		err := dbweek.Insert(&week)
		if err != nil {
			log.Fatal(err)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}
