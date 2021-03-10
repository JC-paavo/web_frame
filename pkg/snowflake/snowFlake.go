package snowflake

import (
	"fmt"

	"github.com/bwmarrin/snowflake"
	"go.uber.org/zap"
)

func GenId() snowflake.ID {
	snowflake.Epoch = 1615383215000

	node, err := snowflake.NewNode(1)
	if err != nil {
		fmt.Println(err)
		zap.L().Error("id生成器失败...",
			zap.String("error", err.Error()))
	}

	// Generate a snowflake ID.
	id := node.Generate()
	return id

	// Print out the ID in a few different ways.
	//fmt.Printf("Int64  ID: %d\n", id)
	//fmt.Printf("String ID: %s\n", id)
	//fmt.Printf("Base2  ID: %s\n", id.Base2())
	//fmt.Printf("Base64 ID: %s\n", id.Base64())
	//
	//// Print out the ID's timestamp
	//unix := time.Unix(0, id.Time()*int64(time.Millisecond))
	//unix := time.Unix(id.Time(), 0)
	//
	//fmt.Printf("ID Time  : %s\n", unix.Format("2006-01-02 15:04:05"))
	//
	//// Print out the ID's node number
	//fmt.Printf("ID Node  : %d\n", id.Node())
	//
	//// Print out the ID's sequence number
	//fmt.Printf("ID Step  : %d\n", id.Step())
	//
	//// Generate and print, all in one.
	//fmt.Printf("ID       : %d\n", node.Generate().Int64())
}
