package dbteam

import (
	"fmt"
	"github.com/gkeele21/topdawgsportsAPI/internal/app/database"
)

type Team struct {
	TeamID           int64               `db:"team_id"`
	Name             string              `db:"name"`
	Abbreviation     database.NullString `db:"abbreviation"`
	Mascot           database.NullString `db:"mascot"`
	Color            database.NullString `db:"color"`
	AltColor         database.NullString `db:"alt_color"`
	Logo             database.NullString `db:"logo"`
	LogoDark         database.NullString `db:"logo_dark"`
	Level            string              `db:"level"`
	FootballDivision database.NullString `db:"football_division"`
	ExternalName     database.NullString `db:"external_name"`
}

// ReadByID reads user by id column
func ReadByID(ID int64) (*Team, error) {
	t := Team{}
	err := database.Get(&t, "SELECT * FROM team where team_id = ?", ID)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ReadAll reads all teams in the database
func ReadAll() ([]Team, error) {
	var teams []Team
	err := database.Select(&teams, "SELECT * FROM team")
	if err != nil {
		return nil, err
	}

	return teams, nil
}

// Delete deletes a team from the database
func Delete(t *Team) error {
	_, err := database.Exec("DELETE FROM team WHERE team_id = ?", t.TeamID)
	if err != nil {
		return fmt.Errorf("team: couldn't delete team %s", err)
	}

	return nil
}

// Insert will create a new record in the database
func Insert(t *Team) error {
	res, err := database.Exec(database.BuildInsert("team", t), database.GetArguments(*t)...)

	if err != nil {
		return fmt.Errorf("team: couldn't insert new %s %#v", err, t)
	}

	ID, err := res.LastInsertId()
	if err != nil {
		return fmt.Errorf("team: couldn't get last inserted ID %S", err)
	}

	t.TeamID = ID

	return nil
}

// Update will update a record in the database
func Update(s *Team) error {
	sql := database.BuildUpdate("team", s)
	_, err := database.Exec(sql, database.GetArgumentsForUpdate(*s)...)

	if err != nil {
		return fmt.Errorf("team: couldn't update %s", err)
	}

	return nil
}

func Save(s *Team) error {
	if s.TeamID > 0 {
		return Update(s)
	} else {
		return Insert(s)
	}
}

// ReadByAbbreviationAndMascot reads user by abbreviation and mascot columns
func ReadByAbbreviationAndMascot(abbrev, mascot string) (*Team, error) {
	t := Team{}
	err := database.Get(&t, "SELECT * FROM team WHERE abbreviation = ? AND mascot = ?", abbrev, mascot)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ReadByAbbreviationAndLevel reads user by abbreviation and level columns
func ReadByAbbreviationAndLevel(abbr, level string) (*Team, error) {
	t := Team{}
	err := database.Get(&t, "SELECT * FROM team WHERE abbreviation = ? AND level = ?", abbr, level)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ReadByNameAndLevel reads user by name and level columns
func ReadByNameAndLevel(name, level string) (*Team, error) {
	t := Team{}
	err := database.Get(&t, "SELECT * FROM team WHERE name = ? AND level = ?", name, level)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

// ReadByExternalNameAndLevel reads user by external_name and level columns
func ReadByExternalNameAndLevel(name, level string) (*Team, error) {
	t := Team{}
	err := database.Get(&t, "SELECT * FROM team WHERE external_name = ? AND level = ?", name, level)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
