package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var requiredParams = []string{
	"KAPITOL_MONGO_URL",
	"KAPITOL_MONGO_DB",
	"KAPITOL_MONGO_PASSWORD",
	"KAPITOL_MONGO_MEMBERS_COLLECTION",
	"KAPITOL_MONGO_SENATORS_COLLECTION",
	"KAPITOL_MONGO_LEGISLATION_COLLECTION",
	"OCTOPUS_LOG_PATH",
	"OCTOPUS_LOG_LEVEL",
	"KAPITOL_PRO_PUBLICA_CONGRESS_API_KEY",
}

//Mongo - Config strings related to the mongo db
type Mongo struct {
	Url                   string
	Db                    string
	Password              string
	MembersCollection     string
	SenatorsCollection    string
	LegislativeCollection string
}

//LogInfo - Config strings related to the logger
type LogInfo struct {
	Path string
	Level int
}

//ApiKeys - list of all the api keys
type ApiKeys struct {
	ProPublicaCongress string
}

//Config - Octopus config
type Config struct {
	Mongo Mongo
	LogInfo LogInfo
	ApiKeys ApiKeys
}

//C - Global config variable
var C Config

func init() {
	missingEnvVars := make([]string, 0, len(requiredParams))
	for _, v := range requiredParams {
		val := os.Getenv(v)
		if val == "" {
			missingEnvVars = append(missingEnvVars, v)
		}
	}

	if len(missingEnvVars) > 0 {
		panic("Octopus environment variables not set properly. Missing:\n" + strings.Join(missingEnvVars, "\n"))
	}

	level, err := strconv.Atoi(os.Getenv("OCTOPUS_LOG_LEVEL"))
	if err != nil {
		fmt.Println("Error: getting octopus log level:", err)
		level = 0
	}

	C = Config{
		Mongo: Mongo{
			Url:                   os.Getenv("KAPITOL_MONGO_URL"),
			Db:                    os.Getenv("KAPITOL_MONGO_DB"),
			Password:              os.Getenv("KAPITOL_MONGO_PASSWORD"),
			MembersCollection:     os.Getenv("KAPITOL_MONGO_MEMBERS_COLLECTION"),
			SenatorsCollection:    os.Getenv("KAPITOL_MONGO_SENATORS_COLLECTION"),
			LegislativeCollection: os.Getenv("KAPITOL_MONGO_LEGISLATION_COLLECTION"),
		},
		LogInfo: LogInfo{Path: os.Getenv("OCTOPUS_LOG_PATH"), Level: level},
		ApiKeys: ApiKeys{ProPublicaCongress: os.Getenv("KAPITOL_PRO_PUBLICA_CONGRESS_API_KEY")},
	}
}
