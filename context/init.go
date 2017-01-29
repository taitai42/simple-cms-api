// Package context initialize component required by the server
// like database, run migration etc...
// it basically make sure the server can run and use every component safely
package context

import (
	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/spf13/viper"
	"github.com/taitai42/simple-cms-api/db"
)


//Init initialize various components for the server
func Init() {

	if viper.GetBool("debug") == true {
		log.SetLevel(log.DebugLevel)
	}

	log.SetFormatter(&log.JSONFormatter{})

	db.GetDB() // launch migrations

	token := viper.GetString("tokens")

	tokens := strings.Split(token, ",")
	viper.Set("tokens", tokens)
}
