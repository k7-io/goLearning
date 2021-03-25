package leetcode

import (
	"fmt"
	"sort"
	"testing"
)

func SortString(src string) (key string) {
	sBytes := []byte(src)
	sort.Slice(sBytes, func(i, j int) bool { return sBytes[i] < sBytes[j] })
	key = string(sBytes)
	return key
}

func groupAnagrams(strs []string) [][]string {
	mp := map[string][]string{}
	idxKeys := make([]string, 0)
	for _, s := range strs {
		key := SortString(s)
		_, ok := mp[key]
		if !ok {
			idxKeys = append(idxKeys, key)
		}
		mp[key] = append(mp[key], s)
	}
	data := make([][]string, 0, len(mp))
	fmt.Printf("idxKeys%v\n", idxKeys)
	for i := len(idxKeys) - 1; i >= 0; i-- {
		key := idxKeys[i]
		data = append(data, mp[key])
	}
	fmt.Printf("data%v\n", data)
	return data
}

func TestGroupAnagrams(t *testing.T) {
	// 切片
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	//expectData := [][]string{
	//	{"bat"},
	//	{"nat","tan"},
	//	{"ate","eat","tea"},
	//}
	data := groupAnagrams(strs)
	t.Logf("data:%v", data)
}
