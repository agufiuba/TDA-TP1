def profit(prices):
	for i, p in enumerate(prices):
		if i == 0:
			m = p
			mi = i
			c = p
			ci = i
			v = p
			vi = i
		else:
			if p >= m:
				if p - m > v - c:
					c = m
					ci = mi
					v = p
					vi = i
			else:
				m = p
				mi = i
	return ci, vi
