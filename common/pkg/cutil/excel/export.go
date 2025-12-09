package excel

import (
	"common/pkg/cutil/base/logger"
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"github.com/xuri/excelize/v2"
)

// CheckFunc 校验函数
//
// 参数:
//   - value: 写入单元格的值
//   - key: 写入单元格对应的键
//
// 返回:
//   - error: 校验失败的错误信息
type CheckFunc func(key string, value any) error

// File excel文件对象
type File struct {
	fileName     string
	file         *excelize.File
	checkFunc    CheckFunc
	errorStyleId *int
	titleStyleId *int
}

func NewFile(fileName string) *File {
	f := &File{fileName: fileName, file: excelize.NewFile()}
	// 添加默认样式
	if f.errorStyleId == nil {
		_ = f.SetErrorStyle(nil)
	}
	return f
}

func (f *File) SetCheckFunc(checkFunc CheckFunc) {
	f.checkFunc = checkFunc
}

func (f *File) SetErrorStyle(style *excelize.Style) error {
	if style == nil {
		styleId, err := f.file.NewStyle(&excelize.Style{
			Fill: excelize.Fill{
				Type:    "pattern",
				Color:   []string{"#FFFF00"}, // 黄色背景
				Pattern: 1,                   // 实心填充
			},
		})
		if err != nil {
			return fmt.Errorf("failed to create style: %w", err)
		}
		f.errorStyleId = &styleId
	} else {
		styleId, err := f.file.NewStyle(style)
		if err != nil {
			return fmt.Errorf("failed to create style: %w", err)
		}
		f.errorStyleId = &styleId
	}
	return nil
}

func (f *File) SetTitleStyle(style *excelize.Style) error {
	if style == nil {
		return fmt.Errorf("title style cannot be nil")
	}
	styleId, err := f.file.NewStyle(style)
	if err != nil {
		return fmt.Errorf("create title style: %w", err)
	}
	f.titleStyleId = &styleId
	return nil
}

func (f *File) GetExcelizeFile() *excelize.File {
	return f.file
}

func (f *File) checkParams(sheetName string, titles []string, keys []string) error {
	if len(titles) != len(keys) {
		return fmt.Errorf("titles and keys length must be equal")
	}
	sheetIndex, err := f.file.GetSheetIndex(sheetName)
	if err != nil {
		return fmt.Errorf("get sheet index: %w", err)
	}
	if sheetIndex == -1 {
		sheet, err := f.file.NewSheet(sheetName)
		if err != nil {
			return fmt.Errorf("create sheet: %w", err)
		}
		f.file.SetActiveSheet(sheet)
	}
	return nil
}

// WriteToFile 将excel数据写入指定位置文件
//
// 参数:
//   - path: 文件路径
//
// 返回:
//   - error: 写入失败的错误信息
func (f *File) WriteToFile(path string) error {
	if f.fileName == "" {
		return fmt.Errorf("file name is empty")
	}
	file, err := os.Create(filepath.Join(path, f.fileName))
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		if err := file.Close(); err != nil {
			logger.Error("close file error: %w", err)
		}
	}(file)

	if err := f.file.Write(file); err != nil {
		return err
	}
	return nil
}

// ================= 公共逻辑 =================

// writeTitles 写标题行（支持流式/非流式）
//
// 参数:
//   - sheetName: 工作表名称
//   - titles: 标题行值
//   - sw: 流式写入器
//
// 返回:
//   - error: 写入失败的错误信息
func (f *File) writeTitles(sheetName string, titles []string, sw *excelize.StreamWriter) error {
	rowIndex := 1
	titleData := make([]any, len(titles))

	for i, title := range titles {
		cell := excelize.Cell{Value: title}
		if f.titleStyleId != nil {
			cell.StyleID = *f.titleStyleId
		}
		titleData[i] = cell

		if sw == nil {
			cellAddr, _ := excelize.CoordinatesToCellName(i+1, rowIndex)
			if err := f.file.SetCellValue(sheetName, cellAddr, title); err != nil {
				return fmt.Errorf("set title cell value failed: %w", err)
			}
			if cell.StyleID > 0 {
				if err := f.file.SetCellStyle(sheetName, cellAddr, cellAddr, cell.StyleID); err != nil {
					return fmt.Errorf("set title cell style failed: %w", err)
				}
			}
		}
	}
	cellIndex, _ := excelize.CoordinatesToCellName(1, rowIndex)
	if sw != nil {
		return sw.SetRow(cellIndex, titleData)
	}
	return nil
}

// writeRow 写一行数据（支持流式/非流式）
//
// 参数:
//   - sheetName: 工作表名称
//   - rowIndex: 行索引（从 1 开始）
//   - keys: 数据键
//   - row: 数据值
//   - sw: 流式写入器
//
// 返回:
//   - error: 写入失败的错误信息
func (f *File) writeRow(sheetName string, rowIndex int, keys []string, row map[string]any, sw *excelize.StreamWriter) error {
	data := make([]any, len(keys))

	for j, key := range keys {
		cellAddr, _ := excelize.CoordinatesToCellName(j+1, rowIndex)
		cell := excelize.Cell{Value: row[key]}

		if f.checkFunc != nil && f.errorStyleId != nil {
			if err := f.checkFunc(key, row[key]); err != nil {
				// 添加样式
				cell.StyleID = *f.errorStyleId
				// 添加批注
				if err := f.file.AddComment(sheetName, excelize.Comment{
					Cell:   cellAddr,
					Author: "System",
					Text:   fmt.Sprintf("%v", err.Error()),
				}); err != nil {
					return fmt.Errorf("failed to add comment: %w", err)
				}
			}
		}

		data[j] = cell

		if sw == nil {
			if err := f.file.SetCellValue(sheetName, cellAddr, cell.Value); err != nil {
				return fmt.Errorf("failed to set cell value: %w", err)
			}
			if cell.StyleID > 0 {
				if err := f.file.SetCellStyle(sheetName, cellAddr, cellAddr, cell.StyleID); err != nil {
					return fmt.Errorf("set cell style failed: %w", err)
				}
			}
		}
	}
	cellIndex, _ := excelize.CoordinatesToCellName(1, rowIndex)
	if sw != nil {
		return sw.SetRow(cellIndex, data)
	}
	return nil
}

// WriteCell 将值写入 Excel 单元格，并对写入值进行校验
//
// 参数:
//   - sheetName: 工作表名称
//   - rowIndex: 行索引（从 1 开始）
//   - colIndex: 列索引（从 1 开始）
//   - value: 要写入的值
//
// 返回:
//   - error: 写入失败的错误信息
func (f *File) WriteCell(sheetName string, rowIndex int, colIndex int, key string, value any) error {
	cellAddr, err := excelize.CoordinatesToCellName(colIndex, rowIndex)
	if err != nil {
		return fmt.Errorf("conversion coordinates (%d, %d) to cell name failed: %w", colIndex, rowIndex, err)
	}
	if f.checkFunc != nil && f.errorStyleId != nil {
		if checkErr := f.checkFunc(key, value); checkErr != nil {
			if err := f.file.SetCellStyle(sheetName, cellAddr, cellAddr, *f.errorStyleId); err != nil {
				return fmt.Errorf("failed to set cell style: %w", err)
			}
			// 添加批注
			if err := f.file.AddComment(sheetName, excelize.Comment{
				Cell:   cellAddr,
				Author: "System",
				Text:   fmt.Sprintf("%v", checkErr.Error()),
			}); err != nil {
				return fmt.Errorf("failed to add comment: %w", err)
			}
		}
	}
	if err := f.file.SetCellValue(sheetName, cellAddr, value); err != nil {
		return fmt.Errorf("failed to set cell %s: %w", cellAddr, err)
	}
	return nil
}

func (f *File) iterateRows(sheetName string, keys []string, db *sql.DB, query string, args []any, sw *excelize.StreamWriter) error {
	rows, err := db.Query(query, args...)
	if err != nil {
		return err
	}
	defer func(rows *sql.Rows) {
		if err := rows.Close(); err != nil {
			logger.Error("rows.Close() error: %w", err)
		}
	}(rows)
	columns, err := rows.Columns()
	if err != nil {
		return err
	}
	columnMap := make(map[string]int)
	for i, col := range columns {
		columnMap[col] = i
	}
	for _, key := range keys {
		if _, exists := columnMap[key]; !exists {
			logger.Warn("column %s not exists", key)
		}
	}
	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	rowIndex := 1
	for rows.Next() {
		rowIndex++
		err = rows.Scan(valuePtrs...)
		if err != nil {
			return err
		}
		rowData := make(map[string]any, len(keys))
		for _, key := range keys {
			if idx, exists := columnMap[key]; exists {
				rowData[key] = values[idx]
			}
		}
		if err := f.writeRow(sheetName, rowIndex, keys, rowData, sw); err != nil {
			return fmt.Errorf("write row %d err: %w", rowIndex, err)
		}
	}
	return nil
}

// ================= 导出函数 =================

// ExportFromDataMap 从数据集合导出数据
//
// 参数:
//   - sheetName: 工作表名称
//   - titles: 标题行
//   - keys: 数据行对应的键
//   - values: 数据键值map集合数组
//
// 返回:
//   - error: 写入失败的错误信息
func (f *File) ExportFromDataMap(sheetName string, titles []string, keys []string, values []map[string]any) error {
	if err := f.checkParams(sheetName, titles, keys); err != nil {
		return err
	}

	err := f.writeTitles(sheetName, titles, nil)
	if err != nil {
		return err
	}
	for i, value := range values {
		err = f.writeRow(sheetName, i+2, keys, value, nil)
		if err != nil {
			return err
		}
	}
	return nil
}

// ExportFromQuery 从数据库查询数据并导出
//
// 参数:
//   - sheetName: 工作表名称
//   - titles: 标题行
//   - keys: 数据行对应的键（列名）
//   - db: 数据源
//   - query: 查询sql
//   - args: 查询参数
//
// 返回:
//   - error: 写入失败的错误信息
func (f *File) ExportFromQuery(sheetName string, titles []string, keys []string, db *sql.DB, query string, args ...any) error {
	if err := f.checkParams(sheetName, titles, keys); err != nil {
		return err
	}

	err := f.writeTitles(sheetName, titles, nil)
	if err != nil {
		return err
	}

	return f.iterateRows(sheetName, keys, db, query, args, nil)
}

// ExportStreamFromDataMap 从数据源导出数据（流式，适合大规模数据）
//
// 参数:
//   - sheetName: 工作表名称
//   - titles: 标题行
//   - keys: 数据行对应的键（列名）
//   - values: 数据键值map集合数组
//
// 返回:
//   - error: 写入失败的错误信息
func (f *File) ExportStreamFromDataMap(sheetName string, titles []string, keys []string, values []map[string]any) error {
	if err := f.checkParams(sheetName, titles, keys); err != nil {
		return err
	}
	sw, err := f.file.NewStreamWriter(sheetName)
	if err != nil {
		return err
	}

	err = f.writeTitles(sheetName, titles, sw)
	if err != nil {
		return err
	}
	for i, value := range values {
		err = f.writeRow(sheetName, i+2, keys, value, sw)
		if err != nil {
			return err
		}
	}

	// 结束流式写入
	if err = sw.Flush(); err != nil {
		return err
	}
	return nil
}

// ExportStreamFromQuery 从数据库查询数据并导出（流式，适合大规模数据）
//
// 参数:
//   - sheetName: 工作表名称
//   - titles: 标题行
//   - keys: 数据行对应的键（列名）
//   - db: 数据源
//   - query: 查询sql
//   - args: 查询参数
//
// 返回:
//   - error: 写入失败的错误信息
func (f *File) ExportStreamFromQuery(sheetName string, titles []string, keys []string, db *sql.DB, query string, args ...any) error {
	if err := f.checkParams(sheetName, titles, keys); err != nil {
		return err
	}
	sw, err := f.file.NewStreamWriter(sheetName)
	if err != nil {
		return err
	}

	if err = f.writeTitles(sheetName, titles, sw); err != nil {
		return err
	}

	err = f.iterateRows(sheetName, keys, db, query, args, sw)
	if err != nil {
		return err
	}

	// 结束流式写入
	if err = sw.Flush(); err != nil {
		return err
	}
	return nil
}
