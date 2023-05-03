package database

import "time"

var (
	errorCreateDatabaseService = map[string]string{
		"message":   "unable to create database service",
		"timestamp": time.Now().Local().Format("2006-01-02T15:04:05Z07:00"),
	}
)
