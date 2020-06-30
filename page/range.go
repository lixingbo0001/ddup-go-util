package page

type Ranger struct {
	offset int
	limit int
}

func (r *Ranger)Offset()int  {
	return  r.offset
}

func (r *Ranger)Limit()int  {
	return  r.limit
}

func (r *Ranger)Invalid()  bool{
	return  r.limit == 0
}
