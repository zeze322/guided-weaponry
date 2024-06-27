package storage

import "context"

type Storage interface {
	ListCategory(ctx context.Context) ([]Category, error)
	WeaponParams(ctx context.Context, category string) ([]*Params, error)
	InsertWeapon(ctx context.Context, params *Params) error
	UpdateWeapon(ctx context.Context, params *Params) error
}
