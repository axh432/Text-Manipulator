package new_regex

func Set(expressions ...Expression) Expression {
	return func(iter *Iterator) MatchTree {
		startingIndex := iter.GetIndex()
		for _, exp := range expressions {
			match := exp(iter)
			if match.isValid {
				return MatchTree{
					isValid:  true,
					Value:    match.Value,
					Label:    "",
					Children: []MatchTree{match},
				}
			} else {
				iter.Reset(startingIndex)
			}
		}
		return MatchTree{}
	}
}
