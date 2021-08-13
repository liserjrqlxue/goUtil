package dataframeUtil

func JoinDb2MapArray(left []map[string]string, db map[string]map[string]string, key string) {
	for _, item := range left {
		var add, ok = db[item[key]]
		if ok {
			for k, v := range add {
				item[k] = v
			}
		}
	}
}

func JoinMapArray(left, right []map[string]string, lKey, rKey string) {
	var db = make(map[string]map[string]string)
	for _, m := range right {
		var key = m[rKey]
		db[key] = m
	}
	JoinDb2MapArray(left, db, lKey)
}
