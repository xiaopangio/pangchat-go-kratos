package uid

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"math/rand"
	"relationship/internal/conf"
)

type FriendRequestUidGenerator struct {
	node *snowflake.Node
}

func NewFriendRequestUidGenerator(cf *conf.Bootstrap, helper *log.Helper) *FriendRequestUidGenerator {
	node, err := snowflake.NewNode(cf.Snowflake.WorkerId)
	if err != nil {
		helper.Errorw("kind", "snowflake", "reason", "SNOWFLAKE_INIT_ERROR", "err", err)
	}
	return &FriendRequestUidGenerator{node: node}
}
func (u *FriendRequestUidGenerator) Generate() snowflake.ID {
	return u.node.Generate()
}

type GroupRequestUidGenerator struct {
}

func NewGroupUidGenerator() *GroupRequestUidGenerator {

	return &GroupRequestUidGenerator{}
}
func (u *GroupRequestUidGenerator) Generate() string {
	//随机生成10-13位的群号
	l := rand.Intn(4) + 10
	res := make([]byte, l)
	for i := 0; i < l; i++ {
		res[i] = byte(rand.Intn(10))
	}
	return string(res)
}
