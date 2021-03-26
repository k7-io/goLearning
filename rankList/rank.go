package rankList

/*
rank: 排行榜
基于redis zset 实现.

api:
	1.参与排名.
		input: uid, score
	2.获取排名.
		input: uid.
		返回: 用户的名次，成绩.
	3.获取榜单.topN
	4.获取榜单all.分页.
test: 现在redis-cli>测试.
*/


type Item interface {
	// zset 集合的name
	Key() (key string)
	// element 的唯一标识
	Uid() (key string)
	// element 的成绩
	Score() (score float64)
}

type Rank interface {
	Join(key, uid string, score float64) (err error)
	Get(key, uid string) (index int, err error)
	Top(key string, n int) (dataItems []interface{}, err error)
	All(key string, pageIndex, pageSize int) (dataItems []interface{}, err error)
}


type RankItem struct {
	Key string
	Uid string
	Score float64
}

func (rankItem *RankItem) Join(key, uid string, score float64) (err error) {
	return
}

func (rankItem *RankItem) Top(key string, n int) (dataItems []interface{}, err error) {
	return
}

func (rankItem *RankItem) All(key string, pageIndex, pageSize int) (dataItems []interface{}, err error) {
	return
}

