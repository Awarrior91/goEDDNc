package eddn

import (
	"fmt"
)

var cmdtFromJMarket = map[string]string{
	"name":          "Name",
	"meanPrice":     "MeanPrice",
	"buyPrice":      "BuyPrice",
	"stock":         "Stock",
	"stockBracket":  "StockBracket",
	"sellPrice":     "SellPrice",
	"demand":        "Demand",
	"demandBracket": "DemandBracket",
}

func SetCommoditiesJ(msg map[string]interface{}, journal map[string]interface{}) error {
	if tmp, ok := journal["StarSystem"]; ok {
		msg["systemName"] = tmp
	} else {
		return fmt.Errorf("missing system name in commodities data: %s", journal)
	}
	if tmp, ok := journal["StationName"]; ok {
		msg["stationName"] = tmp
	} else {
		return fmt.Errorf("missing station name in commodities data: %s", journal)
	}
	if tmp, ok := journal["MarketID"]; ok {
		msg["marketId"] = tmp
	}
	var items []interface{}
	if tmp, ok := journal["Items"]; ok {
		itmls := tmp.([]interface{})
		items = make([]interface{}, 0, len(itmls))
		for _, src := range itmls {
			si := src.(map[string]interface{})
			di := make(map[string]interface{})
			for dnm, snm := range cmdtFromJMarket {
				if tmp, ok := si[snm].(interface{}); ok {
					di[dnm] = tmp
				} else {
					return fmt.Errorf("missing '%s' in %s", snm, si)
				}
			}
			items = append(items, di)
		}
	}
	msg["commodities"] = items
	return nil
}
