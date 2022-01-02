package kumex

import "net/http"

// KLineModel represents the k lines for a symbol.
type KLineModel []interface{}

// A KLinesModel is the set of *KLineModel.
type KLinesModel []*KLineModel

// Data are returned in grouped buckets based on requested type.
// Parameter #2 granularity is the type of granularity patterns(minute).
// Parameter #3 #4 startAt, endAt is millisecond.
func (as *ApiService) KLines(symbol, granularity string, startAt, endAt int64) (*ApiResponse, error) {
	
	params := make(map[string]string)
	params["symbol"] = symbol
	params["granularity"] = granularity
	params["from"] = IntToString(startAt)
	
	if endAt > 0 {
		params["to"] = IntToString(endAt)
	}
	
	req := NewRequest(http.MethodGet, "/api/v1/kline/query", params)
	return as.Call(req)
}
