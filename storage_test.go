package main

import (
	"encoding/json"
	"errors"
	"os"
	"strings"
	"testing"
)

func setupTestStorage(t *testing.T) (*URLStorage, string) {
	t.Helper()
	tempDir := t.TempDir()
	tempFile := tempDir + "/test_urls.json"
	s := &URLStorage{
		filename: tempFile,
	}

	if err := os.WriteFile(tempFile, []byte("{}"), 0666); err != nil {
		t.Fatalf("Failed to initialize temp file: %v", err)
	}

	return s, tempFile
}

func setupTestStorageWithData(t *testing.T) (*URLStorage, string) {
	t.Helper()
	tempDir := t.TempDir()
	tempFile := tempDir + "/test_urls.json"
	s := &URLStorage{
		filename: tempFile,
	}

	initialData := map[string]URLEntry{
		"abc123": {URL: "https://example.com", RedirectCount: 5},
	}
	bytes, err := json.MarshalIndent(initialData, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal initial data: %v", err)
	}
	if err := os.WriteFile(tempFile, bytes, 0666); err != nil {
		t.Fatalf("Failed to write initial data to temp file: %v", err)
	}

	return s, tempFile
}

func TestSuccessfulStorage(t *testing.T) {
	s, tempFile := setupTestStorage(t)
	defer os.Remove(tempFile) // Clean up

	shortCode, err := s.Store("abc123", "https://example.com")
	if err != nil {
		t.Errorf("Store failed: %v", err)
	}

	if shortCode != "abc123" {
		t.Errorf("Expected short code 'abc123', got %q", shortCode)
	}

	data, err := os.ReadFile(tempFile)
	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	var storedData map[string]URLEntry

	if err := json.Unmarshal(data, &storedData); err != nil {
		t.Fatalf("Failed to unmarshal data: %v", err)
	}

	entry, found := storedData["abc123"]
	if !found {
		t.Error("Expected entry with short code 'abc123' not found")
	}

	if entry.URL != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got %q", entry.URL)
	}

	if entry.RedirectCount != 0 {
		t.Errorf("Expected RedirectCount 0, got %d", entry.RedirectCount)
	}
}

func TestCollisionHandling(t *testing.T) {
	s, tempFile := setupTestStorage(t)
	defer os.Remove(tempFile) // Clean up

	_, err := s.Store("abc123", "https://example1.com")
	if err != nil {
		t.Fatalf("First Store failed: %v", err)
	}

	newShortCode, err := s.Store("abc123", "https://example2.com")

	if err != nil {
		t.Errorf("Second Store failed: %v", err)
	}

	if newShortCode == "abc123" {
		t.Error("Expected a new short code due to collision, got the same")
	}

	if !strings.HasPrefix(newShortCode, "abc123_") {
		t.Errorf("Expected new short code to start with 'abc123_', got %q", newShortCode)
	}

	data, err := os.ReadFile(tempFile)

	if err != nil {
		t.Fatalf("Failed to read temp file: %v", err)
	}

	var storedData map[string]URLEntry

	if err := json.Unmarshal(data, &storedData); err != nil {
		t.Fatalf("Failed to unmarshal data: %v", err)
	}

	if _, found := storedData["abc123"]; !found {
		t.Error("Original entry 'abc123' not found")
	}

	if _, found := storedData[newShortCode]; !found {
		t.Errorf("New entry %q not found", newShortCode)
	}
}

func TestSuccessfulRetrieval(t *testing.T) {
	s, tempFile := setupTestStorageWithData(t)
	defer os.Remove(tempFile) // Clean up

	entry, err := s.Get("abc123")
	if err != nil {
		t.Errorf("Get failed: %v", err)
	}

	if entry.URL != "https://example.com" {
		t.Errorf("Expected URL 'https://example.com', got %q", entry.URL)
	}

	if entry.RedirectCount != 5 {
		t.Errorf("Expected RedirectCount 5, got %d", entry.RedirectCount)
	}
}

func TestNotFound(t *testing.T) {
	s, tempFile := setupTestStorageWithData(t)
	defer os.Remove(tempFile) // Clean up

	_, err := s.Get("nonexistent")
	if err == nil {
		t.Error("Expected ErrNotFound for nonexistent short code, got nil")
	}

	if !errors.Is(err, ErrNotFound) {
		t.Errorf("Expected error to be ErrNotFound, got %v", err)
	}
}
