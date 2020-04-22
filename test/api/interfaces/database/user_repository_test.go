package database

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-testfixtures/testfixtures/v3"
	"github.com/jmoiron/sqlx"
	"github.com/mf-sakura/golang_study/test/api/domain"
)

var (
	db       *sqlx.DB
	fixtures *testfixtures.Loader
)

func TestMain(m *testing.M) {
	conn, err := sqlx.Open("mysql", "root:rootpassword@tcp(127.0.0.1:3314)/golang_study_test")
	if err != nil {
		panic(err.Error())
	}

	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	db = conn

	fixtures, err = testfixtures.New(
		testfixtures.Database(db.DB),
		testfixtures.Dialect("mysql"),
		testfixtures.Directory("../../testdata/fixtures"),
	)
	if err != nil {
		panic(err)
	}

	runTests := m.Run()
	os.Exit(runTests)
}

func prepareTestDB() {
	fmt.Printf("%v", fixtures)
	if err := fixtures.Load(); err != nil {
		panic(err)
	}
}

func GetTestTransaction() *sqlx.Tx {
	tx := db.MustBegin()
	return tx
}

func TestStore(t *testing.T) {
	type args struct {
		u domain.User
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "正常系",
			args: args{
				u: domain.User{
					FirstName: "hogehoge",
					LastName:  "hogeohgoehgoehoh",
				},
			},
			wantErr: false,
			// Execなどをモックして異常系のテストしたいけど、interface使わないとしんどいので一旦保留
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := GetTestTransaction()
			fmt.Println("hfoawhef;oawj")
			_, err := Store(tx, tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("Store() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			tx.Rollback()
		})
	}
}

// 課題にするメソッド
func TestFirstNameLike(t *testing.T) {
	type args struct {
		firstName string
	}
	tests := []struct {
		name    string
		args    args
		want    domain.Users
		wantErr bool
	}{
		{
			name:    "正常系1",
			args:    args{firstName: "田"},
			want:    domain.Users{{ID: 1, FirstName: "田中", LastName: "一郎"}},
			wantErr: false,
		},
		{
			name:    "正常系2",
			args:    args{firstName: "鈴"},
			want:    domain.Users{{ID: 2, FirstName: "鈴木", LastName: "二郎"}},
			wantErr: false,
		},
	}

	prepareTestDB()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FirstNameLike(db, tt.args.firstName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstNameLike() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstNameLike() = %v, want %v", got, tt.want)
			}
		})
	}
}
