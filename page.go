package easysql

// Condition ...
type Condition map[string]interface{}

// PageFunc ...
type PageFunc func(Condition, int64, int64) (interface{}, int64, error)

// PageRsp ...
type PageRsp struct {
	Data  interface{}
	Count int64
}
