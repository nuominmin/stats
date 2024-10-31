# stats


## 示例

```shell
go get github.com/nuominmin/stats
```

```go
import "github.com/nuominmin/stats"

func main() {
    stats.UpdateStats()
    fmt.Printf("新代币创建间隔: %v\n", stats.ReadInterval())
    fmt.Printf("平均创建间隔: %v\n", stats.ReadAvg())
}
```