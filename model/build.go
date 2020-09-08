package model

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type BuildTable struct {
	db     string
	table  string
	s      []string
	w      string
	pKey   string
	key    []string
	path   string
	fPath  string
	tableT int
}

const (
	tablePrefix = "gugu_"
)

//加载文件
func Load() error {
	table := Parse()

	file, err := os.Open(table.fPath)
	defer file.Close()

	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		table.analyze(filterString(scanner.Text()))
	}
	table.buildFile()
	table.buildCommon()

	return nil
}

func Parse() BuildTable {
	name := flag.String("name", "", "生成文件的服务名称")
	path := flag.String("path", "", "SQL文件地址")
	tableT := flag.Int("type", 1, "表格类型")

	flag.Parse()

	table := BuildTable{}
	table.setPath(*name)
	table.setFilePath(*path)
	table.setTableT(*tableT)
	return table
}

//解析文件
func (b *BuildTable) analyze(s string) {
	if b.db == "" {
		b.setBase(s)
		return
	}
	b.setSql(s)
}

//解析包名
func (b *BuildTable) setBase(s string) {
	b.db = strings.Replace(s, "use", "", 1)
}

func (b *BuildTable) setPath(s string) {
	b.path = s
}

func (b *BuildTable) setFilePath(s string) {
	b.fPath = s
}

func (b *BuildTable) setTableT(t int) {
	b.tableT = t
}

//解析SQL
func (b *BuildTable) setSql(s string) {
	if s == "" {
		return
	}
	if b.table == "" {
		sList := strings.Split(s, "`")
		for _, v := range sList {
			if strings.Contains(v, tablePrefix) {
				b.table = strings.Replace(v, tablePrefix, "", 1)
				return
			}
		}
	}
	if b.pKey == "" && strings.Contains(s, "primary") {
		b.pKey = "id"
		return
	}

	if b.pKey != "" && strings.Contains(s, "key") {
		b.key = append(b.key, s)
	}

	if !strings.Contains(s, "innodb") {
		b.s = append(b.s, s)
	}
}

func (b *BuildTable) buildFile() error {
	fileName := fmt.Sprintf("../../app/%s/model/%s/%s.go", b.path, strings.TrimSpace(b.db), b.table)
	f, err := os.Create(fileName)

	defer func() {
		f.Close()
		cmd := exec.Command("go", "fmt", fileName)
		cmd.Run()
	}()

	if err != nil {
		return err
	}

	f.WriteString(GenContent(b.table, b.db, b.s))
	return nil
}

func (b *BuildTable) buildCommon() error {
	fileName := fmt.Sprintf("../../app/%s/model/%s/common.go", b.path, strings.TrimSpace(b.db))
	f, err := os.Create(fileName)

	defer func() {
		f.Close()
		cmd := exec.Command("go", "fmt", fileName)
		cmd.Run()
	}()

	if err != nil {
		return err
	}

	f.WriteString(GenCommon(b.db))
	return nil
}
