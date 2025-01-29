package repositories

import "gorm.io/gorm"

type RepositoryResult[T any] struct {
	Result    T
	RowsCount int
	NoRecord  bool
	RowsError error
}

type Repositories[T any] interface {
	FindAllPaginate(offset int, limit int, filter string, id_account int) RepositoryResult[any]
	FindAll() RepositoryResult[any]
	Find() RepositoryResult[any]
	BulkCreate() RepositoryResult[any]
	Create() RepositoryResult[any]
	Update() RepositoryResult[any]
	BulkUpdate() RepositoryResult[any]
	CustomQuery() RepositoryResult[any]
	Delete()
}
type Repository[T any] struct {
	Wrapper T
}

func RepositoryTransaction[T any](rows *gorm.DB, wrapper T) RepositoryResult[T] {
	return RepositoryResult[T]{
		Result:    wrapper,
		RowsCount: int(rows.RowsAffected),
		NoRecord:  bool(int(rows.RowsAffected) == 0),
		RowsError: rows.Error,
	}
}

func (repo *Repository[T]) FindAllPaginate(offset int, limit int, filter string, id_account int) RepositoryResult[T] {
	return RepositoryTransaction(
		db.Limit(limit).Offset(offset).Find(&repo.Wrapper),
		repo.Wrapper,
	)
}

func (repo *Repository[T]) FindAll() RepositoryResult[T] {
	return RepositoryTransaction(
		db.Find(&repo.Wrapper),
		repo.Wrapper,
	)
}

func (repo *Repository[T]) Find() RepositoryResult[T] {
	return RepositoryTransaction(
		db.Find(&repo.Wrapper),
		repo.Wrapper,
	)
}

func (repo *Repository[T]) Create() RepositoryResult[T] {
	return RepositoryTransaction(
		db.Create(&repo.Wrapper),
		repo.Wrapper,
	)
}
