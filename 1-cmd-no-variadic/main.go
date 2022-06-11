package main

var commands = map[string]Command{}

func main() {
	// create new command
	command := cmd("get", "get todo by id", []string{"id"})
	command.Execute()

	// get the existing command
	command = cmd("get", "", nil)
	command.Execute()
}

func cmd(name, description string, args []string) Command {
	c, ok := commands[name]
	if ok {
		return c
	}
	c = Command{
		Name:        name,
		Description: description,
		Args:        args,
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
