package enderecos

import (
	"github.com/raulickis/api-myapp/internal/database"
	"github.com/raulickis/api-myapp/internal/database/domains"
)

func ListAddresses() (*[]domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstance()
	enderecos := &[]domains.Endereco{}
	err := db.
		Order("id").
		Find(enderecos).Error
	return enderecos, err
}

func ListAddressesByUser(userId int) (*[]domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstance()
	enderecos := &[]domains.Endereco{}
	err := db.Where("usuario_id = ?", userId).
		Order("id").
		Find(enderecos).Error
	return enderecos, err
}

func GetAddress(id int) (*domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstance()
	endereco := &domains.Endereco{}
	err := db.Where("id = ?", id).
		Find(endereco).Error
	return endereco, err
}

func RegisterAddress(endereco *domains.Endereco) (*domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstance()
	err := db.Create(endereco).Error
	return endereco, err
}

func UpdateAddress(endereco *domains.Endereco) (*domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstance()
	err := db.Save(endereco).Error
	return endereco, err
}

func DeleteAddress(endereco *domains.Endereco) (error) {
	var db = (&database.Repository{}).GetInstance()
	err := db.Delete(endereco).Error
	return err
}
