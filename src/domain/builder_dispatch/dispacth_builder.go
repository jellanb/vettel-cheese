package builder_dispatch

import "vettel-backend-app/src/domain/entity"

type DispatchBuilderInterfaces interface {
}

type DispatchBuilder struct {
	dispatch entity.Dispatch
}

func NewDispatchBuilder() DispatchBuilderInterfaces {
	return &DispatchBuilder{}
}

func (d *DispatchBuilder) BuildNewDispatch() entity.Dispatch {
	return d.dispatch
}
