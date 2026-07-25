package client

import (
	"database/sql"
	"errors"
	"strings"
	"testing"
)

func TestPgvectorCheckErrorReportsMissingExtension(t *testing.T) {
	err := pgvectorCheckError(sql.ErrNoRows)
	if err == nil {
		t.Fatal("expected error")
	}
	if !strings.Contains(err.Error(), "pgvector extension is not initialized") {
		t.Fatalf("unexpected error: %v", err)
	}
}

func TestPgvectorCheckErrorPreservesDatabaseFailure(t *testing.T) {
	cause := errors.New("password authentication failed")
	err := pgvectorCheckError(cause)
	if err == nil {
		t.Fatal("expected error")
	}
	if !errors.Is(err, cause) {
		t.Fatalf("expected wrapped cause, got %v", err)
	}
	if strings.Contains(err.Error(), "CREATE EXTENSION") {
		t.Fatalf("database failure should not be reported as missing extension: %v", err)
	}
}
