package main

import(
	"fmt"
	"os"
	"github.com/kaiserkimguin/blogAggregator/internal/config"
 _ "github.com/lib/pq"
)

func main()  {
	// initialize a commands, and a state struct to work with 
	cmds := commands{
		cmdMap: make(map[string]func(*state, command)error),
	}
	var s state 
	// read the config file and store it in s
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s.cfg = &cfg
	// load the dbURL from State and connect into the database
	db, err := sql.Open("postgres", dbURL)
	// store the dbQueries in the state struct.
	dbQueries := database.New(db)
	s.db = &dbQueries
	// register all needed functions
	cmds.register("login", handlerLogin)
	// detect all arguments provided by the caler and construct
	// new command struct with it.
	args := os.Args	
	if len(args) < 2 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	}
	cmd := command{
		name: 	args[1],
		args: args[2:],
	}
	err = cmds.run(&s, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

type state struct {
	db 			*database.Queries
	cfg			*config.ConfigJson
}
