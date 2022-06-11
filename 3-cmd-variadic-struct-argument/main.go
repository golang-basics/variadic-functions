package main

var commands = map[string]Command{}

func main() {
	// create new command
	command := cmd(Args{
		Name:        "get",
		Description: "get todo by id",
		Args:        []string{"id"},
	})
	command.Execute()

	// get the existing command
	command = cmd(Args{Name: "get"})
	command.Execute()
}

type Args struct {
	Name        string
	Description string
	Args        []string
}

func cmd(args Args) Command {
	c, ok := commands[args.Name]
	if ok {
		return c
	}

	c = Command{
		Name:        args.Name,
		Description: args.Description,
		Args:        args.Args,
	}
	commands[args.Name] = c
	return c
}

type Command struct {
	Name        string
	Description string
	Args        []string
}

func (c Command) Execute() {
}
