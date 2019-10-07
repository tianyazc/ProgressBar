package ProgressBar

import (
	"fmt"
	"github.com/tianyazc/zlxGo/zlxmodule/color"
	"os"
	"strings"
)

type ProgressBar struct {
	typestr string
	count   int
	syncs   chan int
	color   color.Color
}

func NewPb() *ProgressBar {
	return &ProgressBar{
		syncs: make(chan int),
	}
}

func (pb *ProgressBar) Bar(describe, typestr string, count int) {
	go pb.BarOut(describe, typestr, count)
}

func (pb *ProgressBar) BarOut(describe, typestr string, count int) {
	_, _ = fmt.Fprintln(os.Stdout, describe)
	pb.out(typestr, count)
	_, _ = fmt.Fprintf(os.Stdout, "\n")
}

func (pb *ProgressBar) out(typestr string, count int) {
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

func (pb *ProgressBar) Complete() {
	x := <-pb.syncs
	_ = x
}
