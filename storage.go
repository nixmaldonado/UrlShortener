package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
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

type URLEntry struct {
	URL           string `json:"url"`
	RedirectCount int    `json:"redirect_count"`
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
			data[shortCode] = URLEntry{
				URL:           longURL,
				RedirectCount: 0,
			}
			if err := s.writeFile(data); err != nil {
				return "", err
			}
			return shortCode, nil
		}
	}
}

func (s *URLStorage) Get(shortCode string) (URLEntry, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	entry := URLEntry{}

	data, err := s.readFile()
	if err != nil {
		return entry, err
	}

	log.Printf("Map contents: %v\n", data)

	entry, found := data[shortCode]
	if !found {
		return entry, ErrNotFound
	}

	return entry, nil
}

func (s *URLStorage) IncrementCounter(shortCode string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	data, err := s.readFile()
	if err != nil {
		return fmt.Errorf("failed to read data: %w", err)
	}

	entry, found := data[shortCode]
	if !found {
		return ErrNotFound
	}

	entry.RedirectCount++
	data[shortCode] = entry

	if err := s.writeFile(data); err != nil {
		return fmt.Errorf("failed to write updated data: %w", err)
	}

	return nil
}

func (s *URLStorage) readFile() (map[string]URLEntry, error) {
	data := make(map[string]URLEntry)
	file, err := os.ReadFile(s.filename)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("File does not exist, initializing empty storage")
			return data, nil
		}
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	if len(file) == 0 {
		log.Println("File is empty")
		return data, nil
	}

	if err := json.Unmarshal(file, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal file: %w", err)
	}

	return data, nil
}

func (s *URLStorage) writeFile(data map[string]URLEntry) error {
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
