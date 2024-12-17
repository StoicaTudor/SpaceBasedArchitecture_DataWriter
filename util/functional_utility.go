package util

import "fmt"

type Consumer[T any] interface {
	Consume(T)
}

func Cast[TypeToCastFrom, TypeToCastTo any](objectToCastFrom TypeToCastFrom) (TypeToCastTo, error) {
	y, ok := any(objectToCastFrom).(TypeToCastTo)
	if !ok {
		return *new(TypeToCastTo), fmt.Errorf("failed to cast %T to %T", objectToCastFrom, *new(TypeToCastTo))
	}
	return y, nil
}
