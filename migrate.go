package main

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"

	"github.com/daheige/thinkgo/mysql"
)

type Post struct {
	ID        uint   `gorm:"primary_key;auto_increment"` //unsigned int
	Uid       int    `gorm:"type:int;not null"`
	Title     string `gorm:"type:varchar(255);not null"`
	Content   string `gorm:"type:text;not null"`
	Type      uint8  `gorm:"type:tinyint;default 1;not null"`
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

//指定Post结构体对应的数据表为my_post
func (Post) TableName() string {
	return "my_posts"
}

type Model struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

type Prize struct {
	gorm.Model
	Name string
}

func (Prize) TableName() string {
	return "my_Prizes"
}

type User struct {
	Id       uint `gorm:"primary_key;auto_increment"` //对应数据表的自增id
	Username string
	Password string
	Email    string
	Phone    string
}

func (User) TableName() string {
	return "my_user"
}

//数据迁移简单代码示例
func main() {
	dbconf := &mysql.DbConf{
		Ip:           "127.0.0.1",
		Port:         3306,
		User:         "root",
		Password:     "root",
		Database:     "test",
		MaxIdleConns: 10,
		MaxOpenConns: 100,
		ParseTime:    true,
		SqlCmd:       true,
	}

	//设置db engine name
	dbconf.SetDbPool() //建立db连接池
	dbconf.SetEngineName("default")

	defer dbconf.Close() //关闭当前连接
	log.Println(dbconf)

	db, err := mysql.GetDbObj("default")
	if err != nil {
		log.Println("error: ", err)
	}

	//log.Println(db)
	//创建表
	db.AutoMigrate(&Post{}, &User{})

	//db.CreateTable(&Post{}) //创建posts数据表 如果表已经存在就不创建

	//db.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&Post{}) //创建posts表时指存在引擎

	//db.DropTable(&Post{}, "my_user") //删除posts和users表数据表

	//db.DropTableIfExists(&Post{}, "my_user") //删除前会判断posts和users表是否存在

	//先判断users表是否存在，再删除users表
	if db.HasTable("users") {
		db.DropTable("users")
	}

	//删除数据表字段
	//db.Model(&Post{}).DropColumn("name") //ALTER TABLE `my_posts` DROP COLUMN `name`

	//修改字段数据类型
	//db.Model(&Post{}).ModifyColumn("name", "varchar(500)") //如果字段不存在就添加字段

	//给posts表的title字段添加索引
	//db.Model(&Post{}).AddIndex("idx_title", "title") //CREATE INDEX idx_title ON `my_posts`(`title`)

	//db.Model(&Post{}).AddUniqueIndex("idx_uid", "uid") //添加唯一主键  CREATE UNIQUE INDEX idx_uid ON `my_posts`(`uid`)

	//支持以字段的前几位长度索引
	//db.Model(&Post{}).AddUniqueIndex("idx_name", "name(5)") //添加唯一主键  CREATE UNIQUE INDEX idx_uid ON `my_posts`(`uid`)

	//支持复合索引创建
	// CREATE INDEX idx_title_name ON `my_posts`(title(5),name(5))
	db.Model(&Post{}).AddIndex("idx_title_name", "title(5),name(5)")

	//删除索引 根据索引名称进行删除
	//db.Model(&Post{}).RemoveIndex("idx_title")
}
