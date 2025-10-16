package tcp

import (
	"github.com/Dimas862/xray-core/common"
	"github.com/Dimas862/xray-core/transport/internet"
)

func init() {
	common.Must(internet.RegisterProtocolConfigCreator(protocolName, func() interface{} {
		return new(Config)
	}))
}

