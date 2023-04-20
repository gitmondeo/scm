package db

import (
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

	//和UserInfo表建立一对一关系
	UserInfoID int `gorm:"unique"`
	UserInfo   UserInfo
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

	//和UserInfo表建立一对一关系
	UserInfoID int `gorm:"unique"`
	UserInfo   UserInfo
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

type Admin struct {
	Base
	//和UserInfo表建立一对一关系
	UserInfoID int `gorm:"unique"`
	UserInfo   UserInfo
}

//登录用户表，和teacher表、student表、admin表建立 一对一关联关系
type UserInfo struct {
	ID       int `grom:"primaryKey"`
	Account  string
	Password string
}
