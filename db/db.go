package db

import (
	"runtime"

	"github.com/Sirupsen/logrus"
	"github.com/jmoiron/sqlx"

	"path"

	_ "github.com/lib/pq"
	migrate "github.com/rubenv/sql-migrate"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func GetDB() *sqlx.DB {
	if db != nil {
		return db
	}

	logrus.SetFormatter(new(logrus.JSONFormatter))
	// build database connection
	db2, err := sqlx.Connect("postgres", "host="+viper.GetString("postgresql_addon_host")+" port="+viper.GetString("postgresql_addon_port")+" database="+viper.GetString("postgresql_addon_db")+" user="+
		viper.GetString("postgresql_addon_user")+" password="+viper.GetString("postgresql_addon_password")+
		" sslmode=disable")

	if err != nil {
		logrus.WithField("error", err).WithField("conn", "host="+viper.GetString("postgresql_addon_host")+" port="+viper.GetString("postgresql_addon_port")+" database="+viper.GetString("postgresql_addon_db")+" user="+
			viper.GetString("postgresql_addon_user")+" password="+viper.GetString("postgresql_addon_password")+
			" sslmode=disable").Fatalln("can't connect to db")
	}
	db = db2

	_, filename, _, _ := runtime.Caller(0)

	migrations := &migrate.FileMigrationSource{
		Dir: path.Dir(filename) + "/migrations",
	}

	n, err := migrate.Exec(db.DB, "postgres", migrations, migrate.Up)

	if err != nil {
		logrus.WithField("error", err).Fatalln("can't apply migrations")
	}

	logrus.WithField("MigrationsApplied", n).Infoln("applied migrations")

	return db
}

func SetDB(newdb *sqlx.DB) {
	db = newdb
}
