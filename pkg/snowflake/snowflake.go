package snowflake

import (
	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
	"time"
)

var node *snowflake.Node

func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	if st, err = time.Parse("2006-01-02", startTime); err != nil {
		zap.L().Error("time.Pares err,:", zap.Error(err))
		return err
	}
	snowflake.Epoch = st.UnixNano() / 1000000
	if node, err = snowflake.NewNode(machineID); err != nil {
		zap.L().Error("snowflake.NewNode err,:", zap.Error(err))
		return err
	}
	return nil
}

func GenID() int64 {
	return node.Generate().Int64()
}
