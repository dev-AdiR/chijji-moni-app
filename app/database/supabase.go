package database

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/supabase-community/supabase-go"
)

type SupabaseDb struct {
	Client *supabase.Client
}

func (db *SupabaseDb) Insert(table string, data map[string]any, upsert bool) error {

	_, _, err := db.Client.
		From(table).
		Insert(
			data,
			upsert,
			"representation",
			"exact",
			"").
		Execute()

	if err != nil {
		return err
	}

	return nil
}

func (db *SupabaseDb) Select(table string, columns []string, filters map[string]string) ([]byte, error) {

	data, _, err := db.Client.
		From(table).
		Select(strings.Join(columns, ","), "", false).
		Match(filters).
		Single().
		Execute()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *SupabaseDb) SelectMany(table string, columns []string, filters map[string]int) ([]byte, error) {

	userID, ok := filters["user_id"]
	if !ok {
		return nil, errors.New("user_id missing or invalid type in filters")
	}
	cols := strings.Join(columns, ",")
	fmt.Println(cols)
	data, _, err := db.Client.
		From(table).
		Select(cols, "", false).
		Eq("user_id", strconv.Itoa(userID)).
		// Match(filters).
		Execute()

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (db *SupabaseDb) Update() {
}

func (db *SupabaseDb) Delete() {
}
