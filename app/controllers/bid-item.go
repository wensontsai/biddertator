package controllers

import (
    "biddertator/app/models"
    "github.com/revel/revel"
    "encoding/json"
)

type BidItemCtrl struct {
    GorpController
}

func (c BidItemCtrl) parseBidItem() (models.BidItem, error) {
    biditem := models.BidItem{}
    err := json.NewDecoder(c.Request.Body).Decode(&biditem)
    return biditem, err
}

func (c BidItemCtrl) Add() revel.Result {
  if biditem, err := c.parseBidItem(); err != nil {
    return c.RenderText("Unable to parse the BidItem from JSON.")
  } else {
    // Validate the model
    biditem.Validate(c.Validation)
    if c.Validation.HasErrors(){
      // Do something better here !
      return c.RenderText("You have error in your BidItem.")
    } else {
      if err := c.Txn.Insert(&biditem); err != nil {
        return c.RenderText("Error inserting record into database!")
      } else {
        return c.RenderJson(biditem)
      }
    }
  }
}

func (c BidItemCtrl) Get(id int64) revel.Result {
  biditem := new(models.BidItem)
  err := c.Txn.SelectOne(biditem, `SELECT * FROM BidItem WHERE id = ?`)
  if err != nil {
    return c.RenderText("Error. Item doesn't exist.")
  }
  return c.RenderJson(biditem)
}


