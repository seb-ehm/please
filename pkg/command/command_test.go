package command

import "testing"

func TestString(t *testing.T) {
	c := Command{true, "apt", []string{"install", "expect"}}
	t.Run("sudo apt install expect", testStringfunc(c, "sudo apt install expect"))
	c = Command{false, "ls", []string{"."}}
	t.Run("ls .", testStringfunc(c, "ls ."))
}

func testStringfunc(c Command, expected string) func(*testing.T) {
	return func(t *testing.T) {
		actual := c.String()
		if actual != expected {
			t.Errorf("Expected: \"%s\" but instead got \"%s\"", expected, actual)
		}
	}
}
