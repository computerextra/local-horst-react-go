package mitarbeiter

import (
	"database/sql"
	"time"

	database "github.com/computerextra/local-horst-react-go/Database"
)

type Mitarbeiter struct {
	Id                 string
	Name               string
	Short              sql.NullString
	Gruppenwahl        sql.NullString
	InternTelefon1     sql.NullString
	InternTelefon2     sql.NullString
	FestnetzAlternativ sql.NullString
	FestnetzPrivat     sql.NullString
	HomeOffice         sql.NullString
	MobilBusiness      sql.NullString
	MobilPrivat        sql.NullString
	Email              sql.NullString
	Azubi              sql.NullBool
	Geburtstag         sql.NullString
}

func GetMitarbeiter() ([]Mitarbeiter, error) {
	var AlleMitarbeiter []Mitarbeiter

	connString := database.GetMySqlConnectionString()

	db, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	db.SetConnMaxIdleTime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	rows, err := db.Query("SELECT * FROM Mitarbeiter ORDER BY Name ASC")
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var ma Mitarbeiter
		if err := rows.Scan(&ma.Id, &ma.Name, &ma.Short, &ma.Gruppenwahl, &ma.InternTelefon1, &ma.InternTelefon2, &ma.FestnetzAlternativ, &ma.FestnetzPrivat, &ma.HomeOffice, &ma.MobilBusiness, &ma.MobilPrivat, &ma.Email, &ma.Azubi, &ma.Geburtstag); err != nil {
			return nil, err
		}

		AlleMitarbeiter = append(AlleMitarbeiter, ma)
	}

	return AlleMitarbeiter, nil
}
