package configs

import "github.com/spf13/viper"

//3 structs: servidor, banco de dados porta q vai levantar, unir as structs

var cfg *config

type config struct {
	API APIConfig
	DB  DBConfig
}

type APIConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Pass     string
	Database string
}

func init() { //sempre é chamada no start das aplicações
	viper.SetDefault("api.port", "9000") //viper pra definir valores padroes pras nossas configurações, seria o default caso nao configure nada no load
	viper.SetDefault("database.host", "localhost")
	viper.SetDefault("database.port", "5432") //porta padrao do postgreSQL
}

func Load() error {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}

	cfg = new(config) // new é igual a &cconfig{}, vai criar um ponteiro da struct

	cfg.API = APIConfig{
		Port: viper.GetString("api.port"),
	}

	cfg.DB = DBConfig{
		Host:     viper.GetString("database.host"),
		Port:     viper.GetString("database.port"),
		User:     viper.GetString("database.user"),
		Pass:     viper.GetString("database.pass"),
		Database: viper.GetString("database.name"),
	}
	return nil
}

//função pra retornar objeto de configuração do banco e outra para porta

func GetDB() DBConfig {
	return cfg.DB
}

func GetServerPort() string {
	return cfg.API.Port
}
