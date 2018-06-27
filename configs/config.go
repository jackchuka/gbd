package configs

import "flag"

// Configs contain all configs
type Configs struct {
	DSN    string
	Driver string
}

// GetConfigs returns config
func GetConfigs() Configs {
	dsn := flag.String("dsn", "user:pass@tcp(127.0.0.1:3306)/database?parseTime=true", "Specify datasource")
	driver := flag.String("driver", "mysql", "Database driver")
	flag.Parse()

	config := Configs{
		DSN:    *dsn,
		Driver: *driver,
	}
	return config
}
