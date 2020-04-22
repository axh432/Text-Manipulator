package new_regex

func Label(exp Expression, label string) Expression {
	return func(iter *Iterator) MatchTree {
		match := exp(iter)
		return MatchTree{
			isValid:  match.isValid,
			Value:    match.Value,
			Label:    label,
			DebugLine: match.Label + ": " + match.DebugLine,
			Children: []MatchTree{match},
		}
	}
}
