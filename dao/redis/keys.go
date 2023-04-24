package redis

const (
	Prefix             = "mall:"
	KeyPostTimeZSet    = "post:time"   //zset;帖子及发帖时间          //命名方法，目的是在多个公司合作时，区别不同的redis
	KeyPostScoreZSet   = "post:score"  //zset;帖子及投票时间
	keyPostVotedZSetPF = "post:voted:" //zset;记录用户及投票类型,参数是post id
	keyCommunitySetPF  = "community:"  //保存每个分区下帖子的id
	keyCategorySetPF   = "product:"    //保存每个类别下的id
	KeyFavoriteSetPF   = "favorite:"   //保存每个用户的收藏夹
	KeyOrderSetPF      = "order:"      //保存每个订单的id
)

// 给redis key加上前缀
func getRedisKey(key string) string {
	return Prefix + key
}
