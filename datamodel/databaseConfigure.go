package datamodel

import "path"

var (
	DatabaseDir  = "./config"
	DatabaseName = "main.sqlite"
	DatabasePath = path.Join(DatabaseDir, DatabaseName)
)
