package qservice

import (
	"context"

	model "github.com/iWinston/qk-library/frame/qmodel"
	"github.com/iWinston/qk-library/qutil"

	"github.com/gogf/gf/net/ghttp"
	"gorm.io/gorm"
)

// 上下文管理服务,用于从ghttp.Request.RequestContext中方便地获取数据，数据格式model.ReqContext
var ReqContext = new(reqContextService)

type reqContextService struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *reqContextService) Init(r *ghttp.Request, ctx *model.ReqContext) {
	r.SetCtxVar(model.ContextKey, ctx)
	rCtx := r.Context()
	ctx.RCtx = rCtx
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *reqContextService) Get(ctx context.Context) *model.ReqContext {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.ReqContext); ok {
		return localCtx
	}
	return nil
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *reqContextService) GetByRequest(r *ghttp.Request) *model.ReqContext {
	return s.Get(r.Context())
}

// SetTX 设置事务
func (s *reqContextService) SetTX(ctx context.Context, TX *gorm.DB) {
	s.Get(ctx).TX = TX
}

// GetTX 事务
func (s *reqContextService) GetTX(ctx context.Context) *gorm.DB {
	if s.Get(ctx) == nil {
		return nil
	}
	return s.Get(ctx).TX
}

// SetError 设置报错
func (s *reqContextService) SetError(ctx context.Context, err error) {
	s.Get(ctx).Err = err
}

// GetError 报错
func (s *reqContextService) GetError(ctx context.Context) error {
	if s.Get(ctx) == nil {
		return nil
	}
	return s.Get(ctx).Err
}

// GetDB 数据库
func (s *reqContextService) GetDB(ctx context.Context) *gorm.DB {
	if s.Get(ctx) == nil {
		return nil
	}
	return s.Get(ctx).DB
}

// SetDB 数据库
func (s *reqContextService) SetDB(ctx context.Context, DB *gorm.DB) {
	s.Get(ctx).DB = DB
}

// SetError 设置报错
func (s *reqContextService) SetData(ctx context.Context, key string, value interface{}) {
	s.Get(ctx).Data[key] = value
}

// GetError 报错
func (s *reqContextService) GetData(ctx context.Context, key string) interface{} {
	if s.Get(ctx) == nil {
		return nil
	}
	v, ok := s.Get(ctx).Data[key]
	if !ok {
		return nil
	}
	return v
}

// SetBeforeModel 设置操作之前的Model
func (s *reqContextService) SetBeforeModel(ctx context.Context, result interface{}) {
	qCtx := s.Get(ctx)
	if qCtx != nil {
		beforeMap := qutil.StructToMap(result)
		qCtx.ActionHistory.Before = &beforeMap
	}

}

// SetBeforeModel 根据TX设置操作之前的Model
func (s *reqContextService) SetBeforeModelByTx(tx *gorm.DB, result interface{}) {
	if tx.Statement.Context != nil {
		s.SetBeforeModel(tx.Statement.Context, result)
	}
}

// SetBeforeModel 设置操作之后的Model
func (s *reqContextService) SetAfterModel(ctx context.Context, result interface{}) {
	qCtx := s.Get(ctx)
	if qCtx != nil {
		afterMap := qutil.StructToMap(result)
		qCtx.ActionHistory.After = &afterMap
	}

}

// SetAfterModel 根据TX设置操作之后的Model
func (s *reqContextService) SetAfterModelByTx(tx *gorm.DB, result interface{}) {
	if tx.Statement.Context != nil {
		s.SetAfterModel(tx.Statement.Context, result)
	}
}

// SetRowsAffected 设置操作影响行数
func (s *reqContextService) SetRowsAffected(ctx context.Context, rowsAffected uint) {
	qCtx := s.Get(ctx)
	if qCtx != nil {
		qCtx.ActionHistory.RowsAffected = rowsAffected
	}

}

// SetRowsAffectedByTx 根据Tx设置操作影响行数
func (s *reqContextService) SetRowsAffectedByTx(tx *gorm.DB) {
	if tx.Statement.Context != nil {
		s.SetRowsAffected(tx.Statement.Context, uint(tx.Statement.RowsAffected))
	}
}
