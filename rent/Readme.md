ent 使用学习

安装 ent

``go get github.com/facebookincubator/ent/cmd/entc``

初始化(必须首字母大写)

``entc init User ``

创建字段

```go
func (User) Fields() []ent.Field {
  	return []ent.Field{
  		field.Int("age").Positive(),
  		field.String("name").Default("null"),
  	}
  }
```