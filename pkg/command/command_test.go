package command

import (
	"testing"
)

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

func TestCommandEquals(t *testing.T) {

	testCases := []struct {
		com    Command
		other  Command
		result bool
	}{
		{Command{true, "apt", []string{"install", "expect"}},
			Command{true, "apt", []string{"install", "expect"}},
			true},

		{Command{true, "apt", []string{"install", "expect"}},
			Command{false, "apt", []string{"install", "expect"}},
			false},
	}
	for _, tc := range testCases {
		if !tc.com.Equals(&tc.other) == tc.result {
			if tc.result {
				t.Errorf("Expected \"%s\" to Equal \"%s\" ", tc.com, tc.other)
			} else {
				t.Errorf("Expected \"%s\" to not Equal \"%s\" ", tc.com, tc.other)
			}

		}
	}

}
