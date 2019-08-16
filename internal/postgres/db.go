package postgres

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"taskulu/pkg"
)

type Database struct {
	log *pkg.Logger
	db  *gorm.DB
}

type Option struct {
	Host string
	Port string
	User string
	Pass string
	Db   string
}

type Taskulu struct {
	gorm.Model
	UserId   int `gorm:"primary_key"`
	Username string
	Password string
}

func New(log *pkg.Logger, option Option) *Database {
	url := fmt.Sprintf("host=%v port=%v user=%v dbname=%v password=%v sslmode=disable", option.Host, option.Port, option.Db, option.Pass)
	db, err := gorm.Open("postgres", url)
	if err != nil {
		log.Fatal("failed to connect database", err)
	}
	db.LogMode(true)
	// Migrate the schema
	db.AutoMigrate(&Taskulu{})

	return &Database{
		log: log,
		db:  db,
	}
}

func (d *Database) UpsertTaskulu(userId int, username, password string) error {
	t := Taskulu{
		UserId:   userId,
		Username: username,
		Password: password,
	}

	db := d.db.Model(&t).Update(t)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		db = d.db.Create(&t)
	}
	return db.Error
}

func (d *Database) UpsertTaskuluByUsername(userId int, username string) error {
	t := Taskulu{
		UserId:   userId,
		Username: username,
	}

	db := d.db.Model(&t).Where("user_id=?", userId).Update("username", username)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		db = d.db.Create(&t)
	}
	return db.Error
}

func (d *Database) UpsertTaskuluByPassword(userId int, password string) error {
	t := Taskulu{
		UserId:   userId,
		Password: password,
	}

	db := d.db.Model(&t).Where("user_id=?", userId).Update("password", password)
	if db.Error != nil {
		return db.Error
	}
	if db.RowsAffected == 0 {
		db = d.db.Create(&t)
	}
	return db.Error
}

func (d *Database) GetTaskuluByUser(userId int) ([]Taskulu, error) {
	var taskulu []Taskulu
	r := d.db.Where("user_id = ?", userId).Find(&taskulu)
	return taskulu, r.Error
}

func (d *Database) GetAllUserAuth() ([]Taskulu, error) {
	var taskulu []Taskulu
	r := d.db.Find(&taskulu)
	return taskulu, r.Error
}
