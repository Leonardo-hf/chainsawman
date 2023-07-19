package common

import (
	"bytes"
	"encoding/csv"
	"errors"
	"fmt"
	"go.uber.org/multierr"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/extrame/xls"
	"github.com/tealeg/xlsx"
)

var fmtErr = func(err error) error { return fmt.Errorf("[Excel] wrong excel format, err = %v", err.Error()) }
var convErr = func(column string) error {
	return fmt.Errorf("[Excel] cannot find or parse target column, err = %v", column)
}

type strMap map[string]string
type Record struct {
	strMap
}

func (r *Record) Put(key string, value string) {
	r.strMap[key] = value
}

func (r *Record) Get(key string) (string, error) {
	if v, ok := r.strMap[key]; ok {
		return v, nil
	}
	return "", convErr(key)
}

func (r *Record) GetAsInt(key string) (int64, error) {
	if v, ok := r.strMap[key]; ok {
		vInt, err := strconv.ParseInt(v, 10, 64)
		if err == nil {
			return vInt, nil
		}
	}
	return 0, convErr(key)
}

type ExcelParser interface {
	Next() (*Record, error)
}

type csvParser struct {
	titles []string
	file   os.File
	reader *csv.Reader
}

func initCSVParser(file io.Reader) (*csvParser, error) {
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

func (c *csvParser) Next() (*Record, error) {
	records, err := c.reader.Read()
	if err == io.EOF {
		return nil, err
	}
	if err != nil || len(records) != len(c.titles) {
		return nil, fmtErr(err)
	}
	r := &Record{make(map[string]string)}
	for i, v := range records {
		r.Put(c.titles[i], strings.TrimSpace(v))
	}
	return r, nil
}

type xlsParser struct {
	titles []string
	index  int
	reader *xls.WorkSheet
}

func initXLSParser(content io.Reader) (*xlsParser, error) {
	parser := &xlsParser{}
	bs, _ := io.ReadAll(content)
	workbook, err := xls.OpenReader(bytes.NewReader(bs), "utf-8")
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

func (c *xlsParser) Next() (*Record, error) {
	if c.index >= int(c.reader.MaxRow) {
		return nil, io.EOF
	}
	r := &Record{make(map[string]string)}
	row := c.reader.Row(c.index)
	for i := 0; i < len(c.titles); i++ {
		r.Put(c.titles[i], strings.TrimSpace(row.Col(i)))
	}
	c.index++
	return r, nil
}

type xlsxParser struct {
	titles []string
	index  int
	reader *xlsx.Sheet
}

func initXLSXParser(content io.Reader) (*xlsxParser, error) {
	parser := &xlsxParser{}
	bs, _ := io.ReadAll(content)
	workbook, err := xlsx.OpenBinary(bs)
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

func (c *xlsxParser) Next() (*Record, error) {
	if c.index >= c.reader.MaxRow {
		return nil, io.EOF
	}
	r := &Record{make(map[string]string)}
	row := c.reader.Row(c.index)
	for i := 0; i < len(c.titles); i++ {
		r.Put(c.titles[i], strings.TrimSpace(row.Cells[i].Value))
	}
	c.index++
	return r, nil
}

func NewExcelParser(content io.Reader) (ExcelParser, error) {
	var merr error
	if p, err := initCSVParser(content); err != nil {
		merr = multierr.Append(merr, err)
	} else {
		return p, nil
	}
	if p, err := initXLSParser(content); err != nil {
		merr = multierr.Append(merr, err)
	} else {
		return p, nil
	}
	if p, err := initXLSXParser(content); err != nil {
		merr = multierr.Append(merr, err)
	} else {
		return p, nil
	}
	return nil, fmt.Errorf("[Excel] only suppot csv, xls and xlsx, err=%v", merr)
}
