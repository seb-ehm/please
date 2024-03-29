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
func Suggest(command Command) string {

	return command.String()
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

//New creates a new Command from a string
func New(s string) Command {
	sa := strings.Split(s, " ")

	c := Command{}
	i := 0
	if len(sa) > 1 {
		if sa[i] == "sudo" {
			c.Sudo = true
			i++
		}
		c.Name = sa[i]
		i++
		c.Arguments = sa[i:]

	} else {
		c.Sudo = false
		c.Name = sa[0]
	}
	return c

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

//Equals returns whether two Commands are identical
func (c *Command) Equals(other *Command) bool {
	if c.Sudo != other.Sudo {
		return false
	}
	if c.Name != other.Name {
		return false
	}
	if len(c.Arguments) != len(other.Arguments) {
		return false
	}
	for i := 0; i < len(c.Arguments); i++ {
		if c.Arguments[i] != other.Arguments[i] {
			return false
		}
	}
	return true
}
