package domainservice

import (
	"context"
	"github.com/qsk5yrs/testing/common/util"
	"github.com/qsk5yrs/testing/library"
	"net/url"
	"path/filepath"
)

type IServerDomainSvc struct {
	ctx     context.Context
	iServer *library.IServer
}

func NewIServerDomainSvc(ctx context.Context) *IServerDomainSvc {
	return &IServerDomainSvc{
		ctx:     ctx,
		iServer: library.NewIServer(ctx),
	}

}

// GenerateScenesJson 生成scenes.json文件
func (ids *IServerDomainSvc) GenerateScenesJson(prefixUrl, hostPath string) error {
	urlPath, _ := url.JoinPath(prefixUrl, "/scenes.json")
	parsePath, err := url.Parse(urlPath)
	if err != nil {
		return err
	}
	localPath := filepath.Join(hostPath, parsePath.Path)
	content, err := ids.iServer.GetUrlContentBytes(urlPath)
	if err != nil {
		return err
	}
	err = util.WriteBytesToFile(content, localPath)
	if err != nil {
		return err
	}

	return nil
}

func (ids *IServerDomainSvc) GenerateScenesNameJson(prefixUrl, hostPath string) error {
	scenesUrlPath, _ := url.JoinPath(prefixUrl, "/scenes.json")
	scenesNames := ids.iServer.GetScenesConfigNames(scenesUrlPath)
	for _, scenesName := range scenesNames {
		scenesNameUrlPath, _ := url.JoinPath(prefixUrl, "scenes", scenesName+".json")
		parsePath, err := url.Parse(scenesNameUrlPath)
		if err != nil {
			return err
		}
		localPath := filepath.Join(hostPath, parsePath.Path)
		content, err := ids.iServer.GetUrlContentBytes(scenesNameUrlPath)
		if err != nil {
			return err
		}
		err = util.WriteBytesToFile(content, localPath)
		if err != nil {
			return err
		}
	}

	return nil
}
