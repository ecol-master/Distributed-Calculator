package storage

import (
	cache "distributed_calculator/pkg/cache"
	ex "distributed_calculator/pkg/expression"
	"encoding/json"
	"log"
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
func (s *Storage) loadFileData() (cache.CacheData, error) {
	fileData := make(cache.CacheData)

	file, err := os.Open(s.filepath)
	if err != nil {
		return fileData, err
	}

	decoder := json.NewDecoder(file)
	decoder.Decode(&fileData)

	if fileData == nil {
		return make(cache.CacheData), nil
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
	flags := log.Ldate | log.Ltime | log.Lshortfile
	syncLogFile, err := os.OpenFile("../data/sync_storage.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		os.Exit(1)
	}
	syncLogger := log.New(syncLogFile, "sync_storage", flags)
	syncLogger.Printf("Initialize syncLogger")
	for {
		select {
		case <-ticker.C:
			s.dumpToFile(syncLogger)
		case <-s.stop:
			ticker.Stop()
			return
		}
	}
}

func (s *Storage) dumpToFile(logger *log.Logger) {
	logger.Printf("start dumping new entities to dump storage file")
	file, err := os.OpenFile(s.filepath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return
	}
	fileData := make(cache.CacheData)
	decoder := json.NewDecoder(file)
	decoder.Decode(&fileData)
	logger.Println("get from file", fileData)

	var newEntities, changedEntities int
	for k, v := range s.ch.Data() {
		value, found := fileData[k]
		if !found {
			fileData[k] = v
			newEntities++
			continue
		}
		if value.Status != v.Status {
			fileData[k] = value
			changedEntities++
		}
	}
	// don't update storage if nothing changed
	if newEntities == 0 && changedEntities == 0 {
		return
	}

	encoder := json.NewEncoder(file)
	encoder.Encode(fileData)
	logger.Printf("write to dump storage file new_entities=%d, changed_entities=%d", newEntities, changedEntities)
}

func (s *Storage) Update(expr ex.Expression) {
	s.ch.Update(expr)
}

func (s *Storage) Data() cache.CacheData {
	return s.ch.Data()
}
