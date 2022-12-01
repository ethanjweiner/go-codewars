package arraydiff

func ArrayDiff[T comparable](a []T, b []T) []T {
	diff := make([]T, 0, len(a))
	bSet := toSet(b)

	for _, item := range a {
		if !bSet[item] {
			diff = append(diff, item)
		}
	}

	return diff
}

func toSet[T comparable](slice []T) map[T]bool {
	set := map[T]bool{}

	for _, item := range slice {
		set[item] = true
	}

	return set
}
