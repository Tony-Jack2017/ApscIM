package admin

type Admin struct {
	AdminID   int16  `json:"admin_id,omitempty"`
	AdminNo   int16  `json:"admin_no"`
	AdminName string `json:"admin_name"`
	Avatar    string `json:"avatar"`
}
