package models

type Product struct {
	Id int64 `gorm:"primaryKey;autoIncrement" json:"id"`
	NamaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	Deskripsi string `gorm:"type:text" json:"deskripsi"`
	Gambar string `gorm:"type:text" json:"gambar"`
}