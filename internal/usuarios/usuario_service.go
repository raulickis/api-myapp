package usuarios

import (
	"github.com/autorei/api-myapp/internal/database"
	"github.com/autorei/api-myapp/internal/database/domains"
)

func ListUsers() (*[]domains.Usuario, error) {
	var db = (&database.Repository{}).GetInstance()
	usuarios := &[]domains.Usuario{}
	err := db.
		Preload("Enderecos").
		Order("id").
		Find(usuarios).Error
	return usuarios, err
}

func GetUser(id int) (*domains.Usuario, error) {
	var db = (&database.Repository{}).GetInstance()
	usuario := &domains.Usuario{}
	err := db.Where("id = ?", id).
		Preload("Enderecos").
		Find(usuario).Error
	return usuario, err
}

func RegisterUser(usuario *domains.Usuario) (*domains.Usuario, error) {
	var db = (&database.Repository{}).GetInstance()
	err := db.Create(usuario).Error
	return usuario, err
}

func UpdateUser(usuario *domains.Usuario) (*domains.Usuario, error) {
	var db = (&database.Repository{}).GetInstance()
	err := db.Save(usuario).Error
	return usuario, err
}

func DeleteUser(usuario *domains.Usuario) (error) {
	var db = (&database.Repository{}).GetInstance()
	err := db.Delete(usuario).Error
	return err
}
