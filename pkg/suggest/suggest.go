package suggest

//EnvironmentInfo contains information about the local environment
type EnvironmentInfo struct {
	gitInfo GitInfo
}

//GitInfo contains information about whether the current directory
//is part of a repository and other related information
type GitInfo struct {
	isRepo bool
}

//SuggestCommand returns a command suggestion based on the history of commands
func SuggestCommand(commandHistory []string) string {
	return commandHistory[len(commandHistory)-1]
}

func collectEnvironmentInfo() EnvironmentInfo {
	env := EnvironmentInfo{}
	return env
}
