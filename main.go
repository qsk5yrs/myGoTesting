package main

import (
	"context"
	"github.com/qsk5yrs/testing/common/logger"
	"github.com/qsk5yrs/testing/library"
	"time"
)

func main() {
	// call func
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	iServer := library.NewIServer(ctx)
	// 获取scenes.json的名称
	scenesName := iServer.GetScenesConfigNames("http://localhost:8090/iserver/services/3D-Ltdzname/rest/realspace/scenes.json")
	for _, name := range scenesName {
		logger.New(ctx).Info("scenesName", "name", name)
	}

	// 获取图层名称
	layerNames := iServer.GetScenesLayerNames("http://localhost:8090/iserver/services/3D-Ltdzname/rest/realspace/scenes/Ltdzname.json")
	for i, name := range layerNames {
		//fmt.Printf("第%d个图层名称为：%s\n ", i, name)
		logger.New(ctx).Info("layerName", "index", i, "name", name)
	}
}
