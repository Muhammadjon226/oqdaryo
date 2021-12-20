package main

import (
	"fmt"
)

// Postgres'ga ulanish sozlamalarini e'lon qilamiz
var (
	host = "localhost"
	user = "muhammad"
	password = "12345"
	name = "oqdaryo"
	port = 5432
)

// Configdan foydalanib "connection string" yasab olamiz
var DB_CONFIG = fmt.Sprintf(
	"host=%s user=%s password=%s dbname=%s port=%d sslmode = disable",
	host, user, password, name, port,
)
