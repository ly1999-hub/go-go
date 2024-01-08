package response

import "github.com/thoas/go-funk"

type Code struct {
	Key      string
	Messenge string
	Code     int
}

var notFoundKey = Code{
	Key:      CommonNotFound,
	Messenge: "không tìm thấy",
	Code:     -1,
}

func GetByKey(key string) Code {
	item := funk.Find(list, func(item Code) bool {
		return item.Key == key
	})
	if item == nil {
		return notFoundKey
	}
	return item.(Code)
}

var list []Code

func init() {
	list = append(list, common...)
}
