package snowflake

import (
	"time"

	"github.com/bwmarrin/snowflake"
)

type Snowflake struct {
	node *snowflake.Node
}

// 雪花算法 生成全局唯一ID

func Init(startTime string, format string, machineID int64) (*Snowflake, error) {
	st, err := time.Parse(format, startTime)
	if err != nil {
		return nil, err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	node, err := snowflake.NewNode(machineID)
	if err != nil {
		return nil, err
	}
	return &Snowflake{node: node}, nil
}

func (sn *Snowflake) GetID() int64 {
	return sn.node.Generate().Int64()
}
