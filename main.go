package main

import (
	"encoding/csv"
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sudyusuk/csv_to_db/db_connect"
	"os"
	"strconv"
)

type Databases struct {
	UserDatabase *gorm.DB
}

type TableName struct {
	Name     string    `gorm:"column:name"`
	Age    int `gorm:"column:age"`
	Address       string    `gorm:"column:address"`
}

func main() {

	user_db := database_connect.GormConnect
	databases := Databases{user_db()}

	if err := databases.readJson("./sample.csv");err != nil {
		fmt.Println(err)
	}

}

func (d *Databases) readJson(path string) error {

	csvFile, err := os.Open(path)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	reader.FieldsPerRecord = -1

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, data := range csvData {
		age ,_ := strconv.Atoi(data[1])

		oneLine := CreateOneList(data[0],age,data[2])
		d.CreateDatabase(*oneLine)
	}
	return nil
}

func CreateOneList(name string,age int ,address string) *TableName {
	return &TableName{name,age,address}
}

func (d *Databases) CreateDatabase(oneLine TableName) {
	d.UserDatabase.Create(&oneLine)
}
