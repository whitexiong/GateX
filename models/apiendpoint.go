package models

type APIEndpoint struct {
	ID             uint   `gorm:"primaryKey"`
	Path           string `gorm:"index:idx_path,unique"`
	Type           string
	HitCount       uint
	MinRequestTime uint64 `gorm:"type:bigint"`      // 以毫秒为单位
	MaxRequestTime uint64 `gorm:"type:bigint"`      // 以毫秒为单位
	RequestMethod  string `gorm:"type:varchar(10)"` // 限制为10个字符长度，足以存储常用的HTTP方法
}
