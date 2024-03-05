package channel

type CreateChannelReq struct {
	ChannelName  string  `json:"Channel_name"`
	Description  string  `json:"description"`
	ChannelCover string  `json:"Channel_cover"`
	MemberIDS    []int32 `json:"member_ids"`
}

type CreateChannelResp struct {
	Success     bool  `json:"success"`
	ChannelID   int32 `json:"Channel_id"`
	ChannelName int32 `json:"Channel_name"`
}
