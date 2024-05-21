// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSysLog = "sys_logs"

// SysLog mapped from table <sys_logs>
type SysLog struct {
	ID           int64     `gorm:"column:id;type:bigint;primaryKey;autoIncrement:true;comment:主键id" json:"id"`              // 主键id
	Title        string    `gorm:"column:title;type:varchar(128);not null;comment:操作的模块" json:"title"`                      // 操作的模块
	BusinessType int32     `gorm:"column:business_type;type:tinyint;not null;comment:0其它 1新增 2修改 3删除" json:"business_type"` // 0其它 1新增 2修改 3删除
	URL          string    `gorm:"column:url;type:varchar(128);comment:操作url" json:"url"`                                   // 操作url
	Operation    string    `gorm:"column:operation;type:varchar(200);comment:操作路径" json:"operation"`                        // 操作路径
	Method       string    `gorm:"column:method;type:varchar(20);not null;comment:请求方法" json:"method"`                      // 请求方法
	UserName     string    `gorm:"column:user_name;type:varchar(255);not null;comment:操作人员" json:"user_name"`               // 操作人员
	UserID       int64     `gorm:"column:user_id;type:bigint;not null;comment:用户id" json:"user_id"`                         // 用户id
	IP           string    `gorm:"column:ip;type:varchar(16);not null;comment:操作IP" json:"ip"`                              // 操作IP
	Latency      int32     `gorm:"column:latency;type:int;not null;comment:延迟" json:"latency"`                              // 延迟
	Resp         string    `gorm:"column:resp;type:varchar(255);not null;comment:请求参数" json:"resp"`                         // 请求参数
	Status       int32     `gorm:"column:status;type:tinyint;not null;default:1;comment:1=正常 2=异常" json:"status"`           // 1=正常 2=异常
	ErrorMessage string    `gorm:"column:error_message;type:varchar(191);not null;comment:错误信息" json:"error_message"`       // 错误信息
	CreatedAt    time.Time `gorm:"column:created_at;type:datetime;comment:创建时间" json:"created_at"`                          // 创建时间
}

// TableName SysLog's table name
func (*SysLog) TableName() string {
	return TableNameSysLog
}