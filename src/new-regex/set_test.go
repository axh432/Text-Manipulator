package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSet(t *testing.T) {
	t.Run("when given a string that matches an expression in a set return true", func(t *testing.T) {
		iterA := CreateIterator("a")
		iterB := CreateIterator("b")
		iterC := CreateIterator("c")

		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)

		matchResultA := MatchIter(&iterA, exp)
		expectedA := MatchTree{
			IsValid:   true,
			Value:     "a",
			Type:	   "Set",
			Label:     "",
			Children:  []MatchTree{{IsValid: true, Type:"SetOfCharacters", Value: "a"}},
			DebugLine: "",
		}
		require.Equal(t, expectedA, matchResultA)
		require.Equal(t, 1, iterA.index)

		matchResultB := MatchIter(&iterB, exp)
		expectedB := MatchTree{
			IsValid:   true,
			Value:     "b",
			Type:	   "Set",
			Label:     "",
			Children:  []MatchTree{{IsValid: true, Type:"SetOfCharacters", Value: "b"}},
			DebugLine: "",
		}
		require.Equal(t, expectedB, matchResultB)
		require.Equal(t, 1, iterB.index)

		matchResultC := MatchIter(&iterC, exp)
		expectedC := MatchTree{
			IsValid:   true,
			Value:     "c",
			Type:	   "Set",
			Label:     "",
			Children:  []MatchTree{{IsValid: true, Type:"SetOfCharacters", Value: "c"}},
			DebugLine: "",
		}
		require.Equal(t, expectedC, matchResultC)
		require.Equal(t, 1, iterC.index)
	})

	t.Run("when given a character that is not in the set return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)
		iter := CreateIterator("d")

		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:	   "Set",
			Label:     "",
			Children:  nil,
			DebugLine: "Set:[], NoMatch:string does not match the given subexpressions",
		}
		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given a string and the first character matches return true", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)
		iter := CreateIterator("athguy")

		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   true,
			Value:     "a",
			Type:	   "Set",
			Label:     "",
			Children:  []MatchTree{{IsValid: true, Type:"SetOfCharacters", Value: "a"}},
			DebugLine: "",
		}
		require.Equal(t, expected, matchResult)
		require.Equal(t, 1, iter.index)
	})

	t.Run("when given a string and the first character does not match return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)
		iter := CreateIterator("xthguy")

		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:	   "Set",
			Label:     "",
			Children:  nil,
			DebugLine: "Set:[], NoMatch:string does not match the given subexpressions",
		}
		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)

	})

	t.Run("when given an empty string return false", func(t *testing.T) {
		a := SetOfCharacters("a")
		b := SetOfCharacters("b")
		c := SetOfCharacters("c")

		exp := Set(a, b, c)
		iter := CreateIterator("")

		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:	   "Set",
			Label:     "",
			Children:  nil,
			DebugLine: "Set:[], NoMatch:string does not match the given subexpressions",
		}
		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

	t.Run("when given an empty set return false", func(t *testing.T) {
		exp := Set()
		iter := CreateIterator("abc")

		matchResult := MatchIter(&iter, exp)
		expected := MatchTree{
			IsValid:   false,
			Value:     "",
			Type:	   "Set",
			Label:     "",
			Children:  nil,
			DebugLine: "Set:[], NoMatch:number of subexpressions is zero",
		}
		require.Equal(t, expected, matchResult)
		require.Equal(t, 0, iter.index)
	})

}
