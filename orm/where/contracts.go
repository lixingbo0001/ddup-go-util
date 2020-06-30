package where

type WhereReciever interface {
	Where(query interface{}, args ...interface{}) interface{}
}
