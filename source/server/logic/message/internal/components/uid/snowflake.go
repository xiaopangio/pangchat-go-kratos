package uid

import (
	"github.com/bwmarrin/snowflake"
	"github.com/go-kratos/kratos/v2/log"
	"message/internal/conf"
)

func NewUidGenerator(cf *conf.Bootstrap, helper *log.Helper) *snowflake.Node {
	node, err := snowflake.NewNode(cf.Snowflake.WorkerId)
	if err != nil {
		helper.Errorw("kind", "snowflake", "reason", "SNOWFLAKE_INIT_ERROR", "err", err)
	}
	return node
}
