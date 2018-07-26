package eddn

import (
	"encoding/json"
	"fmt"
	"strings"

	"git.fractalqb.de/fractalqb/ggja"
)

var cmdtCatCopy = map[string]bool{
	"$MARKET_category_chemicals;":            true,
	"$MARKET_category_consumer_items;":       true,
	"$MARKET_category_drugs;":                true,
	"$MARKET_category_foods;":                true,
	"$MARKET_category_industrial_materials;": true,
	"$MARKET_category_machinery;":            true,
	"$MARKET_category_medicines;":            true,
	"$MARKET_category_metals;":               true,
	"$MARKET_category_minerals;":             true,
	"$MARKET_category_salvage;":              true,
	"$MARKET_category_slaves;":               true,
	"$MARKET_category_technology;":           true,
	"$MARKET_category_textiles;":             true,
	"$MARKET_category_waste;":                true,
	"$MARKET_category_weapons;":              true,
}

var cmdtFromJMarket = map[string]string{
	"meanPrice":     "MeanPrice",
	"buyPrice":      "BuyPrice",
	"stock":         "Stock",
	"stockBracket":  "StockBracket",
	"sellPrice":     "SellPrice",
	"demand":        "Demand",
	"demandBracket": "DemandBracket",
}

func cmdtConvert(jMkt ggja.GenObj) (res ggja.GenObj, err error) {
	//TODO: from discord ecdn
	//msg = {
	// 'timestamp': je['timestamp'],
	// 'systemName': je['StarSystem'],
	// 'stationName': je['StationName'],
	// 'marketId': je['MarketID'],
	// 'commodities': [{
	//   'name': c['Name'][1:-5] if c['Name'][0] == '$' and c['Name'][-5:] == '_name' else c['Name'],
	//   '…': c['…'],
	//   'statusFlags': c['StatusFlags'] if 'StatusFlags' in c else []
	//  } for c in je['Items'] if 'nonmarketable' not in c['Category'].lower()]
	//}
	defer func() {
		if x := recover(); x != nil {
			res = nil
			j, _ := json.Marshal(jMkt)
			err = fmt.Errorf("%s: %s", err, string(j))
		}
	}()
	mkt := ggja.Obj{Bare: jMkt}
	mcat := mkt.MStr("Category")
	if cpy, ok := cmdtCatCopy[mcat]; !ok {
		return nil, fmt.Errorf("unknown category of market item: '%s'", mcat)
	} else if cpy {
		res = make(ggja.GenObj)
		name := mkt.MStr("Name")
		if strings.HasPrefix(name, "$") {
			name = name[1:]
		}
		if strings.HasSuffix(name, ";") {
			name = name[:len(name)-1]
		}
		if strings.HasSuffix(name, "_name") {
			name = name[:len(name)-5]
		}
		res["name"] = name
		for edcNm, mktNm := range cmdtFromJMarket {
			if tmp, ok := jMkt[mktNm]; ok {
				res[edcNm] = tmp
			}
		}
	}
	return res, err
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
			di, err := cmdtConvert(si)
			if err != nil {
				return err
			} else if di != nil {
				items = append(items, di)
			}
		}
	}
	msg["commodities"] = items
	return nil
}
