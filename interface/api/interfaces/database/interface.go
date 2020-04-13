package database

import (
	"github.com/pkg/errors"

	"github.com/jmoiron/sqlx"

	"github.com/mf-sakura/golang_study/interface/api/domain"
)

// Provider specifies provider
// Type Aliasを使用して、想定外のproviderが使われる可能性を減らす。
type Provider string

const (
	// ProviderMySQL uses MySQL
	ProviderMySQL Provider = "mysql"
	// ProviderOnMemory uses on memory
	ProviderOnMemory Provider = "on_memory"
)

// UserRepository is interface for persistence of user
type UserRepository interface {
	Store(*sqlx.DB, domain.User) (int, error)
	FindAll(*sqlx.DB) (domain.Users, error)
}

// NewUserRepository returns UserRepository corresponding to provider
func NewUserRepository(provider Provider) (UserRepository, error) {
	switch provider {
	case ProviderMySQL:
		return newMySQLUserRepository(), nil
	case ProviderOnMemory:
		return newOnMemoryUserRepository(), nil
	default:
		return nil, errors.New("unrecognized provider")
	}
}
