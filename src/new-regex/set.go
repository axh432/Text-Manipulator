package new_regex

func Set(expressions ...Expression) Expression {
	return func(iter *Iterator) MatchTree {
		startingIndex := iter.GetIndex()
		mt := MatchTree{}
		mt.Label = "Set"
		mt.DebugLine = "expecting either: "
		for _, exp := range expressions {
			match := exp(iter)
			if match.isValid {
				mt.DebugLine = ""
				mt.isValid = true
				mt.Value = match.Value
				mt.Children = append(mt.Children, match)
			} else {
				mt.DebugLine += match.Label + " or "
				iter.Reset(startingIndex)
			}
		}
		return mt
	}
}
