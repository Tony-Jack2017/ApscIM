package channel

import "ApscIM/pkg/model/base/channel"

type channelServer struct {
}

func Start() {
	channel.NewChannelDatabase()
}
