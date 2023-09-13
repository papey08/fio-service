package valid

const allowedNameSymbols = " abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-'"

func isAllowed(c rune) bool {
	for _, s := range allowedNameSymbols {
		if s == c {
			return true
		}
	}
	return false
}

func Name(name string) bool {
	for _, c := range name {
		if !isAllowed(c) {
			return false
		}
	}
	return true
}
