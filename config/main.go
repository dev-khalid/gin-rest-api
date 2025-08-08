package config

func init() {
	loadEnvVariables()
	initDB()
}
