package new_regex

type EqualityOperator func(left, right int) bool

func LessThan(left, right int) bool           { return left < right }
func LessThanOrEqual(left, right int) bool    { return left <= right }
func GreaterThan(left, right int) bool        { return left > right }
func GreaterThanOrEqual(left, right int) bool { return left >= right }
func Equal(left, right int) bool              { return left == right }
