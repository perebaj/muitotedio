package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseCSV(t *testing.T) {
	tmpfile, err := os.CreateTemp("", "test.csv")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	testData := []string{"question1,1", "question2,2", "question3,3"}
	for _, data := range testData {
		_, err := tmpfile.WriteString(data + "\n")
		if err != nil {
			t.Fatalf("Failed to write to temporary file: %v", err)
		}
	}

	err = tmpfile.Close()
	if err != nil {
		t.Fatalf("Failed to close temporary file: %v", err)
	}

	got, err := parseCSV(tmpfile.Name())
	if err != nil {
		t.Fatalf("Failed to parse CSV: %v", err)
	}

	want := []Problem{
		{"question1", 1},
		{"question2", 2},
		{"question3", 3},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("Unexpected problems. Got %v, want %v", got, want)
	}
}
