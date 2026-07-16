package domain

type DB interface {
	Insert(table string, data map[string]any, upsert bool) error
	Select(table string, columns []string, filters map[string]string) ([]byte, error)
	Update()
	Delete()
	SelectMany(table string, columns []string, filters map[string]int) ([]byte, error)
}

type DBFactory func(env *Env) DB
