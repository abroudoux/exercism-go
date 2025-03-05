package strain

func Keep[T any](collection []T, prediction func(T) bool) []T {
	if len(collection) == 0 {
		return collection
	}

	var collectionToReturn []T
	for _, v := range collection {
		if prediction(v) {
			collectionToReturn = append(collectionToReturn, v)
		}
	}

	return collectionToReturn
}

func Discard[T any](collection []T, prediction func(T) bool) []T {
	if len(collection) == 0 {
		return collection
	}

	var collectionToReturn []T
	for _, v := range collection {
		if !prediction(v) {
			collectionToReturn = append(collectionToReturn, v)
		}
	}

	return collectionToReturn
}
