package biz

import "github.com/google/wire"

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewPingUsecase, NewAdminUsecase, NewNaviUsecase, NewNaviLvl2Usecase)
