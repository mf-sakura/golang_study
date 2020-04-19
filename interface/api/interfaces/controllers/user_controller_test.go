package controllers

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/mf-sakura/golang_study/interface/api/domain"
	"github.com/mf-sakura/golang_study/interface/api/interfaces/database"
	"github.com/pkg/errors"
)

type mockUserRepository struct {
	isError bool
}

var (
	mockID             = 2
	userJSON           = `{"first_name":"John ","last_name":"Doe"}`
	invalidUserJSON    = `{"first_name": 1,"last_name": 2}`
	emptyFirstNameJSON = `{"first_name": "","last_name": "Doe"}`
	emptyLastNameJSON  = `{"first_name": "John","last_name": ""}`
)

func (r *mockUserRepository) Store(db *sqlx.DB, u domain.User) (int, error) {
	if r.isError {
		return 0, errors.New("error")
	}
	return mockID, nil
}

func (r *mockUserRepository) FindAll(db *sqlx.DB) (domain.Users, error) {
	return nil, nil
}

func newUserContext() echo.Context {
	e := echo.New()
	validReq := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(userJSON))
	validReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return e.NewContext(validReq, httptest.NewRecorder())
}

func newInvalidUserContext() echo.Context {
	e := echo.New()
	invalidReq := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(invalidUserJSON))
	invalidReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return e.NewContext(invalidReq, httptest.NewRecorder())
}

func newEmptyFirstNameContext() echo.Context {
	e := echo.New()
	invalidReq := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(emptyFirstNameJSON))
	invalidReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return e.NewContext(invalidReq, httptest.NewRecorder())
}

func newEmptyLastNameContext() echo.Context {
	e := echo.New()
	invalidReq := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(emptyLastNameJSON))
	invalidReq.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	return e.NewContext(invalidReq, httptest.NewRecorder())
}

// echoのtestについては以下を参照
// https://echo.labstack.com/guide/testing
// 課題: ユーザー名が空のケースを通す
func TestUserController_Create(t *testing.T) {
	type fields struct {
		db         *sqlx.DB
		repository database.UserRepository
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			fields: fields{
				db:         nil,
				repository: &mockUserRepository{},
			},
			args: args{c: newUserContext()},
		},
		{
			name: "異常系 mockRepositoryがエラーを返す",
			fields: fields{
				db:         nil,
				repository: &mockUserRepository{isError: true},
			},
			args:    args{c: newUserContext()},
			wantErr: true,
		},
		{
			name: "異常系 JSONの型が違う",
			fields: fields{
				db:         nil,
				repository: &mockUserRepository{},
			},
			args:    args{c: newInvalidUserContext()},
			wantErr: true,
		},
		{
			name: "異常系 ユーザ名(FirstName)が空文字のときエラーを返す",
			fields: fields{
				db:         nil,
				repository: &mockUserRepository{isError: true},
			},
			args:    args{c: newEmptyFirstNameContext()},
			wantErr: true,
		},
		{
			name: "異常系 ユーザ名(LastName)が空文字のときエラーを返す",
			fields: fields{
				db:         nil,
				repository: &mockUserRepository{isError: true},
			},
			args:    args{c: newEmptyLastNameContext()},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &UserController{
				db:         tt.fields.db,
				repository: tt.fields.repository,
			}
			if err := controller.Create(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserController.Create() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// 課題： UserController.Indexのカバレッジを100%にする
func TestUserController_Index(t *testing.T) {
	type fields struct {
		db         *sqlx.DB
		repository database.UserRepository
	}
	type args struct {
		c echo.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			controller := &UserController{
				db:         tt.fields.db,
				repository: tt.fields.repository,
			}
			if err := controller.Index(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("UserController.Index() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
