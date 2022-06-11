package main

var commands = map[string]Command{}

func main() {
	// create new command
	command := createCmd("get", "get todo by id", []string{"id"})
	command.Execute()

	// get the existing command
	command = getCmd("get")
	command.Execute()
}

func createCmd(name, description string, args []string) Command {
	c := Command{
		Name:        name,
		Description: description,
		Args:        args,
	}
	commands[name] = c
	return c
}

func getCmd(name string) Command {
	c, ok := commands[name]
	if ok {
		return c
	}
	return Command{}
}

type Command struct {
	Name        string
	Description string
	Args        []string
}

func (c Command) Execute() {
}
