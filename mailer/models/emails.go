package models

import "gorm.io/gorm"

type Emails struct {
	ID             uint         `json:"id" gorm:"primary_key"`
	Type           string       `json:"type"`
	Sender         string       `json:"sender"`
	Recipient      string       `json:"recipient"`
	Body           string       `json:"body"`
	Subject        string       `json:"subject"`
	Newsletter     string       `json:"newsletter"`
	Status         string       `json:"status"`
	StatusHash     string       `json:"status_hash"`
	EmailServersID int          `json:"used_server"`
	UsedServer     EmailServers `gorm:"foreignKey:EmailServersID,constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	EmailRemoteID  string       `json:"email_remote_id"`
	gorm.Model
}
