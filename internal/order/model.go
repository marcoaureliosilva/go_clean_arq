package order

import "gorm.io/gorm"

// Order representa um pedido na aplicação
type Order struct {
	ID       uint   `gorm:"primaryKey"`
	Item     string `json:"item"`
	Quantity int    `json:"quantity"`
}

// Migrate cria a tabela no banco de dados
func Migrate(db *gorm.DB) error {
	return db.AutoMigrate(&Order{})
}
