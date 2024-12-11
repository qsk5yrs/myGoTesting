package domainservice

import "context"

type IServerDomainSvc struct {
	ctx context.Context
}

func NewIServerDomainSvc(ctx context.Context) *IServerDomainSvc {
	return &IServerDomainSvc{ctx: ctx}
}
