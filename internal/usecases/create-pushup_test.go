package usecases

import (
	"bytes"
	"database/sql"
	"io"
	"os"
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

func TestCreatePushup_Success_Add(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	newPushup := entities.NewPushup{
		Repetitions: 10,
		Type:        entities.OpAdd,
		CreatedAt:   time.Now(),
	}

	// Capture stdout
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := CreatePushup(db, newPushup)

	w.Close()
	os.Stdout = old
	output, _ := io.ReadAll(r)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	outputStr := string(output)
	if outputStr == "" {
		t.Error("expected output, got empty string")
	}

	if !bytes.Contains(output, []byte("flexões foram registradas")) {
		t.Errorf("expected success message for OpAdd, got: %s", outputStr)
	}
}

func TestCreatePushup_Success_Subtract(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	newPushup := entities.NewPushup{
		Repetitions: 5,
		Type:        entities.OpSubtract,
		CreatedAt:   time.Now(),
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := CreatePushup(db, newPushup)

	w.Close()
	os.Stdout = old
	output, _ := io.ReadAll(r)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !bytes.Contains(output, []byte("removidas com sucesso")) {
		t.Errorf("expected success message for OpSubtract, got: %s", output)
	}
}

func TestCreatePushup_Invalid_Operation_Type(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	newPushup := entities.NewPushup{
		Repetitions: 10,
		Type:        entities.PushupOperationType("invalid"),
		CreatedAt:   time.Now(),
	}

	old := os.Stdout
	_, w, _ := os.Pipe()
	os.Stdout = w
	defer func() {
		w.Close()
		os.Stdout = old
	}()

	err := CreatePushup(db, newPushup)

	// The function still succeeds but doesn't print anything for unknown types
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
