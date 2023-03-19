package __

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)

// TestCopyText tests copy a text file a line at a time, with optional filtering
func TestCopyText(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if line != "unwanted" {
			fmt.Println(line)
		}
	}
	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
