package book

import (
	"context"
	"errors"
	"log"
	"remakech7/config"
	"remakech7/module/model"

	"gorm.io/gorm"
)

type GormRepoImpl struct {
	db *gorm.DB
}

func NewBookGormRepo(db *gorm.DB) BookRepo {
	bookRepo := GormRepoImpl{
		db: db,
	}

	if config.MIGRATE {
		bookRepo.doMigration()
	}

	return &GormRepoImpl{
		db: db,
	}
}

func (b *GormRepoImpl) doMigration() (err error) {
	if err = b.db.AutoMigrate(&model.Book{}); err != nil {
		panic(err)
	}

	log.Println("succesfully create book table")

	return
}

func (b *GormRepoImpl) FindBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	tx := b.db.
					Model(&model.Book{}).
					Where("id = ?", bookId).
					Find(&book)

	if err = tx.Error; err != nil {
		return
	}

	if bookId <= 0 {
		err = errors.New("book is not found")
	}

	return
}

func (b *GormRepoImpl) FindAllBooks(ctx context.Context) (books []model.Book, err error) {
	tx := b.db.
					Model(&model.Book{}).
					Find(&books).
					Order("id ASC")

	if err = tx.Error; err != nil {
		return
	}

	return
}

func (b *GormRepoImpl) InsertBook(ctx context.Context, bookIn model.Book) (book model.Book, err error) {
	tx := b.db.
					Model(&model.Book{}).
					Create(&bookIn)

	if err = tx.Error; err != nil {
		return
	}

	return bookIn, nil
}

func (b *GormRepoImpl) UpdateBook(ctx context.Context, bookIn model.Book) (err error) {
	tx := b.db.
					Model(&model.Book{}).
					Where("id = ?", bookIn.BookID).
					Updates(&bookIn)

	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("book is not found")
		return
	}

	return
}

func (b *GormRepoImpl) DeleteBookById(ctx context.Context, bookId uint64) (book model.Book, err error) {
	tx := b.db.
					Unscoped().
					Model(&model.Book{}).
					Where("id = ?", bookId).
					Delete(&model.Book{})

	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected <= 0 {
		err = errors.New("book is not found")
		return
	}


	return
}
