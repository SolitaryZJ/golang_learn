package topic3

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

const schema = `
CREATE TABLE IF NOT EXISTS Employees (
    id        BIGINT PRIMARY KEY AUTO_INCREMENT,
    name      VARCHAR(50)  NOT NULL,
    department      VARCHAR(50)  NOT NULL,
    salary       INT          DEFAULT 0
);`

type Employee struct {
	ID         int
	Name       string
	Department string
	Salary     float32
}

func Run(db *sqlx.DB) {
	db.MustExec(schema)
	db.MustExec("INSERT INTO Employees (name, department, salary) VALUES (?, ?, ?)", "Jane", "技术部", 5000.0)
	var employee Employee
	// 查询 employees 表中所有部门为 "技术部" 的员工信息，并将结果映射到一个自定义的 Employee 结构体切片中
	db.Get(&employee, "SELECT * FROM Employees WHERE id = ?", 1)
	fmt.Println(employee)
	// 查询 employees 表中工资最高的员工信息，并将结果映射到一个 Employee 结构体中
	db.Get(&employee, "SELECT * FROM Employees ORDER BY salary DESC LIMIT 1")
	fmt.Println(employee)
}
