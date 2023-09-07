// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"gorm.io/gorm"
)

const TableNameLogOper = "log_opers"

// LogOper mapped from table <log_opers>
type LogOper struct {
	ID           int64          `gorm:"column:id;type:bigint(20);primaryKey;autoIncrement:true;comment:主键id" json:"id"`
	Title        string         `gorm:"column:title;type:varchar(128);not null;comment:操作的模块" json:"title"`
	BusinessType int32          `gorm:"column:business_type;type:tinyint(2);not null;comment:0其它 1新增 2修改 3删除" json:"business_type"`
	URL          string         `gorm:"column:url;type:varchar(255);not null;comment:操作url" json:"url"`
	Method       string         `gorm:"column:method;type:varchar(20);not null;comment:请求方法" json:"method"`
	UserName     string         `gorm:"column:user_name;type:varchar(255);not null;comment:操作人员" json:"user_name"`
	UserID       int64          `gorm:"column:user_id;type:bigint(20);not null;comment:用户id" json:"user_id"`
	IP           string         `gorm:"column:ip;type:varchar(16);not null;comment:操作IP" json:"ip"`
	Agent        string         `gorm:"column:agent;type:varchar(16);not null;comment:代理" json:"agent"`
	Latency      int32          `gorm:"column:latency;type:int(11);not null;comment:延迟" json:"latency"`
	Resp         string         `gorm:"column:resp;type:varchar(255);not null;comment:请求参数" json:"resp"`
	Status       int32          `gorm:"column:status;type:tinyint(2);not null;default:1;comment:1=正常 2=异常" json:"status"`
	ErrorMessage string         `gorm:"column:error_message;type:varchar(191);not null;comment:错误信息" json:"error_message"`
	CreatedAt    time.Time      `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"column:updated_at;type:datetime;comment:更新时间" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"column:deleted_at;type:datetime;comment:删除时间" json:"deleted_at"`
}

// TableName LogOper's table name
func (*LogOper) TableName() string {
	return TableNameLogOper
}