// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMessage = "messages"

// Message mapped from table <messages>
type Message struct {
	MessageID  int64  `gorm:"column:message_id;primaryKey" json:"message_id"` // 唯一消息标识
	Type       int32  `gorm:"column:type;not null" json:"type"`               // 消息类型：1为文字，2为图片，3为视频，4为文件
	Content    string `gorm:"column:content;not null" json:"content"`         // 消息体
	SenderID   int64  `gorm:"column:sender_id;not null" json:"sender_id"`     // 发送者
	ReceiverID int64  `gorm:"column:receiver_id;not null" json:"receiver_id"` // 接受者
	SendAt     string `gorm:"column:send_at;not null" json:"send_at"`         // 发送时间
}

// TableName Message's table name
func (*Message) TableName() string {
	return TableNameMessage
}
