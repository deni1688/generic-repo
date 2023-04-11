package main

import (
	"testing"
)

func TestInMemRepo_Create(t *testing.T) {
	dr := newInMemRepo[Driver](make(map[int]Driver))

	id, err := dr.Create(newDriver("John"))
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if id == 0 {
		t.Errorf("id is 0")
	}

	d, err := dr.FindByID(id)
	if err != nil {
		t.Errorf("error: %v", err)

	}

	if d.Name != "John" {
		t.Errorf("name is not John")
	}
}

func TestInMemRepo_FindAll(t *testing.T) {
	dr := newInMemRepo[Driver](make(map[int]Driver))

	_, _ = dr.Create(newDriver("John"))
	_, _ = dr.Create(newDriver("Jane"))

	d, err := dr.FindAll()
	if err != nil {
		t.Errorf("error: %v", err)
	}

	if len(d) != 2 {
		t.Errorf("length is not 2")
	}
}

func TestInMemRepo_Update(t *testing.T) {
	dr := newInMemRepo[Driver](make(map[int]Driver))

	id, _ := dr.Create(newDriver("John"))

	driver := newDriver("Jane")
	driver.SetID(id)
	err := dr.Update(driver)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	d, _ := dr.FindByID(id)

	if d.Name != "Jane" {
		t.Errorf("name is not Jane")
	}
}

func TestInMemRepo_DeleteByID(t *testing.T) {
	dr := newInMemRepo[Driver](make(map[int]Driver))

	id, _ := dr.Create(newDriver("John"))

	err := dr.DeleteByID(id)
	if err != nil {
		t.Errorf("error: %v", err)
	}

	_, err = dr.FindByID(id)

	if err.Error() != "not found" {
		t.Errorf("error is nil but should be not found")
	}
}

func TestInMemRepo_DeleteByID_NotFound(t *testing.T) {
	dr := newInMemRepo[Driver](make(map[int]Driver))

	err := dr.DeleteByID(1)
	if err.Error() != "not found" {
		t.Errorf("error is nil but should be not found")
	}
}
