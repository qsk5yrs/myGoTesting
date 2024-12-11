package library

import (
	"context"
	"github.com/qsk5yrs/testing/common/logger"
	"github.com/qsk5yrs/testing/common/util/httptool"
	"github.com/tidwall/gjson"
)

type IServer struct {
	ctx context.Context
}

func NewIServer(ctx context.Context) *IServer {
	return &IServer{ctx: ctx}
}

func (i *IServer) GetScenesConfigNames(url string) (names []string) {
	log := logger.New(i.ctx)
	httpStatusCode, resBody, err := httptool.Get(i.ctx, url)
	if err != nil {
		log.Error("IServer GetScenes request error", "err", err, "httpStatusCode", httpStatusCode)
		return nil
	}
	result := gjson.ParseBytes(resBody).Get("#.name")
	result.ForEach(func(key, value gjson.Result) bool {
		names = append(names, value.String())
		return true
	})

	return
}

func (i *IServer) GetScenesLayerNames(url string) (layerNames []string) {
	log := logger.New(i.ctx)
	httpStatusCode, resBody, err := httptool.Get(i.ctx, url)
	if err != nil {
		log.Error("IServer GetScenes request error", "err", err, "httpStatusCode", httpStatusCode)
		return nil
	}
	result := gjson.ParseBytes(resBody).Get("layers.#.name")
	result.ForEach(func(key, value gjson.Result) bool {
		layerNames = append(layerNames, value.String())
		return true
	})

	return
}
