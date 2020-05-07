package new_regex

func Label(exp Expression, label string) Expression {
	return func(iter *Iterator) MatchTree {
		match := exp(iter)
		return MatchTree{
			IsValid:   match.IsValid,
			Value:     match.Value,
			Label:     label,
			DebugLine: match.Label + ": " + match.DebugLine,
			Children:  []MatchTree{match},
		}
	}
}
