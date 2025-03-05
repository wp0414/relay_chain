package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var (
	dbFilePath = "./"
)

type LevelDB struct {
	Db *leveldb.DB
}

func GetDB(dbFileName string) *LevelDB {
	db, err := leveldb.OpenFile(dbFilePath+dbFileName, nil)
	if err != nil {
		log.Fatalf("create or load leveldb fail [%s]", err)
	}
	return &LevelDB{Db: db}
}

func (ld *LevelDB) Put(key string, value string) error {
	return ld.Db.Put([]byte(key), []byte(value), nil)
}
func (ld *LevelDB) Get(key string) (string, error) {
	data, err := ld.Db.Get([]byte(key), nil)
	return string(data), err
}
func (ld *LevelDB) Delete(key string) error {
	return ld.Db.Delete([]byte(key), nil)
}

func (ld *LevelDB) Close() error {
	return ld.Db.Close()
}

const (
	ReqSuffix  = "req"
	RespSuffix = "resp"
)

func (ld *LevelDB) PutRequest(crossId string, request string) error {
	return ld.Put(crossId+ReqSuffix, request)
}

func (ld *LevelDB) GetRequest(crossId string) (string, error) {
	return ld.Get(crossId + ReqSuffix)
}

func (ld *LevelDB) PutResponse(crossId string, response string) error {
	return ld.Put(crossId+RespSuffix, response)
}

func (ld *LevelDB) GetResponse(crossId string) (string, error) {
	return ld.Get(crossId + RespSuffix)
}
