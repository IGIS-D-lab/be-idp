package orm

import "fmt"

func (d DataBaseObject) SqlLogin(dbAddress, dbName string) (string, error) {
	/*
		Create SQL Login Query
		[username]:[password]@tcp(ipv4:pornNum)/[dbname]
	*/
	s := fmt.Sprintf("%s:%s@tcp(%s)/%s", d.Username, d.Password, dbAddress, dbName)
	return s, nil
}
