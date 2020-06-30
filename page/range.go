package page

type ranger struct {
	offset int
	limit int
}

func (r *ranger)Offset()int  {
	return  r.offset
}

func (r *ranger)Limit()int  {
	return  r.limit
}

func (r *ranger)Invalid()  bool{
	return  r.limit == 0
}
