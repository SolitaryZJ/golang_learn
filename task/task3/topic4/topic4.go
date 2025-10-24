package topic4

import (
	"github.com/jmoiron/sqlx"
)

const schema = `
CREATE TABLE IF NOT EXISTS books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(255),
  author VARCHAR(255),
  price FLOAT
);
`

type Book struct {
	ID     int
	Title  string
	Author string
	Price  float32
}

func Run(db *sqlx.DB) {
	// 查询价格大于 50 元的书籍，并将结果映射到 Book 结构体切片中
	var books []Book

	db.MustExec(schema)
	db.MustExec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)", "《Go 语言》", " effective-go", 50.0)
	db.MustExec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)", "《python 语言》", " effective-go", 30.0)
	db.MustExec("INSERT INTO books (title, author, price) VALUES (?, ?, ?)", "《java 语言》", " effective-go", 80.0)

	db.Select(&books, "SELECT * FROM books WHERE price > ?", 50.0)
	for _, book := range books {
		println(book.Title, book.Author, book.Price)
	}
}
