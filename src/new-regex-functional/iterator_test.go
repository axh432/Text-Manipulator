package new_regex

import (
	"fmt"
	"testing"
)

func Test_Iterator(t *testing.T) {
	iter := CreateIterator("F (l r i){+}")
	for iter.HasNext() {
		fmt.Printf("%s", string(iter.Next()))
	}
}
