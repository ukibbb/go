package main

type Repository[T Entity] struct {
	db DataStore[T]
}

func (r *Repository[Entity]) Get() {}
func (r *Repository[Entity]) Create(e *Entity) (*Entity, error) {
	ce, err := r.db.Create(e)
	if err != nil {
		return nil, err
	}
	return ce, nil
}
func (r *Repository[Entity]) Update() {}
func (r *Repository[Entity]) Delete() {}

func NewRepository[T Entity](db DataStore[T]) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}
