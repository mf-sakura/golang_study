package database

import (
	"fmt"
	"sync"

	"sort"

	"github.com/jmoiron/sqlx"

	"github.com/mf-sakura/golang_study/interface/api/domain"
)

// onMemoryUserRepository is repository of user to save on memory
// mutex should be initialized.
type onMemoryUserRepository struct {
	users map[int]domain.User
	maxID int
	mu    sync.Mutex
}

// newOnMemoryUserRepository returns new OnMemoryUserRepository
func newOnMemoryUserRepository() UserRepository {
	return &onMemoryUserRepository{
		users: make(map[int]domain.User),
		mu:    sync.Mutex{},
	}
}

// Store is a function for creating a user.
func (r *onMemoryUserRepository) Store(db *sqlx.DB, u domain.User) (int, error) {
	// ロックを取得する。
	// 並行して起こったリクエストによる上書きを防ぐ目的
	r.mu.Lock()
	defer r.mu.Unlock()
	r.maxID++
	u.ID = r.maxID
	r.users[r.maxID] = u
	return r.maxID, nil
}

// FindAll is a function for getting all users.
// 実装は課題
// 出来ればuser_id昇順で返す
func (r *onMemoryUserRepository) FindAll(db *sqlx.DB) (domain.Users, error) {
	var sortedUsers []domain.User
	ids := []int{}

	// r.usersから取得するという発想がなくて時間かかってしまった
	for id := range r.users {
		ids = append(ids, id)
	}
	sort.Ints(ids)
	fmt.Printf("sortedKeys: %v", ids)

	for _, id := range ids {
		sortedUsers = append(sortedUsers, r.users[id])
	}

	return sortedUsers, nil
}
