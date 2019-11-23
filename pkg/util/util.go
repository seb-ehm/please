package util

//SuggestCommand returns a command suggestion based on the history of commands
func SuggestCommand(commandHistory []string) string {
	return commandHistory[len(commandHistory)-1]
}
