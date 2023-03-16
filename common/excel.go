package common

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
)

var fmtErr = func(err error) error { return fmt.Errorf("[Excel] wrong excel format, err = %v", err.Error()) }

type Record map[string]string

type ExcelParser interface {
	Next() (Record, error)
}

type csvParser struct {
	titles []string
	file   os.File
	reader *csv.Reader
}

func initCSVParser(file *os.File) (*csvParser, error) {
	parser := &csvParser{}
	parser.reader = csv.NewReader(file)
	var err error
	parser.titles, err = parser.reader.Read()
	for i := range parser.titles {
		parser.titles[i] = strings.ToLower(strings.TrimSpace(parser.titles[i]))
	}
	if err != nil {
		return nil, fmtErr(err)
	}
	return parser, nil
}

func (c *csvParser) Next() (Record, error) {
	records, err := c.reader.Read()
	if err == io.EOF {
		return nil, err
	}
	if err != nil || len(records) != len(c.titles) {
		return nil, fmtErr(err)
	}
	r := make(map[string]string)
	for i, v := range records {
		r[c.titles[i]] = strings.TrimSpace(v)
	}
	return r, nil
}

type xlsParser struct {
	titles []string
	index  int
	reader *xls.WorkSheet
}

func initXLSParser(file *os.File) (*xlsParser, error) {
	parser := &xlsParser{}
	workbook, err := xls.OpenReader(file, "utf-8")
	if err != nil {
		return nil, fmtErr(err)
	}
	reader := workbook.GetSheet(0)
	if reader.MaxRow == 0 {
		return nil, fmtErr(errors.New("empty"))
	}
	row := reader.Row(0)
	for i := 0; i < row.LastCol(); i++ {
		parser.titles = append(parser.titles, strings.ToLower(strings.TrimSpace(row.Col(i))))
	}
	parser.reader = reader
	parser.index = 1
	return parser, nil
}

func (c *xlsParser) Next() (Record, error) {
	if c.index >= int(c.reader.MaxRow) {
		return nil, io.EOF
	}
	r := make(map[string]string)
	row := c.reader.Row(c.index)
	for i := 0; i < len(c.titles); i++ {
		r[c.titles[i]] = strings.TrimSpace(row.Col(i))
	}
	c.index++
	return r, nil
}

type xlsxParser struct {
	titles []string
	index  int
	reader *xlsx.Sheet
}

func initXLSXParser(file *os.File) (*xlsxParser, error) {
	parser := &xlsxParser{}
	stat, _ := file.Stat()
	workbook, err := xlsx.OpenReaderAt(file, stat.Size())
	if err != nil {
		return nil, fmtErr(err)
	}
	reader := workbook.Sheets[0]
	if reader.MaxRow == 0 {
		return nil, fmtErr(errors.New("empty"))
	}
	row := reader.Row(0)
	for i := 0; i < len(row.Cells); i++ {
		parser.titles = append(parser.titles, strings.ToLower(strings.TrimSpace(row.Cells[i].Value)))
	}
	parser.reader = reader
	parser.index = 1
	return parser, nil
}

func (c *xlsxParser) Next() (Record, error) {
	if c.index >= c.reader.MaxRow {
		return nil, io.EOF
	}
	r := make(map[string]string)
	row := c.reader.Row(c.index)
	for i := 0; i < len(c.titles); i++ {
		r[c.titles[i]] = strings.TrimSpace(row.Cells[i].Value)
	}
	c.index++
	return r, nil
}

func NewExcelParser(path string) (ExcelParser, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("[Excel] no such file, path = %v", path)
	}
	if strings.HasSuffix(path, ".csv") {
		return initCSVParser(file)
	} else if strings.HasSuffix(path, ".xls") {
		return initXLSParser(file)
	} else if strings.HasSuffix(path, ".xlsx") {
		return initXLSXParser(file)
	}
	return nil, fmt.Errorf("[Excel] only suppot csv, xls and xlsx, but the file passed is %v", path)
}
