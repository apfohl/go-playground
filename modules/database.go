package modules

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
)

const (
	DSN   string = "postgres://postgres:postgres@localhost:5432/postgres"
	TABLE string = "items"
)

type (
	Data struct {
		FirstName string `json:"first_name,omitempty"`
		LastName  string `json:"last_name,omitempty"`
		IsDefault bool   `json:"-"`
	}

	Item struct {
		Id   uuid.UUID `db:"id" goqu:"skipinsert, skipupdate"`
		Name string    `db:"name"`
		Data *Data     `db:"data"`
	}
)

func (u Data) String() string {
	return fmt.Sprintf("%s %s", u.FirstName, u.LastName)
}

func (i Item) String() string {
	dataStr := "nil"
	if !i.Data.IsDefault {
		dataStr = i.Data.String()
	}
	return fmt.Sprintf("Id: %s, Item: %s, Data: %s", i.Id, i.Name, dataStr)
}

func (u *Data) Scan(src any) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("type assertion to byte failed")
	}

	// Check if the JSON is an empty object
	if string(b) == "{}" {
		u = nil
		//(*u).IsDefault = true
		return nil // Keep u as nil
	}

	return json.Unmarshal(b, u)
}

func (u *Data) Value() (driver.Value, error) {
	if u == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(u)
}

func NewDatabaseCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "database",
		Short: "Run database operations",
		Run:   runDatabase,
	}
}

func runDatabase(_ *cobra.Command, _ []string) {
	database, err := sql.Open("pgx", DSN)
	if err != nil {
		log.Fatalf("Unable to open connection: %v\n", err)
		return
	}
	defer func(database *sql.DB) {
		if err = database.Close(); err != nil {
			log.Fatalf("Unable to close connection: %v\n", err)
			return
		}
	}(database)

	readDb := goqu.New("postgres", database)

	newItem := Item{
		Name: "Phone",
		Data: &Data{
			FirstName: "John",
			LastName:  "Doe",
		},
	}

	var insertedID uuid.UUID
	insert := readDb.Insert(TABLE).Rows(newItem).Returning("id")
	if _, err = insert.Executor().ScanVal(&insertedID); err != nil {
		log.Fatalf("Unable to insert record: %v\n", err)
		return
	}

	fmt.Printf("Inserted record with ID: %s\n", insertedID)

	var items []Item
	err = readDb.From(TABLE).Prepared(false).Where(goqu.Ex{"id": insertedID}).ScanStructs(&items)
	if err != nil {
		log.Fatalf("Unable to scan structs: %v\n", err)
		return
	}

	for _, item := range items {
		fmt.Println(item)
	}
}
