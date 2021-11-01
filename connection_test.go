package go_crud

import (
	"fmt"
	"testing"
)

func TestConnection(t *testing.T) {
	conn := GetConnection()

	defer conn.Close()
	fmt.Println("Done")
}
