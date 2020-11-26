package where

type Where struct {
	Name string
	Con  string
	Val  interface{}
}

func New(name string) *Where {
	return &Where{
		Name: name,
	}
}

func (w *Where) Eq(value interface{}) *Where {

	w.Con = "="
	w.Val = value

	return w
}

func (w *Where) Neq(value interface{}) *Where {

	w.Con = "<>"
	w.Val = value

	return w
}

func (w *Where) Gt(value interface{}) *Where {

	w.Con = ">"
	w.Val = value

	return w
}

func (w *Where) Lt(value interface{}) *Where {

	w.Con = "<"
	w.Val = value

	return w
}

func (w *Where) Egt(value interface{}) *Where {

	w.Con = ">="
	w.Val = value

	return w
}

func (w *Where) Elt(value interface{}) *Where {

	w.Con = "<="
	w.Val = value

	return w
}

func (w *Where) In(value interface{}) *Where {

	w.Con = "in"
	w.Val = value

	return w
}

func (w *Where) Like(value interface{}) *Where {

	w.Con = "like"
	w.Val = value

	return w
}
