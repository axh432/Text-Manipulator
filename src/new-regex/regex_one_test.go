package new_regex

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

/* These are the regex tutorials found at: https://regexone.com/ implemented in new-regex */

func TestRegexOne(t *testing.T) {

	t.Run("Lesson 1: the ABCs", func(t *testing.T) {
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

		fmt.Println(result2.toMermaidDiagram())

		result1.acceptVisitor(visitor)
		result2.acceptVisitor(visitor)
		result3.acceptVisitor(visitor)

		require.Equal(t, []string{"123", "123", "123"}, numbers)
	})

	t.Run("Lesson 2: the 'any' character", func(t *testing.T) {
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

	//Todo: Do we need a 'not' expression? When would you need it?
	//Todo: the main problem with it is that I dont know if the iterator is supposed to move on or not.
	//Todo: maybe 'not' can only exist for sets of characters?
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

	//lesson 5 character ranges are not supported atm
	//Todo: do we need a special type of SetOfCharacters e.g. RangedSetOfCharacters('A', 'm'), RangedSetOfCharacters('0', '6')?
	//Todo: in regex our 'Ranges' are called 'Repetitions' maybe we should rename them?
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
	})

	t.Run("Lesson 9: All this whitespace", func(t *testing.T) {
		exp := Sequence(Range(Whitespace, 1, -1), SequenceOfCharacters("abc"))

		require.True(t, Match("   abc", exp).IsValid)
		require.True(t, Match("\tabc", exp).IsValid)
		require.True(t, Match("           abc", exp).IsValid)
		require.False(t, Match("abc", exp).IsValid)
	})

	//Todo: do we need the start or end of string markers?
	t.Run("Lesson 10: Starting and ending", func(t *testing.T) {
		exp := SequenceOfCharacters("Mission: successful")

		require.True(t, Match("Mission: successful", exp).IsValid)
		require.False(t, Match("Last Mission: unsuccessful", exp).IsValid)
		require.False(t, Match("Next Mission: successful upon capture of target", exp).IsValid)
	})

	t.Run("Lesson 11: Match groups", func(t *testing.T) {
		//capture only the file name and not the extension
		filename := Label(Range(Set(Letter, Number, SetOfCharacters("_")), 1, -1), "filename")
		fileExtension := SequenceOfCharacters(".pdf")
		nonSpace := Set(Letter, Number, Punctuation, Symbol) //you can build not expressions
		noCharactersAfter := Range(nonSpace, 0, 0) //custom string end expression
		exp := Sequence(filename, fileExtension, noCharactersAfter)

		filenames := []string{}
		visitor := func(mt *MatchTree) {
			if mt.Label == "filename" {
				filenames = append(filenames, mt.Value)
			}
		}

		result1 := Match("file_record_transcript.pdf", exp)
		result2 := Match("file_07241999.pdf", exp)
		result3 := Match("testfile_fake.pdf.tmp", exp)

		require.True(t, result1.IsValid)
		require.True(t, result2.IsValid)
		require.False(t, result3.IsValid)

		result1.acceptVisitor(visitor)
		result2.acceptVisitor(visitor)
		require.Equal(t, []string{"file_record_transcript", "file_07241999"}, filenames)
	})

	t.Run("Lesson 12: Nested groups", func(t *testing.T) {
		//capture the full date and the year of the date
		year := Label(Range(Number, 4, 4), "year")
		space := SequenceOfCharacters(" ")
		month := Label(Set(
			SequenceOfCharacters("Jan"),
			SequenceOfCharacters("May"),
			SequenceOfCharacters("Aug")), "month")
		date := Label(Sequence(month, space, year), "date")

		dates := []string{}
		years := []string{}
		visitor := func(mt *MatchTree) {
			if mt.Label == "date" {
				dates = append(dates, mt.Value)
			}
			if mt.Label == "year" {
				years = append(years, mt.Value)
			}
		}

		result1 := Match("Jan 1987", date)
		result2 := Match("May 1969", date)
		result3 := Match("Aug 2011", date)

		result1.acceptVisitor(visitor)
		result2.acceptVisitor(visitor)
		result3.acceptVisitor(visitor)

		require.True(t, result1.IsValid)
		require.True(t, result2.IsValid)
		require.True(t, result3.IsValid)
		require.Equal(t, []string{"Jan 1987", "May 1969", "Aug 2011"}, dates)
		require.Equal(t, []string{"1987", "1969", "2011"}, years)

	})

	t.Run("Lesson 13: More group work", func(t *testing.T) {
		//capture the individual dimensions
		dimension := Label(Range(Number, 1, 4), "dimension")
		x := SequenceOfCharacters("x")
		size := Sequence(dimension, x, dimension)

		dimensions := []string{}
		visitor := func(mt *MatchTree) {
			if mt.Label == "dimension" {
				dimensions = append(dimensions, mt.Value)
			}
		}

		result1 := Match("1280x720", size)
		result2 := Match("1920x1600", size)
		result3 := Match("1024x768", size)

		result1.acceptVisitor(visitor)
		result2.acceptVisitor(visitor)
		result3.acceptVisitor(visitor)

		require.True(t, result1.IsValid)
		require.True(t, result2.IsValid)
		require.True(t, result3.IsValid)
		require.Equal(t, []string{"1280", "720", "1920", "1600", "1024", "768"}, dimensions)
	})

	t.Run("Lesson 14: It's all conditional", func(t *testing.T) {
		cats := SequenceOfCharacters("cats")
		dogs := SequenceOfCharacters("dogs")
		catsOrDogs := Set(cats, dogs)
		iLove := SequenceOfCharacters("I love ")
		exp := Sequence(iLove, catsOrDogs)

		require.True(t, Match("I love cats", exp).IsValid)
		require.True(t, Match("I love dogs", exp).IsValid)
		require.False(t, Match("I love logs", exp).IsValid)
		require.False(t, Match("I love cogs", exp).IsValid)
	})

	//Todo: it is possible to go into an infinite loop if the Range of the word is set to min=0
	//Todo: every match if successful needs to move the iterator on in some way. it is not valid to a have 'non' character.
	//Todo: some sort of logging 
	t.Run("Lesson 15: Other special characters", func(t *testing.T) {

		word := Label(Range(Letter, 0, -1), "word")
		space := SetOfCharacters(" ")
		integer := Label(Range(Number, 1, -1), "integer")
		decimal := Sequence(integer, SetOfCharacters("."), integer)
		percentage := Label(Sequence(decimal, SetOfCharacters("%")), "percentage")
		fullStop := Label(SequenceOfCharacters("."), "fullstop")
		expletive := Label(SequenceOfCharacters("&$#*@!"), "expletive")
		sentence := Sequence(Range(Set(word, space, integer, percentage, expletive), 1, -1), fullStop)

		fmt.Println(Match("There were 614 instances of students getting 90.0% or above.", sentence).toMermaidDiagram())

		/*require.True(t, Match("The quick brown fox jumps over the lazy dog.", sentence).IsValid)
		require.True(t, Match("There were 614 instances of students getting 90.0% or above.", sentence).IsValid)
		require.True(t, Match("The FCC had to censor the network for saying &$#*@!.", sentence).IsValid)*/
	})

}
