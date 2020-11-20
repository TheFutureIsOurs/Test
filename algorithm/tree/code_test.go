package tree

import (
	"fmt"
	"testing"
)

func TestSeRe(t *testing.T) {
	tree := BuildTree()
	PrintTree(tree)

	str := Serialize(tree)
	fmt.Println(str)
	tree = UnSerialize(str)
	PrintTree(tree)

}
