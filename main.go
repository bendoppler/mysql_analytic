package main

import (
	"fmt"

	"github.com/piendop/mysql_analytic/database"
)

type Driver struct {
	Surname                string
	Name                   string
	Position               int
	PositionPerConstructor string `gorm:"column:position_per_constructor"`
}

func main() {
	//connect to postgres
	db := database.GetConnectionDb()
	rows, err := db.Raw("SELECT drivers.surname, constructors.name, results.position, concat(ROW_NUMBER() OVER (PARTITION by constructors.constructorId ORDER BY position),'/', COUNT(*) OVER (PARTITION BY constructors.constructorId)) AS position_per_constructor FROM results JOIN drivers using(driverId) JOIN constructors using(constructorId) WHERE raceId = 22 ORDER BY position DESC;").Rows()
	if err != nil {
		panic(err)
	}

	fmt.Println(rows)
	for rows.Next() {
		var driver Driver
		db.ScanRows(rows, &driver)
		fmt.Println(driver)
	}
	defer db.Close()
}
