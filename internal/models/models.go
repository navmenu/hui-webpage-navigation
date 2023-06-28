package models

const AdminTable = "admin"

type Admin struct {
	ID             uint64 `gorm:"primary_key" json:"id"`
	Username       string `gorm:"unique" json:"username,omitempty"`
	Password       string `json:"password,omitempty"`
	Nickname       string `json:"nickname,omitempty"`
	CanCreateAdmin bool   `json:"can_create_admin,omitempty"`
	CanSelectAdmin bool   `json:"can_select_admin,omitempty"`
	CanEdit        bool   `json:"can_edit,omitempty"`
	CanSort        bool   `json:"can_sort,omitempty"`
	Token          string `gorm:"unique" json:"token,omitempty"`
	CreatedByUname string `gorm:"index" json:"created_by_uname,omitempty"`
}

func (*Admin) TableName() string {
	return AdminTable
}

const NaviTable = "navis"

type Navi struct {
	ID         uint64 `gorm:"primary_key" json:"id"`
	Name       string `gorm:"unique; uniqueIndex:udx_same_name_same_parent" json:"name,omitempty"`
	Sort       int32  `gorm:"index" json:"sort,omitempty"`
	ParentNvid uint64 `gorm:"default:0; index; uniqueIndex:udx_same_name_same_parent" json:"parent_nvid,omitempty"`
}

func (*Navi) TableName() string {
	return NaviTable
}

const NaviLvl2Table = "navi_lvl2"

type NaviLvl2 struct {
	ID       uint64 `gorm:"primary_key" json:"id"`
	NaviName string `gorm:"index:udx_navi_name_lvl2_name,unique;index:idx_navi_name_lvl2_sort" json:"navi_name,omitempty"` // 分类名称
	Name     string `gorm:"index:udx_navi_name_lvl2_name,unique;" json:"name,omitempty"`                                   // 名称：按钮上的简洁文字
	Text     string `json:"text,omitempty"`                                                                                // 中转页面内容部分包含以下元素 文案：后台自定义，比如“这是汇旺担保群，点击直达：”，文案单独一行；
	Link     string `json:"link,omitempty"`                                                                                // 中转页面内容部分包含以下元素 链接：后台自定义链接文案，自定义链接的地址；字号比前面一行大两号；
	IsEscrow bool   `json:"is_escrow,omitempty"`                                                                           // 需求-添加分类内容时可以勾选担保标志
	Sort     int32  `gorm:"index:idx_navi_name_lvl2_sort" json:"sort,omitempty"`                                           // 当前序号
}

func (*NaviLvl2) TableName() string {
	return NaviLvl2Table
}

const GuestSettingsTable = "guest_settings"

// GuestSettings 客户端设置
// 其实这个信息应该放在redis里面，设置30天过期，但目前功能极少只用1个db就行，没必要上两个组件吧
type GuestSettings struct {
	ID              uint64 `gorm:"primary_key" json:"id"`
	Ip              string `gorm:"index:udx_guest_ip_cookie,unique" json:"ip,omitempty"`     // 根据cookie和IP判断，是否首次登陆
	Cookie          string `gorm:"index:udx_guest_ip_cookie,unique" json:"cookie,omitempty"` // 根据cookie和IP判断，是否首次登陆
	IsNotRemindInfo bool   `json:"is_not_remind_info,omitempty"`                             ////根据cookie和IP判断，是否首次登陆
}

func (*GuestSettings) TableName() string {
	return GuestSettingsTable
}
