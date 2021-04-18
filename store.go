package main

// DB or REDIS
var mp = make(map[string]bool)

func exist(h string) bool {
	if _, ok := mp[h]; ok {
		return true
	}
	return false
}
func checkValue(h string) bool {
	if v, ok := mp[h]; ok {
		return v
	}
	return false
}
func addKey(h string) {
	mp[h] = false
}

func changeStatus(h string) {
	mp[h] = true
}
