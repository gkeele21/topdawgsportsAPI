package dbsportlevel

import (
	"fmt"
	"github.com/gkeele21/topdawgsportsAPI/pkg/database"
)

type SportLevel struct {
	SportLevelID int64  `db:"sport_level_id"`
	SportID      int64  `db:"sport_id"`
	Level        string `db:"level"`
}

type SportLevelFull struct {
	SportLevelID int64  `db:"sport_level_id"`
	SportID      int64  `db:"sport_id"`
	SportLevel   string `db:"level"`
	SportName    string `db:"name"`
}

// ReadByID reads by id column
func ReadByID(ID int64) (*SportLevel, error) {
	d := SportLevel{}
	err := database.Get(&d, "SELECT * FROM sport_level where sport_level_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// ReadByIDFull reads by id column and also includes sport table info
func ReadByIDFull(ID int64) (*SportLevelFull, error) {
	d := SportLevelFull{}
	err := database.Get(&d, "SELECT * FROM sport_level sl inner join sport s on s.sport_id = sl.sport_id where sport_level_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &d, nil
}

// ReadAll reads all records in the database
func ReadAll() ([]SportLevel, error) {
	var recs []SportLevel
	err := database.Select(&recs, "SELECT * FROM sport_level")
	if err != nil {
		return nil, err
	}

	return recs, nil
}

// ReadAllFull reads all records in the database, including the sport name
func ReadAllFull(orderBy string) ([]SportLevelFull, error) {
	var recs []SportLevelFull
	if orderBy == "" {
		orderBy = "sport_level_id asc"
	}

	err := database.Select(&recs, "SELECT * FROM sport_level sl inner join sport s on s.sport_id = sl.sport_id ORDER BY "+orderBy)
	if err != nil {
		return nil, err
	}

	return recs, nil
}

// Delete deletes a record from the database
func Delete(d *SportLevel) error {
	_, err := database.Exec("DELETE FROM sport_level WHERE sport_level_id = ?", d.SportLevelID)
	if err != nil {
		return fmt.Errorf("sport_level: couldn't delete record %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(d *SportLevel) error {
	res, err := database.Exec(database.BuildInsert("sport_level", d), database.GetArguments(*d)...)

	if err != nil {
		return fmt.Errorf("sport_level: couldn't insert new %s", err)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("sport_level: couldn't get last inserted ID %S", err)
	}

	d.SportLevelID = ID

	return nil
}
