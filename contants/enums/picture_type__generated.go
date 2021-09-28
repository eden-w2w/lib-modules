package enums

import (
	"bytes"
	"encoding"
	"errors"

	github_com_eden_framework_enumeration "github.com/eden-framework/enumeration"
)

var InvalidPictureType = errors.New("invalid PictureType")

func init() {
	github_com_eden_framework_enumeration.RegisterEnums("PictureType", map[string]string{
		"VIDEO": "视频",
		"IMAGE": "图片",
	})
}

func ParsePictureTypeFromString(s string) (PictureType, error) {
	switch s {
	case "":
		return PICTURE_TYPE_UNKNOWN, nil
	case "VIDEO":
		return PICTURE_TYPE__VIDEO, nil
	case "IMAGE":
		return PICTURE_TYPE__IMAGE, nil
	}
	return PICTURE_TYPE_UNKNOWN, InvalidPictureType
}

func ParsePictureTypeFromLabelString(s string) (PictureType, error) {
	switch s {
	case "":
		return PICTURE_TYPE_UNKNOWN, nil
	case "视频":
		return PICTURE_TYPE__VIDEO, nil
	case "图片":
		return PICTURE_TYPE__IMAGE, nil
	}
	return PICTURE_TYPE_UNKNOWN, InvalidPictureType
}

func (PictureType) EnumType() string {
	return "PictureType"
}

func (PictureType) Enums() map[int][]string {
	return map[int][]string{
		int(PICTURE_TYPE__VIDEO): {"VIDEO", "视频"},
		int(PICTURE_TYPE__IMAGE): {"IMAGE", "图片"},
	}
}

func (v PictureType) String() string {
	switch v {
	case PICTURE_TYPE_UNKNOWN:
		return ""
	case PICTURE_TYPE__VIDEO:
		return "VIDEO"
	case PICTURE_TYPE__IMAGE:
		return "IMAGE"
	}
	return "UNKNOWN"
}

func (v PictureType) Label() string {
	switch v {
	case PICTURE_TYPE_UNKNOWN:
		return ""
	case PICTURE_TYPE__VIDEO:
		return "视频"
	case PICTURE_TYPE__IMAGE:
		return "图片"
	}
	return "UNKNOWN"
}

var _ interface {
	encoding.TextMarshaler
	encoding.TextUnmarshaler
} = (*PictureType)(nil)

func (v PictureType) MarshalText() ([]byte, error) {
	str := v.String()
	if str == "UNKNOWN" {
		return nil, InvalidPictureType
	}
	return []byte(str), nil
}

func (v *PictureType) UnmarshalText(data []byte) (err error) {
	*v, err = ParsePictureTypeFromString(string(bytes.ToUpper(data)))
	return
}
