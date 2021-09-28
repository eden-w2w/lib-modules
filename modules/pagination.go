package modules

type Pagination struct {
	// 分页大小
	Size int32 `in:"query" default:"10" name:"size,omitempty" validate:"@int32[-1,]"`
	// 偏移量
	Offset int32 `in:"query" default:"0" name:"offset,omitempty" validate:"@int32[0,]"`
}
