package enderecos

import (
	"context"
	"github.com/raulickis/api-myapp/internal/database"
	"github.com/raulickis/api-myapp/internal/database/domains"
)

func ListAddresses(ctx context.Context) (*[]domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	enderecos := &[]domains.Endereco{}
	err := db.
		Order("id").
		Find(enderecos).Error
	return enderecos, err
}

func ListAddressesByUser(ctx context.Context, userId int) (*[]domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	enderecos := &[]domains.Endereco{}
	err := db.Where("usuario_id = ?", userId).
		Order("id").
		Find(enderecos).Error
	return enderecos, err
}

func GetAddress(ctx context.Context, id int) (*domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	endereco := &domains.Endereco{}
	err := db.Where("id = ?", id).
		Find(endereco).Error
	return endereco, err
}

func RegisterAddress(ctx context.Context, endereco *domains.Endereco) (*domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	err := db.Create(endereco).Error
	return endereco, err
}

func UpdateAddress(ctx context.Context, endereco *domains.Endereco) (*domains.Endereco, error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	err := db.Save(endereco).Error
	return endereco, err
}

func DeleteAddress(ctx context.Context, endereco *domains.Endereco) (error) {
	var db = (&database.Repository{}).GetInstanceCtx(ctx)
	err := db.Delete(endereco).Error
	return err
}
