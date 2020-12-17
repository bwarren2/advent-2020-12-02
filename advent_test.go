package advent_test

import (
	"advent"
	"fmt"
	"testing"
)

func TestValidPasswordCount(t *testing.T) {
	fmt.Print(advent.ValidPasswordCount("input.txt"))
}
func TestNewValidPasswordCount(t *testing.T) {
	fmt.Print(advent.NewValidPasswordCount("input.txt"))
}
