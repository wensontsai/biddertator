package controllers

import(
  "github.com/coopernurse/gorp"
  "github.com/revel/revel"
  // "github.com/go-sql-driver/mysql"
  // "database/sql"
)

// used to perform database creation
var(
  Dbm *gorp.DbMap
)

// Txn variable executes queries/commands against MYSQL in controller
type GorpController struct {
  *revel.Controller
  Txn *gorp.Transaction
}

func (c *GorpController) Begin() revel.Result{
  txn, err := Dbm.Begin()
  if err != nil {
    panic(err)
  }
  c.Txn = txn
  return nil
}

func (c *GorpController) Commit() revel.Result{
  if c.Txn == nil {
    return nil
  }
  // if err := c.Txn.Commit(); err != nil && err != sql.ErrTxdone {
  //     panic(err)
  // }
  c.Txn = nil
  return nil
}

func (c *GorpController) Rollback() revel.Result{
  if c.Txn == nil {
    return nil
  }
  // if err := c.Txn.Rollback(); err != nil && err != sql.ErrTxdone {
  //     panic(err)
  // }
  c.Txn = nil
  return nil
}
