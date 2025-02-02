package commands

import (
	"fmt"

	"github.com/wolv89/bootgator/internal/state"
)

type Command struct {
	Args []string
	Name string
}

type Commands struct {
	List map[string]func(*state.State, Command) error
}

func (cmds *Commands) Register(name string, f func(*state.State, Command) error) error {

	if len(name) == 0 {
		return fmt.Errorf("command name cannot be empty")
	}

	if _, ok := cmds.List[name]; ok {
		return fmt.Errorf("%s is already registered as a command", name)
	}

	cmds.List[name] = f
	// fmt.Printf("%s command registered\n", name)

	return nil

}

func (cmds *Commands) Run(s *state.State, cmd Command) error {

	if _, ok := cmds.List[cmd.Name]; !ok {
		return fmt.Errorf("unknown command: %s", cmd.Name)
	}

	return cmds.List[cmd.Name](s, cmd)

}
