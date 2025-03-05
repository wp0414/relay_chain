package server

import (
	"relay-chain/module/db"
	"relay-chain/module/plugin"
	"relay-chain/module/verifier"
)

var (
	DefaultServer Server
)

type Server struct {
	Plugin   plugin.ChainPlugin
	DB       *db.LevelDB
	Verifier *verifier.Verifier
	Mode     string
}
