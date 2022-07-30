package utils

// PagePos 计算分页落点
func PagePos(total, page, size int) (offset, endpos int) {
	if page == 0 {
		page = 1
	}
	offset = (page - 1) * size
	if total > offset {
		if total > offset+size {
			endpos = offset + size
		} else {
			endpos = total
		}
	} else {
		offset = total
		endpos = total
	}
	return
}
