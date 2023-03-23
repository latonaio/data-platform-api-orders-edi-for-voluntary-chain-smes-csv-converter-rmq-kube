package dpfm_api_output_formatter

import (
	"encoding/json"
	"os"

	"github.com/gocarina/gocsv"
	"golang.org/x/xerrors"
)

func ConvertToConcatMessage(filePath string) (*[]DataConcatenation, error) {
	data := make([]DataConcatenation, 0)
	f, err := os.OpenFile(filePath, os.O_RDONLY, 0)
	if err != nil {
		return nil, xerrors.Errorf("file open error: %w", err)
	}
	// defer f.Close()

	allHeaders := make([]OrdersHeader, 0)
	allItems := make([]OrdersItem, 0)

	err = gocsv.UnmarshalFile(f, &allHeaders)
	if err != nil {
		return nil, xerrors.Errorf("read order header error: %w", err)
	}
	f.Close()
	f, _ = os.OpenFile(filePath, os.O_RDONLY, 0)
	err = gocsv.UnmarshalFile(f, &allItems)
	if err != nil {
		return nil, xerrors.Errorf("read order items error: %w", err)
	}
	f.Seek(0, 0)

	allHeaders = toNull(allHeaders)
	allItems = toNull(allItems)
	allHeaders = getUniqueHeaders(&allHeaders)

	for headerIdx, v := range allHeaders {
		oID := v.ExchangedOrdersDocumentIdentifier
		items := getUniqueItem(&allItems, oID)
		if len(items) == 0 {
			return nil, xerrors.Errorf("items length is 0")
		}
		// itemIDs := getItemIDs(items, oID)
		orderAll := DataConcatenation{
			Header: allHeaders[headerIdx],
			Item:   items,
		}
		data = append(data, orderAll)
	}

	return &data, nil
}

func getOrderIDs(headers []OrdersHeader) []string {
	idCheck := make(map[string]struct{})
	for _, v := range headers {
		idCheck[v.ExchangedOrdersDocumentIdentifier] = struct{}{}
	}
	ids := make([]string, 0)
	for k := range idCheck {
		ids = append(ids, k)
	}
	return ids
}

func getUniqueHeaders(headers *[]OrdersHeader) []OrdersHeader {
	idCheck := make(map[string]struct{})
	unique := make([]OrdersHeader, 0, len(*headers))
	for i, v := range *headers {
		_, ok := idCheck[v.ExchangedOrdersDocumentIdentifier]
		if ok {
			continue
		}
		idCheck[v.ExchangedOrdersDocumentIdentifier] = struct{}{}
		unique = append(unique, (*headers)[i])
	}
	return unique
}

func getUniqueItem(items *[]OrdersItem, oID string) []OrdersItem {
	itemIDs := make([]string, 0)
	unique := make([]OrdersItem, 0, len(*items))
	for i, v := range *items {
		if v.ExchangedOrdersDocumentIdentifier != oID || isContain(itemIDs, v.ExchangedOrdersDocumentIdentifier) {
			continue
		}
		itemIDs = append(itemIDs, v.ExchangedOrdersDocumentIdentifier)
		unique = append(unique, (*items)[i])
	}
	return unique
}

func isContain[T comparable](obj []T, targ T) bool {
	for _, v := range obj {
		if v == targ {
			return true
		}
	}
	return false
}

func toNull[T any](obj []T) []T {
	b, _ := json.Marshal(obj)
	j := make([]map[string]interface{}, 0)
	json.Unmarshal(b, &j)
	for i := range j {
		for k, v := range j[i] {
			switch typedV := v.(type) {
			case int:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case int32:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case int64:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case float32:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case float64:
				if isZero(typedV) {
					j[i][k] = nil
				}
			case string:
				if isZero(typedV) {
					j[i][k] = nil
				}
			}
		}
	}
	b, _ = json.Marshal(j)
	json.Unmarshal(b, &obj)
	return obj
}
func isZero[T comparable](obj T) bool {
	var zero T
	return obj == zero
}
