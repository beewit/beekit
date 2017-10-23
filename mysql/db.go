package mysql

import (
	"database/sql"
	"fmt"
	"strconv"

	"errors"
	"strings"
	"time"

	"github.com/beewit/beekit/conf"
	"github.com/beewit/beekit/log"
	"github.com/beewit/beekit/utils"
	"github.com/beewit/beekit/utils/convert"
	_ "github.com/go-sql-driver/mysql"
	"sort"
)

type SqlConnPool struct {
	DriverName     string
	DataSourceName string
	MaxOpenConns   int64
	MaxIdleConns   int64
	SqlDB          *sql.DB // 连接池
}

var (
	DB  *SqlConnPool
	cfg = conf.New("config.json")
)

func init() {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&loc=Local",
		cfg.Get("mysql.user").(string),
		cfg.Get("mysql.password").(string),
		cfg.Get("mysql.host").(string),
		cfg.Get("mysql.database").(string))

	maxOpenConns, _ := strconv.ParseInt(cfg.Get("mysql.maxOpenConns").(string), 10, 32)
	maxIdleConns, _ := strconv.ParseInt(cfg.Get("mysql.maxIdleConns").(string), 10, 32)

	DB = &SqlConnPool{
		DriverName:     "mysql",
		DataSourceName: dataSourceName,
		MaxOpenConns:   maxOpenConns,
		MaxIdleConns:   maxIdleConns,
	}
	if err := DB.open(); err != nil {
		panic("init db failed")
	}
}

// 封装的连接池的方法
func (p *SqlConnPool) open() error {
	var err error
	p.SqlDB, err = sql.Open(p.DriverName, p.DataSourceName)
	p.SqlDB.SetMaxOpenConns(int(p.MaxOpenConns))
	p.SqlDB.SetMaxIdleConns(int(p.MaxIdleConns))
	return err
}

func (p *SqlConnPool) Close() error {
	return p.SqlDB.Close()
}

func (p *SqlConnPool) QueryPage(page *utils.PageTable, args ...interface{}) (*utils.PageData, error) {
	if page.Where != "" {
		page.Where = " WHERE " + page.Where
	}
	if page.Order != "" {
		page.Order = " ORDER BY " + page.Order
	}
	sql := fmt.Sprintf("SELECT COUNT(1) count FROM  %s %s ", page.Table, page.Where)
	m, err := p.Query(sql, args...)
	if err != nil {
		log.Logger.Error(err.Error())
		return nil, err
	}
	c := convert.MustInt64(m[0]["count"])

	sql = fmt.Sprintf("SELECT %s FROM %s %s %s limit %d,%d", page.Fields, page.Table, page.Where, page.Order, (page.PageIndex-1)*page.PageSize, page.PageSize)
	log.Logger.Info(sql)
	m, err = p.Query(sql, args...)
	if err != nil {
		log.Logger.Error(err.Error())
		return nil, err
	}
	return &utils.PageData{
		PageIndex: page.PageIndex,
		PageSize:  page.PageSize,
		Count:     c,
		Data:      m,
	}, nil
}

func (p *SqlConnPool) Query(queryStr string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := p.SqlDB.Query(queryStr, args...)
	defer rows.Close()
	if err != nil {
		return []map[string]interface{}{}, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	list := []map[string]interface{}{}
	// 这里需要初始化为空数组，否则在查询结果为空的时候，返回的会是一个未初始化的指针
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		ret := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				ret[columns[i]] = nil
			} else {
				switch val := (*scanArgs[i].(*interface{})).(type) {
				case []byte:
					ret[columns[i]] = string(val)
					break
				case time.Time:
					ret[columns[i]] = val.Format("2006-01-02 15:04:05")
					break
				default:
					ret[columns[i]] = val
				}
			}
		}
		list = append(list, ret)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (p *SqlConnPool) execute(sqlStr string, args ...interface{}) (sql.Result, error) {
	return p.SqlDB.Exec(sqlStr, args...)
}

func (p *SqlConnPool) Update(updateStr string, args ...interface{}) (int64, error) {
	result, err := p.execute(updateStr, args...)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	return affect, err
}

func (p *SqlConnPool) Insert(insertStr string, args ...interface{}) (int64, error) {
	result, err := p.execute(insertStr, args...)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	return lastid, err

}

func (p *SqlConnPool) InsertMap(table string, m map[string]interface{}) (int64, error) {
	c := len(m)
	var keys = make([]string, c)
	var pars = make([]string, c)
	var values = make([]interface{}, c)
	i := 0

	var keysSort []string
	for k := range m {
		keysSort = append(keysSort, k)
	}
	sort.Strings(keysSort)

	for _, k := range keysSort {
		keys[i] = k
		pars[i] = "?"
		values[i] = m[k]
		i++
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s)VALUES(%s)", table, strings.Join(keys, ","), strings.Join(pars, ","))
	println(sql)
	println(convert.ToString(values))
	result, err := p.execute(sql, values...)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	return lastid, err

}

func (p *SqlConnPool) InsertMapList(table string, ms []map[string]interface{}) (int64, error) {
	var key string
	var par = make([]string, len(ms))
	var values = make([]interface{}, len(ms)*len(ms[0]))

	for j := 0; j < len(ms); j++ {
		c := len(ms[j])
		var keys = make([]string, c)
		var pars = make([]string, c)
		i := 0

		var keysSort []string
		for k := range ms[j] {
			keysSort = append(keysSort, k)
		}
		sort.Strings(keysSort)

		for _, k := range keysSort {
			keys[i] = k
			pars[i] = "?"
			l := i + (j * len(ms[j]))
			values[l] = ms[j][k]
			i++
			fmt.Println("Key:", k, "Value:", values[l])
		}
		if key == "" {
			key = strings.Join(keys, ",")
		}

		par[j] = fmt.Sprintf("(%s)", strings.Join(pars, ","))

	}
	sql := fmt.Sprintf("INSERT INTO %s (%s)VALUES%s", table, key, strings.Join(par, ","))
	println(sql)
	result, err := p.execute(sql, values...)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	return lastid, err

}

func (p *SqlConnPool) Delete(deleteStr string, args ...interface{}) (int64, error) {
	result, err := p.execute(deleteStr, args...)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	return affect, err
}

type SqlConnTransaction struct {
	SqlTx *sql.Tx // 单个事务连接
}

//// 开启一个事务
func (p *SqlConnPool) Begin() (*SqlConnTransaction, error) {
	var oneSqlConnTransaction = &SqlConnTransaction{}
	var err error
	if pingErr := p.SqlDB.Ping(); pingErr == nil {
		oneSqlConnTransaction.SqlTx, err = p.SqlDB.Begin()
	}
	return oneSqlConnTransaction, err
}

func (p *SqlConnPool) Tx(f func(tx *SqlConnTransaction), errFunc func(err error)) error {
	tx, _ := p.Begin()
	return tx.Tx(f, errFunc)
}

func (t *SqlConnTransaction) Tx(f func(tx *SqlConnTransaction), errFunc func(err error)) error {
	defer func() {
		if err := recover(); err != nil {
			e := fmt.Sprintf("%v", err)
			t.Rollback()
			errFunc(errors.New(e))
		}
	}()
	f(t)
	return t.Commit()
}

// 封装的单个事务连接的方法
func (t *SqlConnTransaction) Rollback() error {
	return t.SqlTx.Rollback()
}

func (t *SqlConnTransaction) Commit() error {
	return t.SqlTx.Commit()
}

func (t *SqlConnTransaction) Query(queryStr string, args ...interface{}) ([]map[string]interface{}, error) {
	rows, err := t.SqlTx.Query(queryStr, args...)
	defer rows.Close()
	if err != nil {
		return []map[string]interface{}{}, err
	}
	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}
	values := make([]interface{}, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	list := []map[string]interface{}{}
	// 这里需要初始化为空数组，否则在查询结果为空的时候，返回的会是一个未初始化的指针
	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			return nil, err
		}

		ret := make(map[string]interface{})
		for i, col := range values {
			if col == nil {
				ret[columns[i]] = nil
			} else {
				switch val := (*scanArgs[i].(*interface{})).(type) {
				case []byte:
					ret[columns[i]] = string(val)
					break
				case time.Time:
					ret[columns[i]] = val.Format("2006-01-02 15:04:05")
					break
				default:
					ret[columns[i]] = val
				}
			}
		}
		list = append(list, ret)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return list, nil
}

func (t *SqlConnTransaction) execute(sqlStr string, args ...interface{}) (sql.Result, error) {
	return t.SqlTx.Exec(sqlStr, args...)
}

func (t *SqlConnTransaction) Update(updateStr string, args ...interface{}) (int64, error) {
	result, err := t.execute(updateStr, args...)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	return affect, err
}

func (t *SqlConnTransaction) Insert(insertStr string, args ...interface{}) (int64, error) {
	result, err := t.execute(insertStr, args...)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	return lastid, err
}

func (t *SqlConnTransaction) InsertMap(table string, m map[string]interface{}) (int64, error) {
	c := len(m)
	var keys = make([]string, c)
	var pars = make([]string, c)
	var values = make([]interface{}, c)
	i := 0

	var keysSort []string
	for k := range m {
		keysSort = append(keysSort, k)
	}
	sort.Strings(keysSort)

	for _, k := range keysSort {
		keys[i] = k
		pars[i] = "?"
		values[i] = m[k]
		i++
	}
	sql := fmt.Sprintf("INSERT INTO %s (%s)VALUES(%s)", table, strings.Join(keys, ","), strings.Join(pars, ","))
	println(sql)
	return t.Insert(sql, values...)
}

func (t *SqlConnTransaction) InsertMapList(table string, ms []map[string]interface{}) (int64, error) {

	var key string
	var par = make([]string, len(ms))
	var values = make([]interface{}, len(ms)*len(ms[0]))

	for j := 0; j < len(ms); j++ {
		c := len(ms[j])
		var keys = make([]string, c)
		var pars = make([]string, c)
		i := 0

		var keysSort []string
		for k := range ms[j] {
			keysSort = append(keysSort, k)
		}
		sort.Strings(keysSort)

		for _, k := range keysSort {
			keys[i] = k
			pars[i] = "?"
			l := i + (j * len(ms[j]))
			values[l] = ms[j][k]
			i++
			fmt.Println("Key:", k, "Value:", values[l])
		}
		if key == "" {
			key = strings.Join(keys, ",")
		}

		par[j] = fmt.Sprintf("(%s)", strings.Join(pars, ","))

	}
	sql := fmt.Sprintf("INSERT INTO %s (%s)VALUES%s", table, key, strings.Join(par, ","))
	println(sql)
	result, err := t.execute(sql, values...)
	if err != nil {
		return 0, err
	}
	lastid, err := result.LastInsertId()
	return lastid, err

}

func (t *SqlConnTransaction) Delete(deleteStr string, args ...interface{}) (int64, error) {
	result, err := t.execute(deleteStr, args...)
	if err != nil {
		return 0, err
	}
	affect, err := result.RowsAffected()
	return affect, err
}

// others
func bytes2RealType(src []byte, columnType string) interface{} {
	srcStr := string(src)
	var result interface{}
	switch columnType {
	case "TINYINT":
		fallthrough
	case "SMALLINT":
		fallthrough
	case "INT":
		fallthrough
	case "BIGINT":
		result, _ = strconv.ParseInt(srcStr, 10, 64)
	case "CHAR":
		fallthrough
	case "VARCHAR":
		fallthrough
	case "BLOB":
		fallthrough
	case "TIMESTAMP":
		fallthrough
	case "DATETIME":
		result = srcStr
	case "FLOAT":
		fallthrough
	case "DOUBLE":
		fallthrough
	case "DECIMAL":
		result, _ = strconv.ParseFloat(srcStr, 64)
	default:
		result = nil
	}
	return result
}
