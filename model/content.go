package model

import (
	"fmt"
)

//生成内容
func GenContent(t, db string, s []string) string {
	str := ""
	fTable := formatString(t)

	str += fmt.Sprintf("package %s \n\n", db)
	str += genImport()
	str += Sql2Struct(t, s)
	str += genTable(fTable, tablePrefix+t)
	str += genNewStruct(fTable)
	str += genModel(t, fTable)
	str += genGetInfo(t, fTable)
	str += genCreate(t, fTable)
	str += genUpdate(t, fTable)
	str += genDelete(t, fTable)
	str += genBatch(t, fTable)

	return str
}

//生成通用内容
func GenCommon(db string) string {
	return fmt.Sprintf("package %s\n\nimport clog \"micro-common/library/commom/log\"\n\nconst"+
		" (\n\tDBName = \"%s\"\n)\n\nvar (\n\tlog = clog.GetLogger()\n)", db, db)

}

//生成import
func genImport() string {
	return fmt.Sprintf("import (\n\t\"micro-common/plugins/db\"\n)\n")
}

//生成table
func genTable(s, fs string) string {
	return fmt.Sprintf("const (\n %sTableName = \"%s\"\n)\n\n", s, fs)
}

//生成new方法
func genNewStruct(s string) string {
	return fmt.Sprintf("func New%s(p string) *%s {\n\t return &%s{\n\tPostfix:p,\n\t}\n}\n\n", s, s, s)
}

//生成model
func genModel(s, fs string) string {
	return fmt.Sprintf("func (%s *%s) Model() *db.Model {\n\t "+
		"tn := %sTableName\n\t"+
		"if %s.Postfix != \"\" {\n\t\t tn += %s.Postfix \n\t}\n\t"+
		"return &db.Model{\n\t DB: DBName,\n\t Table: tn,\n\t} \n}\n\n", s[:1], fs, fs, s[:1], s[:1])
}

//生成详情
func genGetInfo(s, fs string) string {
	str := "//获取详情\n"
	str += fmt.Sprintf("func (%s *%s) GetInfo(id int64) error {\n\t err:= db.Slave(%s.Model())."+
		"Where(\"id=?\",id).First(%s).Error\n if db.CheckIsError(err){"+
		"\t\tlog.Infof(\"Get err:\",err)\n"+
		"}\n \t return err \n}\n\n", s[:1], fs, s[:1], s[:1])
	return str
}

//生成创建方法
func genCreate(s, fs string) string {
	str := "//创建单个信息\n"
	str += fmt.Sprintf("func (%s *%s) Create() error {\n\t err:= db.Master(%s.Model())."+
		"Save(%s).Error\n if db.CheckIsError(err){"+
		"\t\tlog.Infof(\"Create err:\",err)\n"+
		"}\n \t return err \n}\n\n", s[:1], fs, s[:1], s[:1])
	return str
}

//生成更新方法
func genUpdate(s, fs string) string {
	str := "//更新详情\n"
	str += fmt.Sprintf("func (%s *%s) Update() error {\n\t err:= db.Master(%s.Model())."+
		"Update(%s).Error\n if db.CheckIsError(err){"+
		"\t\tlog.Infof(\"Update err:\",err)\n"+
		"}\n \t return err \n}\n\n", s[:1], fs, s[:1], s[:1])
	return str
}

//生成删除方法
func genDelete(s, fs string) string {
	str := "//删除详情\n"
	str += fmt.Sprintf("func (%s *%s) Delete(id int64) error {\n\t err:= db.Master(%s.Model())."+
		"Where(\"id=?\",id).Delete(%s).Error\n if db.CheckIsError(err){"+
		"\t\tlog.Infof(\"Delete err:\",err)\n"+
		"}\n \t return err \n}\n\n", s[:1], fs, s[:1], s[:1])
	return str
}

//生成批量方法
func genBatch(s, fs string) string {
	str := "//批量获取信息\n"
	str += fmt.Sprintf("func (%s *%s) "+
		"BatchList(condition map[string]interface{}, page, pageSize int64) (*[]%s, error) {\n\t "+
		"var %ss []%s \n\t"+
		"err:= db.Slave(%s.Model())."+
		"Where(condition).Offset((page - 1) * pageSize).Limit(pageSize).Find(&%ss).Error\n if db.CheckIsError(err){"+
		"\t\tlog.Infof(\"Batch err:\",err)\n"+
		"}\n \t return &%ss,err \n}\n\n", s[:1], fs, fs, s[:1], fs, s[:1], s[:1], s[:1])
	return str
}
