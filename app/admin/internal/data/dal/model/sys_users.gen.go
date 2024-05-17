// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameSysUser = "sys_users"

// SysUser mapped from table <sys_users>
type SysUser struct {
	ID        int64          `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键id" json:"id"`    // 主键id
	UUID      string         `gorm:"column:uuid;type:varchar(64);not null;comment:用户UUID" json:"uuid"`              // 用户UUID
	Username  string         `gorm:"column:username;type:varchar(64);not null;comment:用户名(登入)" json:"username"`     // 用户名(登入)
	NickName  string         `gorm:"column:nick_name;type:varchar(64);not null;comment:昵称" json:"nick_name"`        // 昵称
	Password  string         `gorm:"column:password;type:varchar(128);not null;comment:密码" json:"password"`         // 密码
	Phone     string         `gorm:"column:phone;type:varchar(16);not null;comment:手机" json:"phone"`                // 手机
	RoleID    int64          `gorm:"column:role_id;type:bigint;not null;comment:角色id" json:"role_id"`               // 角色id
	Salt      string         `gorm:"column:salt;type:varchar(255);not null;comment:盐" json:"salt"`                  // 盐
	Avatar    string         `gorm:"column:avatar;type:varchar(255);not null;comment:头像" json:"avatar"`             // 头像
	Sex       int32          `gorm:"column:sex;type:tinyint;not null;comment:性别 0-未知 1-男 2-女" json:"sex"`           // 性别 0-未知 1-男 2-女
	Email     string         `gorm:"column:email;type:varchar(128);not null;comment:邮箱" json:"email"`               // 邮箱
	DeptID    int64          `gorm:"column:dept_id;type:bigint;not null;comment:部门id" json:"dept_id"`               // 部门id
	PostID    int64          `gorm:"column:post_id;type:bigint;not null;comment:岗位id" json:"post_id"`               // 岗位id
	Remark    string         `gorm:"column:remark;type:varchar(255);not null;comment:备注" json:"remark"`             // 备注
	Status    int32          `gorm:"column:status;type:tinyint;not null;default:1;comment:1=正常 2=异常" json:"status"` // 1=正常 2=异常
	RoleIds   string         `gorm:"column:role_ids;type:varchar(255);not null;comment:多角色" json:"role_ids"`        // 多角色
	PostIds   string         `gorm:"column:post_ids;type:varchar(255);not null;comment:多岗位" json:"post_ids"`        // 多岗位
	CreateBy  string         `gorm:"column:create_by;type:varchar(128);not null;comment:创建人" json:"create_by"`      // 创建人
	UpdateBy  string         `gorm:"column:update_by;type:varchar(128);not null;comment:更新人" json:"update_by"`      // 更新人
	CreatedAt time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`                // 创建时间
	UpdatedAt time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`                // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`                // 删除时间
	Secret    string         `gorm:"column:secret;type:varchar(255);not null;comment:google密钥" json:"secret"`       // google密钥
}

// TableName SysUser's table name
func (*SysUser) TableName() string {
	return TableNameSysUser
}
