package database

import (
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
func (db *SupabaseDb) Update() {
}

func (db *SupabaseDb) Delete() {
}
