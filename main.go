package main

import (
	"flag"
	"gu2list/ent"
	"gu2list/utils"
)

var (
	SqlClient *ent.Client
	Token     = flag.String("t", "", "Input Your Bot Token")
)

func init() {
	flag.Parse()
}

func main() {
	/*

		var err error

		//Choose Database on your own
		//Here is Example Using SqlClient Database.
		//Testing env Using XAMPP , SqlClient
		//Connect To Database

		SqlClient, err = database.CreateClient("root", "", "test", "localhost", 3306)
		if err != nil {
			panic(err)
		}

	*/
	utils.DiscordClient(*Token)
}
