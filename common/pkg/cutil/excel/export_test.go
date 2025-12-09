package excel

import (
	"database/sql"
	"fmt"
	"strings"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestExport(t *testing.T) {

	f := NewFile("test.xlsx")
	f.SetCheckFunc(func(key string, value any) error {
		if key == "age" && value != nil {
			if v, ok := value.(int); ok && v > 22 {
				return fmt.Errorf("age is greater than 22")
			}
			if v, ok := value.(int64); ok && v > 22 {
				return fmt.Errorf("age is greater than 22")
			}
		}
		return nil
	})
	title := []string{"id", "姓名", "年龄"}
	keys := []string{"id", "name", "age"}
	data := []map[string]any{
		{"id": 1, "name": "张三", "age": 18},
		{"id": 2, "name": "李四", "age": 19},
		{"id": 3, "name": "王五", "age": 20},
		{"id": 4, "name": "赵六", "age": 21},
		{"id": 5, "name": "孙七", "age": 22},
		{"id": 6, "name": "周八", "age": 23},
		{"id": 7, "name": "吴九", "age": 24},
		{"id": 8, "name": "郑十", "age": 25},
	}

	db := InitDB(t)

	t.Run("ExportFromDataMap", func(t *testing.T) {
		err := f.ExportFromDataMap(t.Name()[strings.LastIndex(t.Name(), "/")+1:], title, keys, data)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("ExportFromQuery", func(t *testing.T) {
		err := f.ExportFromQuery(t.Name()[strings.LastIndex(t.Name(), "/")+1:], title, keys, db, "select * from user")
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("ExportStreamFromDataMap", func(t *testing.T) {
		err := f.ExportStreamFromDataMap(t.Name()[strings.LastIndex(t.Name(), "/")+1:], title, keys, data)
		if err != nil {
			t.Error(err)
		}
	})
	t.Run("ExportStreamFromQuery", func(t *testing.T) {
		err := f.ExportStreamFromQuery(t.Name()[strings.LastIndex(t.Name(), "/")+1:], title, keys, db, "select * from user")
		if err != nil {
			t.Error(err)
		}
	})

	err := f.WriteToFile(".")
	if err != nil {
		t.Error(err)
	}
	defer func(db *sql.DB) {
		_ = db.Close()
	}(db)
}

func InitDB(t *testing.T) *sql.DB {
	if testing.Short() {
		t.Skip("skip db tests in short mode")
	}
	db, err := sql.Open("mysql", "root:mysql@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		t.Fatal(err)
	}
	return db
}
