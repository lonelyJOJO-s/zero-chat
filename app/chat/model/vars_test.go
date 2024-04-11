package model

import (
	"fmt"
	"testing"
)

func TestColumns(t *testing.T) {
	columns := NewColumns(WithContent("welcome"), WithType(1))
	fmt.Println(columns)
}
