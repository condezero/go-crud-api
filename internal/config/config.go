package config

import "os"

type Config struc {
	Environment string
	ConnectionString string
}

func getValue(envName string, defaultValue string) string {
	if val, ok := os.LookupEnv(envName); ok {
		return val
	}
	return defaultValue
}
func NewConfig() *Config {
	return &Config{
		Environment: getValue("ENV", "local")
		ConnectionString: getValue("DB_Conn", "User ID=root;Password=myPassword;Host=localhost;Port=5432;Database=myDataBase;Pooling=true;Min Pool Size=0;Max Pool Size=100;Connection Lifetime=0;")
	}
}