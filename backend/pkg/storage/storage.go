package storage

import (
	cache "distributed_calculator/pkg/cache"
	ex "distributed_calculator/pkg/expression"
	"encoding/json"
	"os"
	"time"
)

type Storage struct {
	filepath string
	ch       *cache.Cache
	interval time.Duration
	stop     chan struct{}
}

func NewStorage(fp string, syncInterval time.Duration) *Storage {

	storage := &Storage{
		filepath: fp,
		ch:       cache.NewCache(),
		interval: syncInterval,
		stop:     make(chan struct{}),
	}

	fileData, err := storage.loadFileData()
	if err != nil {
	} else {
		storage.ch.SetData(fileData)
	}

	if syncInterval > 0 {
		go storage.syncWithFile()
	}

	return storage
}

func (s *Storage) GetExpressionByID(exprID string) (ex.Expression, bool) {
	expr, found := s.ch.GetExpressionByID(exprID)

	if found {
		return expr, found
	}

	fileData, err := s.loadFileData()

	if err != nil {
		return ex.Expression{}, false
	}

	s.ch.SetData(fileData)
	expr, found = s.ch.GetExpressionByID(exprID)
	return expr, found
}

// function open file and add new data to cache
func (s *Storage) loadFileData() (map[string]ex.Expression, error) {
	fileData := make(map[string]ex.Expression)

	file, err := os.Open(s.filepath)
	if err != nil {
		return fileData, err
	}

	decoder := json.NewDecoder(file)
	decoder.Decode(&fileData)

	if fileData == nil {
		return make(map[string]ex.Expression), nil
	}
	return fileData, nil
}

// function add new expression to cache and to data file
// every 5 seconds
func (s *Storage) AddExpression(newExpr ex.Expression) {
	s.ch.AddExpression(newExpr)
}

func (s *Storage) syncWithFile() {
	ticker := time.NewTicker(s.interval)
	for {
		select {
		case <-ticker.C:
			s.dumpToFile()
		case <-s.stop:
			ticker.Stop()
			return
		}
	}
}

func (s *Storage) dumpToFile() {
	file, err := os.OpenFile(s.filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		return
	}
	fileData := make(map[string]ex.Expression)
	decoder := json.NewDecoder(file)
	decoder.Decode(&fileData)

	for k, v := range s.ch.Data() {
		if fileData[k] != v {
			fileData[k] = v
		}
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(fileData)
}

func (s *Storage) Update(expr ex.Expression) {
	s.ch.Update(expr)
}

func (s *Storage) Data() map[string]ex.Expression {
	return s.ch.Data()
}
