package group

type CreateGroupReq struct {
	GroupName   string  `json:"group_name"`
	Description string  `json:"description"`
	GroupCover  string  `json:"group_cover"`
	MemberIDS   []int16 `json:"member_ids"`
}

type CreateGroupResp struct {
	Success   bool  `json:"success"`
	GroupID   int16 `json:"group_id"`
	GroupName int16 `json:"group_name"`
}
