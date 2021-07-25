package qservice

import (
	"context"

	model "github.com/iWinston/qk-library/frame/qmodel"

	"github.com/gogf/gf/net/ghttp"
	"gorm.io/gorm"
)

// 上下文管理服务,用于从ghttp.Request.RequestContext中方便地获取数据，数据格式model.QContext
var QContext = new(qContextService)

type qContextService struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *qContextService) Init(r *ghttp.Request, ctx *model.QContext) {
	r.SetCtxVar(model.ContextKey, ctx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *qContextService) Get(ctx context.Context) *model.QContext {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.QContext); ok {
		return localCtx
	}
	return nil
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *qContextService) GetByRequest(r *ghttp.Request) *model.QContext {
	return s.Get(r.Context())
}

// SetTX 设置事务
func (s *qContextService) SetTX(ctx context.Context, TX *gorm.DB) {
	s.Get(ctx).TX = TX
}

// GetTX 事务
func (s *qContextService) GetTX(ctx context.Context) *gorm.DB {
	if s.Get(ctx) == nil {
		return nil
	}
	return s.Get(ctx).TX
}

// SetError 设置报错
func (s *qContextService) SetError(ctx context.Context, err error) {
	s.Get(ctx).Err = err
}

// GetError 报错
func (s *qContextService) GetError(ctx context.Context) error {
	if s.Get(ctx) == nil {
		return nil
	}
	return s.Get(ctx).Err
}

// GetDB 数据库
func (s *qContextService) GetDB(ctx context.Context) *gorm.DB {
	if s.Get(ctx) == nil {
		return nil
	}
	return s.Get(ctx).DB
}
