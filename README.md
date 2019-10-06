# ProcessBar
### Install
```bash
go get github.com/tianyazc/ProcessBar
```
### Example:
```golang
package main

import 
(
"time"
"github.com/tianyazc/ProcessBar"
)

func main() {
	Pb := ProcessBar.NewPb()
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
☁  ProcessBar [master] ⚡  go run *.go
测试bar1
[==========]    100 %
测试bar2
[>>>>>>>>>>]    100 %
测试bar3
[◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆◆]  100 %
```