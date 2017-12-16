package entities

import (
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3" //for real
)

//var mydb *sql.DB
var engine *xorm.Engine

func init() {
	en, err := xorm.NewEngine("sqlite3", "./agendaDate.db")
	checkErr(err)
	engine = en
	engine.SetMapper(core.SameMapper{})
	u := &UserInfo{}
	m := &MeetingInfo{}
	p := &Participated{}
	exist, err2 := engine.IsTableExist(u)
	checkErr(err2)
	if !exist {
		err3 := engine.CreateTables(u)
		checkErr(err3)
	}
	exist, err2 = engine.IsTableExist(m)
	checkErr(err2)
	if !exist {
		err3 := engine.CreateTables(m)
		checkErr(err3)
	}
	exist, err2 = engine.IsTableExist(p)
	checkErr(err2)
	if !exist {
		err3 := engine.CreateTables(p)
		checkErr(err3)
	}
}

/*
// SQLExecer interface for supporting sql.DB and sql.Tx to do sql statement
type SQLExecer interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Prepare(query string) (*sql.Stmt, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

// DaoSource Data Access Object Source
type DaoSource struct {
	// if DB, each statement execute sql with random conn.
	// if Tx, all statements use the same conn as the Tx's connection
	SQLExecer
}
*/
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
