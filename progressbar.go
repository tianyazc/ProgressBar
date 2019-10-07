package ProgressBar

import (
	"fmt"
	"github.com/tianyazc/zlxGo/zlxmodule/color"
	"os"
	"strings"
)

type ProcessBar struct {
	typestr string
	count   int
	syncs   chan int
	color   color.Color
}

func NewPb() *ProcessBar {
	return &ProcessBar{
		syncs: make(chan int),
	}
}

func (pb *ProcessBar) Bar(describe, typestr string, count int) {
	go pb.BarOut(describe, typestr, count)
}

func (pb *ProcessBar) BarOut(describe, typestr string, count int) {
	_, _ = fmt.Fprintln(os.Stdout, describe)
	pb.out(typestr, count)
	_, _ = fmt.Fprintf(os.Stdout, "\n")
}

func (pb *ProcessBar) out(typestr string, count int) {
	for i := 0; i <= count; i++ {
		fill := strings.Repeat(" ", count-i)
		tstr := strings.Repeat(typestr, i)
		_, _ = fmt.Fprint(os.Stdout, "\r[")
		_, _ = fmt.Fprintf(os.Stdout, "%s%s]", tstr, fill)
		p := int(i * 100 / count)
		_, _ = fmt.Fprintf(os.Stdout, "\t%d %%", p)
		pb.syncs <- i
	}
}

func (pb *ProcessBar) Complete() {
	x := <-pb.syncs
	_ = x
}
