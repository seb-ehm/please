package command

import "strings"

//EnvironmentInfo contains information about the local environment
type EnvironmentInfo struct {
	gitInfo GitInfo
}

//GitInfo contains information about whether the current directory
//is part of a repository and other related information
type GitInfo struct {
	isRepo bool
}

//Suggest returns a command suggestion based on the history of commands
func Suggest(commandHistory []string) string {
	return commandHistory[len(commandHistory)-1]
}

func collectEnvironmentInfo() EnvironmentInfo {
	env := EnvironmentInfo{}
	return env
}

//Details is an interface implemented for all commands corrected by please
type Details interface {
	Name() string
	Misspellings() []string
	Correct(command Command) Command
}

//Command is a structured representations of Commands
type Command struct {
	Sudo      bool
	Name      string
	Arguments []string
}

func (c Command) String() string {
	var sb strings.Builder
	if c.Sudo {
		sb.WriteString("sudo ")
	}
	sb.WriteString(c.Name)
	sb.WriteRune(' ')
	sb.WriteString(strings.Join(c.Arguments[:], " "))
	return sb.String()
}
