package service

import (
	"context"

	"github.com/iWinston/qk-library/frame/model"

	"github.com/gogf/gf/net/ghttp"
	"gorm.io/gorm"
)

// 上下文管理服务,用于从ghttp.Request.RequestContext中方便地获取数据，数据格式model.RequestContext
var RequestContext = new(requestContextService)

type requestContextService struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *requestContextService) Init(r *ghttp.Request, customCtx *model.RequestContext) {
	r.SetCtxVar(model.ContextKey, customCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *requestContextService) Get(ctx context.Context) *model.RequestContext {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.RequestContext); ok {
		return localCtx
	}
	return nil
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *requestContextService) GetByRequest(r *ghttp.Request) *model.RequestContext {
	return s.Get(r.Context())
}

// SetTX 设置事务
func (s *requestContextService) SetTX(ctx context.Context, TX *gorm.DB) {
	s.Get(ctx).TX = TX
}

// GetTX 事务
func (s *requestContextService) GetTX(ctx context.Context) *gorm.DB {
	if s.Get(ctx) == nil {
		return nil
	}
	return s.Get(ctx).TX
}

// SetError 设置报错
func (s *requestContextService) SetError(ctx context.Context, err error) {
	s.Get(ctx).Err = err
}

// GetError 报错
func (s *requestContextService) GetError(ctx context.Context) error {
	if s.Get(ctx) == nil {
		return nil
	}
	return s.Get(ctx).Err
}

// GetDB 数据库
func (s *requestContextService) GetDB(ctx context.Context) *gorm.DB {
	if s.Get(ctx) == nil {
		return nil
	}
	return s.Get(ctx).DB
}