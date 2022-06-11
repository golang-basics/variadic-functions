package main

var commands = map[string]Command{}

func main() {
	// create new command
	command := cmd("get", "get todo by id", "id")
	command.Execute()

	// get the existing command
	command = cmd("get")
	command.Execute()
}

func cmd(args ...string) Command {
	name, description, arguments := "", "", make([]string, 0)
	if len(args) > 0 {
		name = args[0]
	}
	c, ok := commands[name]
	if ok {
		return c
	}

	if len(args) > 1 {
		description = args[1]
	}
	if len(args) > 2 {
		arguments = args[2:]
	}
	c = Command{
		Name:        name,
		Description: description,
		Args:        arguments,
	}
	commands[name] = c
	return c
}

type Command struct {
	Name        string
	Description string
	Args        []string
}

func (c Command) Execute() {
}
