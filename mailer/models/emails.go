package models

import "gorm.io/gorm"

type Emails struct {
	gorm.Model
	Type           string       `json:"type"`
	Sender         string       `json:"sender"`
	Recipient      string       `json:"recipient"`
	Body           string       `json:"body"`
	Subject        string       `json:"subject"`
	Newsletter     string       `json:"newsletter"`
	Status         string       `json:"status"`
	StatusHash     string       `json:"status_hash"`
	EmailServersID int          `json:"used_server"`
	UsedServer     EmailServers `gorm:"foreignKey:EmailServersID"`
	EmailRemoteID  string       `json:"email_remote_id"`
}
