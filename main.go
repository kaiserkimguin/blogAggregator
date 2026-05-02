package main

import(
	"fmt"
	"encoding/json"
	"github.com/kaiserkimguin/blogAggregator/internal/config"
)

func main()  {
	// read the config file
	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}
	// set the username to current user 
	err = cfg.SetUser("kaiserkimguin")
	// read the config file again
	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}
	// Marshal the cfg for better readability, then print it.
	data, err := json.MarshalIndent(cfg, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}
