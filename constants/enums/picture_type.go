package enums

//go:generate eden generate enum --type-name=PictureType
// api:enum
type PictureType uint8

// 图片类型
const (
	PICTURE_TYPE_UNKNOWN PictureType = iota
	PICTURE_TYPE__IMAGE              // 图片
	PICTURE_TYPE__VIDEO              // 视频
)
