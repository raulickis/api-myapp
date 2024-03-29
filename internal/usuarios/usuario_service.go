package usuarios

import (
	"context"
	"github.com/raulickis/api-myapp/internal/database"
	"github.com/raulickis/api-myapp/internal/database/domains"
)

func ListUsers(ctx context.Context) (*[]domains.Usuario, error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	usuarios := &[]domains.Usuario{}
	err := db.
		Preload("Enderecos").
		Order("id").
		Find(usuarios).Error
	return usuarios, err
}

func GetUser(ctx context.Context,id int) (*domains.Usuario, error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	usuario := &domains.Usuario{}
	err := db.Where("id = ?", id).
		Preload("Enderecos").
		Find(usuario).Error
	return usuario, err
}

func RegisterUser(ctx context.Context, usuario *domains.Usuario) (*domains.Usuario, error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	err := db.Create(usuario).Error
	return usuario, err
}

func UpdateUser(ctx context.Context, usuario *domains.Usuario) (*domains.Usuario, error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	err := db.Save(usuario).Error
	return usuario, err
}

func DeleteUser(ctx context.Context, usuario *domains.Usuario) (error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	err := db.Delete(usuario).Error
	return err
}
