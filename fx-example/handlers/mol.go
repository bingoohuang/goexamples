package handlers

import (
	"fx-example/cache"

	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttprouter"
	"go.uber.org/zap"
)

type MeaningOfLife struct {
	cache  cache.MeaningOfLifeCache
	logger *zap.Logger
}

func NewMeaningOfLifeHandler(cache cache.MeaningOfLifeCache, logger *zap.Logger) *MeaningOfLife {
	return &MeaningOfLife{
		cache:  cache,
		logger: logger,
	}
}

func (mol *MeaningOfLife) Handle(ctx *fasthttp.RequestCtx, _ fasthttprouter.Params) {
	res, err := mol.cache.LoadOrStore(func() (string, error) {
		return `这个概念通过许多相关问题体现出来，例如：“我为何在此”“什么是生命？”“生命的真谛是什么？”。
在历史长河中，它也是哲学，科学以及神学一直所思索的主题。前人在不同的文化环境与意识形态背景下也给出了很多的多元化答案。
阿尔贝·加缪指出，作为一个存在的人，人类用生命的价值和意义来说服自己：人的存在不是荒诞的。`, nil
	})
	if err != nil {
		ctx.Error(err.Error(), 500)
		return
	}
	ctx.SetBody([]byte(res))
}
