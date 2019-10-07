# ProgressBar
### Install
```bash
go get github.com/tianyazc/ProgressBar
```
### Example:
```golang
package main

import 
(
"time"
"github.com/tianyazc/ProgressBar"
)

func main() {
	Pb := ProgressBar.NewPb()
	Pb.Bar("测试bar1","=",10)
	for i:=0;i<=10;i++ {
		Pb.Complete()
		time.Sleep(time.Second/10)
	}
	Pb.Bar("测试bar2",">",10)
	for i:=0;i<=10;i++ {
		Pb.Complete()
		time.Sleep(time.Second/10)
	}

	Pb.Bar("测试bar3","◆", 20)
	for i:=0;i<=20;i++ {
		Pb.Complete()
		time.Sleep(time.Second/10)
	}
}
// out
☁  ProgressBar [master] ⚡  go run *.go
测试bar1
[==========]    100 %
测试bar2
[>>>>>>>>>>]    100 %
测试bar3
[◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆]  100 %
```
