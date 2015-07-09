package controllers

// GORP relies on database/sql package (set of interfaces)
// which go-sql-driver/mysql package implements

import (
  "github.com/revel/revel"
  "github.com/coopernurse/gorp"
  "database/sql"
  _ "github.com/go-sql-driver/mysql"
  "fmt"
  "strings"
  "biddertator/app/models"
)

var InitDb func() = func(){
  connectionString := getConnectionString()
  if db, err := sql.Open("mysql", connectionString); err != nil {
    revel.ERROR.Fatal(err)
  } else {
    // we set Dbm variable defined in GorpController file gorp.go
    Dbm = &gorp.DbMap{ Db: db, Dialect: gorp.MySQLDialect{"InnoDB","UTF8"} }
    // Defines the table for use by GORP
    // using connection Dbm we define BidItem schema and create table if it doesn't exist
    defineBidItemTable(Dbm)
    if err := Dbm.CreateTablesIfNotExists(); err != nil {
      revel.ERROR.Fatal(err)
    }
  }
}

// using 'define' and not 'create' - because db is not actually created until call to Dbm.CreateTablesIfNotExists()
func defineBidItemTable(dbm *gorp.DbMap){
    // set "id" as primary key and autoincrement
    t := dbm.AddTable(models.BidItem{}).SetKeys(true, "id")
    // e.g. VARCHAR(25)
    t.ColMap("name").SetMaxSize(25)
}


// kick off func
// gives transaction support for controller actions
func init(){
  revel.OnAppStart(InitDb)
  revel.InterceptMethod((*GorpController).Begin, revel.BEFORE)
  revel.InterceptMethod((*GorpController).Commit, revel.AFTER)
  revel.InterceptMethod((*GorpController).Rollback, revel.FINALLY)
}

// helper functions
func getParamString(param string, defaultValue string) string {
  p, found := revel.Config.String(param)
  if !found {
    if defaultValue == "" {
      revel.ERROR.Fatal("Could not find parameter: " + param)
    } else {
      return defaultValue
    }
  }
  return p
}

func getConnectionString() string{
  host := getParamString("db.host", "")
  // port := getParamString("db.port", 3306)
  user := getParamString("db.user", "")
  pass := getParamString("db.password", "")
  port := getParamString("db.port", "3306")
  dbname := getParamString("db.name", "gogogo")
  protocol := getParamString("db.protocol", "tcp")
  dbargs := getParamString("dbargs", " ")

  if strings.Trim(dbargs, " ") != "" {
    dbargs = "?" + dbargs
  } else {
    dbargs = ""
  }

  return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
        user, pass, protocol, host, port, dbname, dbargs)
}
