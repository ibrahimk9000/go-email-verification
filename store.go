package main

// DB or REDIS
var mp = make(map[string]bool)

func exist(h string) bool {
	if _, ok := mp[h]; ok {
		return true
	}
	return false
}
func checkvalue(h string) bool {
	if v, ok := mp[h]; ok {
		return v
	}
	return false
}
func addkey(h string) {
	mp[h] = false
}

func changestatus(h string) {
	mp[h] = true
}
