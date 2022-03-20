models
封装了各种实体类对象，与数据库交互的、与 UI 交互的等等，任何的实体类都应该放在这里。如：

```go
import "time"

type Article struct {
 ID        int64     `json:"id"`
 Title     string    `json:"title"`
 Content   string    `json:"content"`
 UpdatedAt time.Time `json:"updated_at"`
 CreatedAt time.Time `json:"created_at"`
}
```