package functional

type Statement func() bool

func Func(statements ...Statement) Statement {
	return func() bool {
		for _, statement := range statements {
			if statement() {
				return true
			}
		}
		return false
	}
}

func Else(statement Statement) Statement {
	return statement
}

func Whilst(condition Statement, statements ...Statement) Statement {
	return func() bool {
		for condition() {
			for _, statement := range statements {
				if statement() {
					return true
				}
			}
		}
		return false
	}
}

func If(condition Statement, branches ...Statement) Statement {
	return func() bool {
		numOfBranches := len(branches)
		if condition() {
			if numOfBranches > 0 {
				return branches[0]()
			}
		} else {
			if numOfBranches > 1 {
				return branches[1]()
			}
		}
		return false
	}
}

func Not(statement Statement) Statement {
	return func() bool {
		return !statement()
	}
}

func And(left Statement, right Statement) Statement {
	return func() bool {
		return left() && right()
	}
}

func Or(left Statement, right Statement) Statement {
	return func() bool {
		return left() || right()
	}
}
