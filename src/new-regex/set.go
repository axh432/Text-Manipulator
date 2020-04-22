package new_regex

func Set(expressions ...Expression2) Expression2 {
	return func(iter *Iterator) MatchTree {
		startingIndex := iter.GetIndex()
		for _, exp := range expressions {
			match := exp(iter)
			if match.isValid {
				return MatchTree{
					isValid:  true,
					Value:    match.Value,
					Name:     "",
					Children: []MatchTree{match},
				}
			} else {
				iter.Reset(startingIndex)
			}
		}
		return MatchTree{}
	}
}
