package main

import (
	"fmt"
	"math/rand"
	"time"
)

type inMemRepo[T Model] struct {
	data map[int]T
}

func newInMemRepo[T Model](data map[int]T) *inMemRepo[T] {
	return &inMemRepo[T]{data}
}

func (i *inMemRepo[T]) Create(t T) (int, error) {
	id := randomID()

	t.SetID(id)
	t.SetCreatedAt(time.Now().Format(time.RFC3339))
	t.SetUpdatedAt(time.Now().Format(time.RFC3339))

	i.data[id] = t

	return id, nil
}

func (i *inMemRepo[T]) FindByID(id int) (*T, error) {
	if t, ok := i.data[id]; ok {
		return &t, nil
	}

	return nil, fmt.Errorf("not found")
}

func (i *inMemRepo[T]) Update(t T) error {
	t.SetUpdatedAt(time.Now().Format(time.RFC3339))
	i.data[t.GetID()] = t
	return nil
}

func (i *inMemRepo[T]) DeleteByID(id int) error {
	if _, ok := i.data[id]; !ok {
		return fmt.Errorf("not found")
	}

	delete(i.data, id)
	return nil
}

func (i *inMemRepo[T]) FindAll() ([]*T, error) {
	var result []*T

	for _, v := range i.data {
		result = append(result, &v)
	}

	return result, nil
}
func randomID() int {
	seed := time.Now().UnixNano()
	rand.Seed(seed)

	min := 1_000_0000
	max := 9_999_9999

	return rand.Intn(max-min) + min
}
