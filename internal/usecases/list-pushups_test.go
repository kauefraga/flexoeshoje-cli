package usecases

import (
	"bytes"
	"io"
	"os"
	"testing"
	"time"

	"github.com/kauefraga/flexoeshoje-cli/v2/internal/entities"
	"github.com/kauefraga/flexoeshoje-cli/v2/internal/infra"
	_ "modernc.org/sqlite"
)

func TestListPushups_Empty(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := ListPushups(db)

	w.Close()
	os.Stdout = old
	output, _ := io.ReadAll(r)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	outputStr := string(output)
	if !bytes.Contains(output, []byte("Nenhuma flexão")) {
		t.Errorf("expected empty message, got: %s", outputStr)
	}
}

func TestListPushups_Single_Add(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Insert a test pushup
	newPushup := entities.NewPushup{
		Repetitions: 15,
		Type:        entities.OpAdd,
		CreatedAt:   time.Now(),
	}
	if err := infra.CreateOnePushup(db, newPushup); err != nil {
		t.Fatalf("failed to create test data: %v", err)
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := ListPushups(db)

	w.Close()
	os.Stdout = old
	output, _ := io.ReadAll(r)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	outputStr := string(output)
	// Should show total of 15
	if !bytes.Contains(output, []byte("15")) {
		t.Errorf("expected total of 15, got: %s", outputStr)
	}

	if !bytes.Contains(output, []byte("Hoje você fez")) {
		t.Errorf("expected summary message, got: %s", outputStr)
	}
}

func TestListPushups_Multiple_Operations(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Insert multiple test pushups
	pushups := []entities.NewPushup{
		{Repetitions: 10, Type: entities.OpAdd, CreatedAt: time.Now().Add(-1 * time.Hour)},
		{Repetitions: 5, Type: entities.OpAdd, CreatedAt: time.Now().Add(-30 * time.Minute)},
		{Repetitions: 3, Type: entities.OpSubtract, CreatedAt: time.Now()},
	}

	for _, p := range pushups {
		if err := infra.CreateOnePushup(db, p); err != nil {
			t.Fatalf("failed to create test data: %v", err)
		}
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := ListPushups(db)

	w.Close()
	os.Stdout = old
	output, _ := io.ReadAll(r)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	outputStr := string(output)
	// Total should be 10 + 5 - 3 = 12
	if !bytes.Contains(output, []byte("12")) {
		t.Errorf("expected total of 12 (10+5-3), got: %s", outputStr)
	}
}

func TestListPushups_All_Subtract(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	// Insert subtract-only pushups
	pushups := []entities.NewPushup{
		{Repetitions: 5, Type: entities.OpSubtract, CreatedAt: time.Now().Add(-1 * time.Hour)},
		{Repetitions: 3, Type: entities.OpSubtract, CreatedAt: time.Now()},
	}

	for _, p := range pushups {
		if err := infra.CreateOnePushup(db, p); err != nil {
			t.Fatalf("failed to create test data: %v", err)
		}
	}

	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := ListPushups(db)

	w.Close()
	os.Stdout = old
	output, _ := io.ReadAll(r)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	outputStr := string(output)
	// Total should be -8 (all subtractions)
	if !bytes.Contains(output, []byte("-8")) {
		t.Errorf("expected total of -8, got: %s", outputStr)
	}
}

func TestListPushups_Timestamp_Format(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()

	now := time.Now()
	newPushup := entities.NewPushup{
		Repetitions: 10,
		Type:        entities.OpAdd,
		CreatedAt:   now,
	}
	if err := infra.CreateOnePushup(db, newPushup); err != nil {
		t.Fatalf("failed to create test data: %v", err)
	}

	// Query back to verify timestamp parsing works
	pushups, err := infra.FindTodayPushups(db)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if len(pushups) != 1 {
		t.Errorf("expected 1 pushup, got %d", len(pushups))
	}

	// Verify timestamp was parsed correctly (same day, within reasonable time difference)
	if pushups[0].CreatedAt.Format("2006-01-02") != now.Format("2006-01-02") {
		t.Errorf("expected date %s, got %s",
			now.Format("2006-01-02"),
			pushups[0].CreatedAt.Format("2006-01-02"))
	}
}
