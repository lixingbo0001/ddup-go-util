package page

func GetRange(page, size int, total int) *Ranger {
	if page <= 1 {
		page = 1
	}

	offset := size * (page - 1)

	if offset >= total {
		//设置为无效
		size = 0
	}

	return &Ranger{
		offset: offset,
		limit:  size,
	}
}

//当偏移量大于=总条数的时候，limit返回0
func GetOffsetLimit(page, size, total int) (offset, limit int) {
	if page <= 1 {
		page = 1
	}

	offset = size * (page - 1)

	if offset >= total {
		return offset, 0
	}

	return offset, size
}
