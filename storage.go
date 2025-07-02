package main

type URLStorage interface {
	Save(shortCode string, longUrl string) error
	Get(shortCode string) error
}

type Storage struct {
	store map[string]string
}

func NewStorage() *Storage {
	return &Storage{make(map[string]string)}
}

func (s *Storage) Save(shortCode string, longUrl string) error {
	v, exists := s.store[shortCode]
	if exists {
		if v == longUrl {
			return nil
		}

	}

}
