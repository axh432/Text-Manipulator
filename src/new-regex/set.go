package new_regex

func Set(expressions ...Expression) Expression {
	return func(iter *Iterator) MatchTree {
		if len(expressions) == 0 {
			return invalidMatchTree("", "Set", nil, "Set:[], NoMatch:number of subexpressions is zero")
		}

		startingIndex := iter.GetIndex()
		for _, exp := range expressions {
			match := exp(iter)
			if match.IsValid {
				return validMatchTree(match.Value, "Set", []MatchTree{match})
			}else{
				iter.Reset(startingIndex)
			}
		}
		iter.Reset(startingIndex)
		return invalidMatchTree("", "Set", nil, "Set:[], NoMatch:string does not match the given subexpressions")
	}
}
