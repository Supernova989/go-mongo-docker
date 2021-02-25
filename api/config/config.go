package config

import (
	"github.com/spf13/viper"
	"log"
)

type Configuration struct {
	Environment string
	Mongo       MongoConfiguration
}

type MongoConfiguration struct {
	Server      string
	Database    string
	Collections MongoCollections
}

type MongoCollections struct {
	Users string
	Posts string
}

var GetConfig = func() Configuration {
	conf := Configuration{}
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = viper.Unmarshal(&conf)
	if err != nil {
		log.Panic(err)
	}
	collections := MongoCollections{
		Users: "Users",
		Posts: "Posts",
	}
	conf.Mongo.Collections = collections
	return conf
}
