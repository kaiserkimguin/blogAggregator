package main 

import(
	"errors"
)

type commands struct {
	cmdMap	map[string]func (*state, command) error
}
func (c *commands) run(s *state, cmd command) error {
	 // check wether func exists
	 handler, ok := c.cmdMap[cmd.name]
	 if !ok {
		 return errors.New("unknown command")
	 }
		// return what the handler returns, if the function call was successul
		// the return will be nill, else the function will return the handlers 
		// error 
		return handler(s, cmd) 
 }

func (c *commands) register(name string,f func(s *state, cmd command)error) {
// simply take the func and store it in c's comannd map
// attention: no checks implemented here. This function will overwrite if not careful
c.cmdMap[name] = f 
} 

type command struct {
	name 		string
	args		[]string
}


