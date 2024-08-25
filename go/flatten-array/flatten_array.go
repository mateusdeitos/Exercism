package flatten

func Flatten(nested interface{}) []interface{} {
	r := make([]interface{}, 0)
	_, ok := nested.([]interface{})
	if !ok {
		return r
	}
	for _, v := range nested.([]interface{}) {
		if _, ok := v.([]interface{}); ok {
			r = append(r, Flatten(v)...)
			continue
		}

		if v == nil {
			continue
		}

		r = append(r, v)
	}
	return r
}
