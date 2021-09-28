package goods

import (
	"github.com/eden-framework/sqlx"
	"github.com/eden-framework/sqlx/builder"

	"github.com/eden-w2w/lib-modules/modules"
)

type GetGoodsParams struct {
	modules.Pagination
}

func (p GetGoodsParams) Conditions(db sqlx.DBExecutor) builder.SqlCondition {
	return nil
}

func (p GetGoodsParams) Additions() []builder.Addition {
	var additions = make([]builder.Addition, 0)

	if p.Size != 0 {
		limit := builder.Limit(int64(p.Size))
		if p.Offset != 0 {
			limit.Offset(int64(p.Offset))
		}
		additions = append(additions, limit)
	}

	return additions
}
