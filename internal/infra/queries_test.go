package infra

import (
	"database/sql"
	"testing"
	"time"

	"github.com/kauefraga/flexoeshoje-cli/v2/internal/entities"
	_ "modernc.org/sqlite"
)

func setupTestDB(t *testing.T) *sql.DB {
	t.Helper()

	db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		t.Fatalf("Ocorreu um erro ao conectar com o banco de dados: %v", err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS pushups (
		id INTEGER PRIMARY KEY,
		repetitions INTEGER NOT NULL,
		type TEXT NOT NULL CHECK(type in ('add', 'subtract')) DEFAULT 'add',
		created_at TEXT NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`)
	if err != nil {
		t.Fatalf("Ocorreu um erro ao criar a tabela de flexões: %v", err)
	}

	return db
}

func TestCreateOnePushup(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	pushup := entities.NewPushup{
		Repetitions: 10,
		Type:        entities.OpAdd,
		CreatedAt:   time.Now(),
	}

	err := CreateOnePushup(db, pushup)
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}
}

func TestFindTodayPushups(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	now := time.Now()

	t.Run("returns empty when no records", func(t *testing.T) {
		pushups, err := FindTodayPushups(db)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if len(pushups) != 0 {
			t.Errorf("expected 0 pushups, got %d", len(pushups))
		}
	})

	t.Run("returns today's records", func(t *testing.T) {
		_ = CreateOnePushup(db, entities.NewPushup{Repetitions: 10, Type: entities.OpAdd, CreatedAt: now})
		_ = CreateOnePushup(db, entities.NewPushup{Repetitions: 5, Type: entities.OpSubtract, CreatedAt: now})

		pushups, err := FindTodayPushups(db)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		if len(pushups) != 2 {
			t.Errorf("expected 2 pushups, got %d", len(pushups))
		}
	})

	t.Run("does not return old records", func(t *testing.T) {
		yesterday := now.AddDate(0, 0, -1)
		_ = CreateOnePushup(db, entities.NewPushup{Repetitions: 20, Type: entities.OpAdd, CreatedAt: yesterday})

		pushups, err := FindTodayPushups(db)
		if err != nil {
			t.Fatalf("expected no error, got: %v", err)
		}
		// still 2 from previous subtest, yesterday's record excluded
		if len(pushups) != 2 {
			t.Errorf("expected 2 pushups, got %d", len(pushups))
		}
	})
}
