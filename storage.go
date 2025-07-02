package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"sync"
)

const FileName = "urls.json"

var (
	ErrNotFound = errors.New("short code not found")
)

type Storage interface {
	Store(shortCode string, longUrl string) (string, error)
	Get(url string) (string, error)
}

type URLStorage struct {
	filename string
	mutex    sync.RWMutex
}

func NewStorage() (*URLStorage, error) {
	s := &URLStorage{
		filename: FileName,
	}

	if _, err := os.Stat(s.filename); os.IsNotExist(err) {
		if err := os.WriteFile(s.filename, []byte("{}"), 0666); err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
	}

	return s, nil
}

func (s *URLStorage) Store(shortCode string, longURL string) (string, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data, err := s.readFile()
	if err != nil {
		return "", err
	}

	for i := 0; ; i++ {
		if i > 0 {
			shortCode = fmt.Sprintf("%s_%d", shortCode, i)
		}

		if _, found := data[shortCode]; !found {
			data[shortCode] = longURL
			if err := s.writeFile(data); err != nil {
				return "", err
			}
			return shortCode, nil
		}
	}
}

func (s *URLStorage) Get(shortCode string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	data, err := s.readFile()
	if err != nil {
		return "", err
	}

	longURL, found := data[shortCode]
	if !found {
		return "", ErrNotFound
	}

	return longURL, nil
}

func (s *URLStorage) readFile() (map[string]string, error) {
	data := make(map[string]string)
	file, err := os.ReadFile(s.filename)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(file) == 0 {
		return data, nil
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return data, nil
}

func (s *URLStorage) writeFile(data map[string]string) error {
	bytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	tempFile := s.filename + ".tmp"
	if err := os.WriteFile(tempFile, bytes, 0666); err != nil {
		return fmt.Errorf("failed to write temp file: %w", err)
	}

	if err := os.Rename(tempFile, s.filename); err != nil {
		return fmt.Errorf("failed to rename temp file: %w", err)
	}

	return nil
}
