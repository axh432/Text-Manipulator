package new_regex

import (
	"github.com/stretchr/testify/require"
	"testing"
)

/* These are the regex tutorials found at: https://regexone.com/ implemented in new-regex */

func TestRegexOne(t *testing.T) {

	t.Run("Lesson 1: the ABCs", func(t *testing.T) {
		//pattern that matches the following strings.

		exp := Range(SetOfCharacters("abcdefg"), 1, 7)

		require.True(t, Match("abcdefg", exp).IsValid)
		require.True(t, Match("abcde", exp).IsValid)
		require.True(t, Match("abc", exp).IsValid)
	})

	t.Run("Lesson 1.5: the 123s", func(t *testing.T) {
		//write a pattern that matches only the numbers.

		integer := Label(Range(Number, 1, -1), "integer")
		notANumber := Set(Whitespace, Letter, Punctuation, Symbol)
		notInteger := Range(notANumber, 1, -1)

		exp := Range(Set(notInteger, integer), 0, -1)

		numbers := []string{}
		visitor := func(mt *MatchTree) {
			if mt.Label != "" {
				numbers = append(numbers, mt.Value)
			}
		}

		result1 := Match("var g = 123;", exp)
		result2 := Match(`define "123"`, exp)
		result3 := Match(`var g = 123;`, exp)

		result1.acceptVisitor(visitor)
		result2.acceptVisitor(visitor)
		result3.acceptVisitor(visitor)

		require.Equal(t, []string{"123", "123", "123"}, numbers)
	})

	t.Run("Lesson 2: the 'any' character", func(t *testing.T) {
		//write a pattern the first three but not the last

		any := Set(Whitespace, Number, Letter, Punctuation, Symbol)
		dot := SetOfCharacters(".")
		exp := Sequence(any, any, any, dot)

		require.True(t, Match("cat.", exp).IsValid)
		require.True(t, Match("896.", exp).IsValid)
		require.True(t, Match("?=+.", exp).IsValid)
		require.False(t, Match("abc1", exp).IsValid)
	})

	t.Run("Lesson 3: Matching specific characters", func(t *testing.T) {
		//match the specific characters: cmf at the beginning of the string

		cmf := SetOfCharacters("cmf")
		an := SequenceOfCharacters("an")
		exp := Sequence(cmf, an)

		require.True(t, Match("can", exp).IsValid)
		require.True(t, Match("man", exp).IsValid)
		require.True(t, Match("fan", exp).IsValid)
		require.False(t, Match("dan", exp).IsValid)
		require.False(t, Match("ran", exp).IsValid)
		require.False(t, Match("pan", exp).IsValid)

	})

	//lesson 4 the 'not' expression has not been implemented yet
	t.Run("Lesson 4: Excluding specific characters", func(t *testing.T) {
		/*
			Match	can
			Match	man
			Match	fan
			Skip	dan
			Skip	ran
			Skip	pan
		*/
	})

	//lesson 5 character ranges is not supported atm
	t.Run("Lesson 5: Character ranges", func(t *testing.T) {
		/*
		Match	Ana
		Match	Bob
		Match	Cpc
		Skip	aax
		Skip	bby
		Skip	ccz
		*/
	})

	t.Run("Lesson 6: Catching some zzz's", func(t *testing.T) {
		//use ranges to match the strings that need to be matched.
		wa := SequenceOfCharacters("wa")
		z := SetOfCharacters("z")
		up := SequenceOfCharacters("up")
		exp := Sequence(wa, Range(z, 3, 5), up)

		require.True(t, Match("wazzzzzup", exp).IsValid)
		require.True(t, Match("wazzzup", exp).IsValid)
		require.False(t, Match("wazup", exp).IsValid)
	})

	t.Run("Lesson 7: Matching Repeated Characters", func(t *testing.T) {
		a := Range(SetOfCharacters("a"), 2, 4)
		b := Range(SetOfCharacters("b"), 0, 4)
		c := Range(SetOfCharacters("c"), 1, 2)
		exp := Sequence(a, b, c)

		require.True(t, Match("aaaabcc", exp).IsValid)
		require.True(t, Match("aabbbbc", exp).IsValid)
		require.True(t, Match("aacc", exp).IsValid)
		require.False(t, Match("a", exp).IsValid)
	})

	t.Run("Lesson 8: Characters optional", func(t *testing.T) {
		integer := Range(Number, 1, -1)
		space := SetOfCharacters(" ")
		file := SequenceOfCharacters("file")
		files := SequenceOfCharacters("files")
		found := SequenceOfCharacters("found?")
		exp := Sequence(integer, space, Set(file, files), space, found)

		require.True(t, Match("1 file found?", exp).IsValid)
		require.True(t, Match("2 files found?", exp).IsValid)
		require.True(t, Match("24 files found?", exp).IsValid)
		require.False(t, Match("No files found.", exp).IsValid)
		/*
		Match	1 file found?
		Match	2 files found?
		Match	24 files found?
		Skip	No files found.
		*/
	})

	t.Run("Lesson 9: All this whitespace", func(t *testing.T) {
		/*
			Match	1.   abc
			Match	2.	abc
			Match	3.           abc
			Skip	4.abc
		*/
	})

	t.Run("Lesson 10: Starting and ending", func(t *testing.T) {
		/*
			Match	Mission: successful
			Skip	Last Mission: unsuccessful
			Skip	Next Mission: successful upon capture of target
		*/
	})

	t.Run("Lesson 11: Match groups", func(t *testing.T) {
		//capture only the file name and not the extension
		/*
			Capture	file_record_transcript.pdf	file_record_transcript
			Capture	file_07241999.pdf	file_07241999
			Skip	testfile_fake.pdf.tmp
		*/
	})

	t.Run("Lesson 12: Nested groups", func(t *testing.T) {
		//capture the full date and the year of the date
		/*
		Capture	Jan 1987	Jan 1987 1987
		Capture	May 1969	May 1969 1969
		Capture	Aug 2011	Aug 2011 2011
		*/
	})

	t.Run("Lesson 13: More group work", func(t *testing.T) {
		//capture the individual dimensions
		/*
		Capture	1280x720	1280 720
		Capture	1920x1600	1920 1600
		Capture	1024x768	1024 768
		*/
	})

	t.Run("Lesson 14: It's all conditional", func(t *testing.T) {
		/*
			Match	I love cats
			Match	I love dogs
			Skip	I love logs
			Skip	I love cogs
		*/
	})

	t.Run("Lesson 15: Other special characters", func(t *testing.T) {
		/*
			Match	The quick brown fox jumps over the lazy dog.
			Match	There were 614 instances of students getting 90.0% or above.
			Match	The FCC had to censor the network for saying &$#*@!.
		*/
	})

}