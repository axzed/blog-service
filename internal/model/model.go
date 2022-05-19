package model

// 公共的model
type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	CreatedOn  uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified_on"`
	DeleteOn   uint32 `json:"delete_on"`
	IsDel      uint8  `json:"is_del"`
}
