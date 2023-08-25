package redis

import (
	"context"
	"errors"
	"fmt"
	"math"
	"time"

	//rd "github.com/go-redis/redis/v8"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const (
	oneWeekInSeconds = 7 * 24 * 3600
	scorePreVote     = 432
)

var (
	ErrVoteTimeExpire = errors.New("投票时间已过了")
)

func CreatePost(postID int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	// 帖子时间
	err := rdb.ZAdd(ctx, getRedisKey(KeyPostTime), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postID,
	}).Err()
	if err != nil {
		zap.L().Debug("rdb.zadd post time", zap.Error(err))
	}
	// 帖子分数
	err = rdb.ZAdd(ctx, getRedisKey(KeyPostScore), redis.Z{
		Score:  0,
		Member: postID,
	}).Err()
	if err != nil {
		zap.L().Debug("rdb.zadd post score", zap.Error(err))
	}

	return err
}

// 投一票 +432分 86400/200 -> 200 张赞成票可以给帖子占首页一天
func VoteForPost(userID, postID string, value float64) error {
	// 判断投票限制
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	postTime := rdb.ZScore(ctx, getRedisKey(KeyPostTime), postID).Val()

	fmt.Println("postTime")
	fmt.Println(postTime)

	if float64(time.Now().Unix())-postTime > oneWeekInSeconds {
		//return ErrVoteTimeExpire
	}

	// 更新分数
	// 先查单前用户给当前帖子的投票记录
	ov := rdb.ZScore(ctx, getRedisKey(KeyPostVoted+postID), userID).Val()
	var dir float64
	if value > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - value) // 计算2次投票的差值
	_, err := rdb.ZIncrBy(ctx, getRedisKey(KeyPostScore), dir*diff*scorePreVote, postID).Result()
	if err != nil {
		return err
	}

	// 记录用户为该帖子投票
	if value == 0 {
		fmt.Println(777)
		_, err = rdb.ZRem(ctx, getRedisKey(KeyPostVoted+postID), userID).Result()
	} else {
		fmt.Println(8888)
		rdb.ZAdd(ctx, getRedisKey(KeyPostVoted+postID), redis.Z{
			Score:  value,
			Member: userID,
		}).Result()
	}

	return err
}
