package modelBase
import (
	"xfund/core/model/base/entity"
)

//执行数据迁移

func migration() {
	// 自动迁移模式
	DB.AutoMigrate(&entity.User{})
}
