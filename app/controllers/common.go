package controllers

import (
    "github.com/revel/revel"
    "strconv"
    "biddertator/app/models"
)

func parseUintOrDefault(intStr string, _default uint64) uint64 {
    if value, err := strconv.ParseUint(intStr, 0, 64); err != nil {
        return _default
    } else {
        return value
    }
}

func parseIntOrDefault(intStr string, _default int64) int64 {
    if value, err := strconv.ParseInt(intStr, 0, 64); err != nil {
        return _default
    } else {
        return value
    }
}

func (c BidItemCtrl) List() revel.Result {
    lastId := parseIntOrDefault(c.Params.Get("lid"), -1)
    limit := parseUintOrDefault(c.Params.Get("limit"), uint64(25))
    biditems, err := c.Txn.Select(models.BidItem{},
        `SELECT * FROM BidItem WHERE Id > ? LIMIT ?`, lastId, limit)
    if err != nil {
        return c.RenderText(
            "Error trying to get records from DB.")
    }
    return c.RenderJson(biditems)
}

func (c BidItemCtrl) Update(id int64) revel.Result {
    biditem, err := c.parseBidItem()
    if err != nil {
        return c.RenderText("Unable to parse the BidItem from JSON.")
    }
    // Ensure the Id is set.
    biditem.Id = id
    success, err := c.Txn.Update(&biditem)
    if err != nil || success == 0 {
        return c.RenderText("Unable to update bid item.")
    }
    return c.RenderText("Updated %v", id)
}

func (c BidItemCtrl) Delete(id int64) revel.Result {
    success, err := c.Txn.Delete(&models.BidItem{Id: id})
    if err != nil || success == 0 {
        return c.RenderText("Failed to remove BidItem")
    }
    return c.RenderText("Deleted %v", id)
}


