package ProcessBar

import (
	"fmt"
	"strings"
)

type ProcessBar struct {
	typestr string
	count int
}

func NewBar() *ProcessBar  {
	return &ProcessBar{
		typestr: "=",
		count: 100,
	}
}

func (pb *ProcessBar)Start()  {

}

func (pb *ProcessBar)Bar()  {
	for i:=0;i<pb.count;i++ {
		fill := strings.Repeat(" ",100-pb.count)
		fmt.Printf("%s%s",pb.typestr,fill)
		fmt.Printf("%d%",i)
	}
}

func (pb *ProcessBar)End()  {

}