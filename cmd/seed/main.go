package main

import (
	"apart_community/config"
	"apart_community/database"
	"apart_community/database/seeds"
)

func main() {
	flags := config.ParseFlags()
	env := config.LoadEnv(&flags)

	db := database.ConnectToPostgres(env)

	if flags.Seed {
		seeds.ApartmentSeed(db)
		seeds.UserSeed(db)
	}
}
