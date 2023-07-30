package snowflake

import (
	"fmt"
	"time"

	sf "github.com/bwmarrin/snowflake"
)

var node *sf.Node

// 基于雪花算法生成用户ID
func Init(startTime string, machineID int64) (err error) {
	var st time.Time
	if st, err = time.Parse("2006-01-02", startTime); err != nil {
		fmt.Println("time parse failed. err :" + err.Error())
		return err
	}

	//初始化开始的时间
	sf.Epoch = st.UnixNano() / 1000000
	//指定机器码
	node, err = sf.NewNode(machineID)
	return
}

func GenID() int64 {
	return node.Generate().Int64()
}
