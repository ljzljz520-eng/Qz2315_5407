package main

import "time"

type Config struct {
	Addr, DBPath string
	Shutdown     time.Duration
}

func DefaultConfig() Config {
	return Config{Addr: ":8080", DBPath: "journal.db", Shutdown: 5 * time.Second}
}
func ValidConfig(c Config) bool { return c.Addr != "" && c.DBPath != "" && c.Shutdown > 0 }
