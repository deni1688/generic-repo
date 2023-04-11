package main

type Repo[T Model] interface {
	Create(entity T) (int, error)
	Update(entity T) error
	FindByID(id int) (*T, error)
	FindAll() ([]*T, error)
	DeleteByID(id int) error
}

type Model interface {
	Driver | Vehicle
	GetID() int
	SetID(id int)
	SetCreatedAt(createdAt string)
	SetUpdatedAt(updatedAt string)
}
type Base struct {
	ID        int    `json:"id"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

func (b *Base) GetID() int {
	return b.ID
}

func (b *Base) SetID(id int) {
	b.ID = id
}

func (b *Base) SetCreatedAt(createdAt string) {
	b.CreatedAt = createdAt
}

func (b *Base) SetUpdatedAt(updatedAt string) {
	b.UpdatedAt = updatedAt
}
