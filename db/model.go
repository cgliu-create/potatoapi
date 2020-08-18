package db

import (
	"github.com/jinzhu/gorm"
)

var (
  Database *gorm.DB
)
type Product struct {
	gorm.Model
	Name  string
	Price uint
}

func MigrateProduct() {
  Database.AutoMigrate(&Product{})
}
func CreateProduct(p *Product) (err error) {
	if err = Database.Create(p).Error; err != nil {
		return err
	}
	return nil
}
func ReadAllProduct(p *[]Product) (err error) {
	if err = Database.Find(p).Error; err != nil {
		return err
	}
	return nil
}
func ReadIDProduct(p *Product, id string) (err error) {
	if err := Database.Where("id = ?", id).First(p).Error; err != nil {
		return err
	}
	return nil
}
func UpdateProduct(p *Product, id string) (err error) {
	Database.Save(p)
	return nil
}
func DeleteProduct(p *Product, id string) (err error) {
	Database.Where("id = ?", id).Delete(p)
	return nil
}


