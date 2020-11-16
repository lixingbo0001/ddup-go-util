package where

func News(whs ...*Where) []*Where {
	wh := []*Where{}
	wh = append(wh, whs...)

	return wh
}
