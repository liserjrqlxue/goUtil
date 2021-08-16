package dataframeUtil

func JoinDb2MapArray(left []map[string]string, db map[string][]map[string]string, key string) []map[string]string {
	var join []map[string]string
	for _, item := range left {
		var adds, ok = db[item[key]]
		if ok {
			for _, add := range adds {
				var joinItem = make(map[string]string)
				for k, v := range item {
					joinItem[k] = v
				}
				for k, v := range add {
					joinItem[k] = v
				}
				join = append(join, joinItem)
			}
		}
	}
	return left
}

func JoinMapArray(left, right []map[string]string, lKey, rKey string) []map[string]string {
	var db = make(map[string][]map[string]string)
	for _, m := range right {
		var key = m[rKey]
		db[key] = append(db[key], m)
	}
	return JoinDb2MapArray(left, db, lKey)
}
