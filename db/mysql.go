package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

//------------------------模型创建start-------------------------------------
//teacher表：教师ID、教师名称、账号、密码、性别、生日、年龄、备注			创建时间、修改时间、删除时间
//class表：班级ID、班级名称、人数、导员    							创建时间、修改时间、删除时间

//student表：学生ID、姓名、年龄、性别、账号、密码、生日、备注、班级ID	    创建时间、修改时间、删除时间
//course表：课程ID、课程名称、学分、周期、教师ID					创建时间、修改时间、删除时间
//关联表 student2course: ID、学生ID、课程ID							创建时间、修改时间、删除时间

/*
教师表  --->   课程表  （一对多）
班级表  --->   学生表  （一对多）
学生表  --->   课程表  （多对多）
*/

type Base struct {
	ID       int        `gorm:"primaryKey"`
	Name     string     `gorm:"type:varchar(32);unique;not null"`
	CreateAt *time.Time `gorm:"autoCreateTime"`
	UpdateAt *time.Time `gorm:"autoCreateTime"`
	DeleteAt *time.Time `gorm:"autoCreateTime"`
}

//教师表
type Teacher struct {
	Base
	Tno    int        //账号
	Pwd    string     `gorm:"type:varchar(100);not null"` //密码
	Tel    string     `gorm:"type:char(11);"`             //电话
	Birth  *time.Time //生日
	Remark string     `gorm:"type:varchar(255);"` //备注
	Gender string     `gorm:"type:varchar(1);"`   //性别
	Age    int
}

//班级表
type Class struct {
	Base
	Num     int     //人数
	TutorID int     //教师ID
	Tutor   Teacher //教师属于Teacher类型
}

//学生表
type Student struct {
	Base
	Gender string     `gorm:"type:varchar(1);"` //性别
	Age    int        //年龄
	Sno    int        //账号
	Pwd    string     `gorm:"type:varchar(100);not null"` //密码
	Tel    string     `gorm:"type:char(11);"`             //电话
	Birth  *time.Time //生日
	Remark string     `gorm:"type:varchar(255);"` //备注

	//多对一
	ClassID int   //班级ID
	Class   Class `gorm:"foreignKey:ClassID"` //班级ID是外键

	//多对多
	Course []Course `gorm:"many2many:student2course;constraint:OnDelete:CASCADE;"`
}

//课程表
type Course struct {
	Base
	Period int //学时
	Credit int //学分

	//多对一
	TeacherID int
	Teacher   Teacher
}

//------------------------模型创建end-------------------------------------

//-------------------------连接数据库-------------------------------------
var db *gorm.DB

func InitDB() *gorm.DB {

	dsn := "root:root@tcp(10.117.54.37:3306)/godb?charset=utf8mb4&parseTime=True&loc=Local"
	//配置日志对象
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Info,
		},
	)
	db, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	//迁移模型类，将模型类转换成SQL
	db.AutoMigrate(
		&Teacher{},
		&Class{},
		&Student{},
		&Course{},
	)
	return db

}
