package database

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/mf-sakura/golang_study/test/api/domain"
)

var db *sqlx.DB

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
	runTests := m.Run()
	os.Exit(runTests)
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
			name:    "正常系: FirstNameに「太郎」が含まれるあいまい検索ができる",
			args:    args{firstName: "太郎"},
			want:    domain.Users{{ID: 1, FirstName: "一太郎", LastName: "兼進"}},
			wantErr: false,
		},
		{
			name:    "正常系: FirstNameに「郎」が含まれるあいまい検索ができる",
			args:    args{firstName: "郎"},
			want:    domain.Users{{ID: 1, FirstName: "一太郎", LastName: "兼進"}, {ID: 2, FirstName: "次郎坊", LastName: "兼進"}},
			wantErr: false,
		},
		{
			name:    "正常系: FirstName「花子」に該当するデータがないのであいまい検索できるが空配列が返る",
			args:    args{firstName: "花子"},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := GetTestTransaction()
			// INSERT INTO `users` (first_name, last_name) VALUES ("一太郎", "兼進")
			// INSERT INTO `users` (first_name, last_name) VALUES ("次郎坊", "兼進")
			// 前提となるSeedデータの登録の仕方がわからない
			// 上記SQL相当のことがしたいだけなのだが
			// 今回はやりかたがわからなかったのでmysqlに入って上記のSQLを実行してからテストを行っています
			got, err := FirstNameLike(db, tt.args.firstName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FirstNameLike() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// 比較でIDまで比較されてしまうのがテストとして微妙。
			// 登録されたデータだけが確認できるようにしたいが……
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FirstNameLike() = %+v, want %+v", got, tt.want)
			}
			tx.Rollback()
		})
	}
}
