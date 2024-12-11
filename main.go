package main

import (
	"context"
	"fmt"
	"github.com/qsk5yrs/testing/common/logger"
	"github.com/qsk5yrs/testing/common/util/httptool"
	"github.com/tidwall/gjson"
	"time"
)

func main() {
	// call func
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Minute)
	// 获取scenes.json的名称
	scenesName := GetScenesConfigNames(ctx, "http://localhost:8090/iserver/services/3D-Ltdzname/rest/realspace/scenes.json")
	for _, name := range scenesName {
		fmt.Println("场景名称为: ", name)
	}

	// 获取图层名称
	layerNames := GetScenesNameJson(ctx, "http://localhost:8090/iserver/services/3D-Ltdzname/rest/realspace/scenes/Ltdzname.json")
	for i, name := range layerNames {
		fmt.Printf("第%d个图层名称为：%s\n ", i, name)
	}
}

func GetScenesConfigNames(ctx context.Context, url string) (names []string) {
	log := logger.New(ctx)
	httpStatusCode, resBody, err := httptool.Get(ctx, url)
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

func GetScenesNameJson(ctx context.Context, url string) (layerNames []string) {
	log := logger.New(ctx)
	httpStatusCode, resBody, err := httptool.Get(ctx, url)
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
