package user

import (
	"github.com/eden-framework/sqlx/builder"
	"github.com/eden-w2w/lib-modules/databases"
	"github.com/eden-w2w/lib-modules/modules"
)

type CreateUserByWechatSessionParams struct {
	// 微信OpenID
	OpenID string `in:"body" json:"openID" name:"openID"`
	// 微信UnionID
	UnionID string `in:"body" json:"unionID" name:"unionID"`
	// 微信SessionKey
	SessionKey string `in:"body" json:"sessionKey" name:"sessionKey"`
}

type UpdateUserInfoParams struct {
	// 用户名
	UserName string `in:"body" json:"userName" default:""`
	// 手机号
	Mobile string `in:"body" json:"mobile" default:""`
	// 昵称
	NickName string `in:"body" json:"nickName" default:""`
	// 头像地址
	AvatarUrl string `in:"body" json:"avatarUrl" default:""`
	// 推荐人ID
	RefererID uint64 `in:"body" json:"refererID,string" default:""`
	// 微信SessionKey
	SessionKey string `in:"body" json:"sessionKey" default:""`
}

func (p UpdateUserInfoParams) Diff(model *databases.User) (change bool) {
	change = false
	if p.UserName != "" && p.UserName != model.UserName {
		model.UserName = p.UserName
		change = true
	}
	if p.Mobile != "" && p.Mobile != model.Mobile {
		model.Mobile = p.Mobile
		change = true
	}
	if p.NickName != "" && p.NickName != model.NickName {
		model.NickName = p.NickName
		change = true
	}
	if p.AvatarUrl != "" && p.AvatarUrl != model.AvatarUrl {
		model.AvatarUrl = p.AvatarUrl
		change = true
	}
	if p.RefererID != 0 && model.RefererID == 0 && model.UserID != p.RefererID {
		model.RefererID = p.RefererID
		change = true
	}
	if p.SessionKey != "" && p.SessionKey != model.SessionKey {
		model.SessionKey = p.SessionKey
		change = true
	}
	return
}

type GetUsersParams struct {
	// 业务ID
	UserID uint64 `in:"query" name:"userID" default:""`
	// 推荐人ID
	RefererID uint64 `in:"query" name:"refererID" default:""`
	// 用户名
	UserName string `in:"query" name:"userName" default:""`
	// 手机号
	Mobile string `in:"query" name:"mobile" default:""`
	// 昵称
	NickName string `in:"query" name:"nickName" default:""`
	// 微信OpenID
	OpenID string `in:"query" name:"openID" default:""`
	// 微信UnionID
	UnionID string `in:"query" name:"unionID" default:""`

	modules.Pagination
}

func (p GetUsersParams) Conditions() builder.SqlCondition {
	var condition builder.SqlCondition
	table := databases.User{}

	if p.UserID != 0 {
		condition = builder.And(condition, table.FieldUserID().Eq(p.UserID))
	}
	if p.RefererID != 0 {
		condition = builder.And(condition, table.FieldRefererID().Eq(p.RefererID))
	}
	if p.UserName != "" {
		condition = builder.And(condition, table.FieldUserName().Eq(p.UserName))
	}
	if p.Mobile != "" {
		condition = builder.And(condition, table.FieldMobile().Eq(p.Mobile))
	}
	if p.NickName != "" {
		condition = builder.And(condition, table.FieldNickName().Like(p.NickName))
	}
	if p.OpenID != "" {
		condition = builder.And(condition, table.FieldOpenID().Eq(p.OpenID))
	}
	if p.UnionID != "" {
		condition = builder.And(condition, table.FieldUnionID().Eq(p.UnionID))
	}

	return condition
}

func (p GetUsersParams) Additions() []builder.Addition {
	var additions = make([]builder.Addition, 0)

	if p.Size != 0 {
		limit := builder.Limit(int64(p.Size))
		if p.Offset != 0 {
			limit = limit.Offset(int64(p.Offset))
		}
		additions = append(additions, limit)
	}

	additions = append(additions, builder.OrderBy(builder.DescOrder((&databases.Order{}).FieldCreatedAt())))

	return additions
}

type GetUserByNameOrOpenIDParams struct {
	// 关键字
	Keywords string `in:"query" name:"keywords"`
}

func (p GetUserByNameOrOpenIDParams) Conditions() builder.SqlCondition {
	model := databases.User{}
	condition := model.FieldNickName().Like(p.Keywords)
	condition = builder.Or(condition, model.FieldOpenID().Eq(p.Keywords))

	return condition
}

type CreateShippingAddressParams struct {
	// 用户ID
	UserID uint64 `json:"userID,string" default:""`
	// 收件人
	Recipients string `json:"recipients"`
	// 省市区街道
	District string `json:"district"`
	// 详细地址
	Address string `json:"address"`
	// 联系电话
	Mobile string `json:"mobile"`
}

type UpdateShippingAddressParams struct {
	// 业务ID
	ShippingID uint64 `json:"shippingID,string" default:""`
	// 收件人
	Recipients string `json:"recipients" default:""`
	// 省市区街道
	District string `json:"district" default:""`
	// 详细地址
	Address string `json:"address" default:""`
	// 联系电话
	Mobile string `json:"mobile" default:""`
}
