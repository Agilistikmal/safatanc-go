package model

import (
	"time"
)

type Project struct {
	Id           int       `json:"id,omitempty" validate:"required,number" gorm:"primaryKey,autoIncrement"`
	Title        string    `json:"title,omitempty" validate:"required,min=3"`
	Description  string    `json:"description,omitempty"`
	TechStack    string    `json:"tech_stack,omitempty"`
	Contributors string    `json:"contributors,omitempty" validate:"required"`
	Image        string    `json:"image,omitempty"`
	PreviewUrl   string    `json:"preview_url,omitempty" validate:"url"`
	CreatedAt    time.Time `json:"created_at" validate:"required" gorm:"autoCreateTime"`
	UpdatedAt    time.Time `json:"updated_at" validate:"required" gorm:"autoUpdateTime"`
}

type ProjectRequest struct {
	Title        string   `json:"title,omitempty" validate:"required,min=3"`
	Description  string   `json:"description,omitempty"`
	TechStack    []string `json:"tech_stack,omitempty"`
	Contributors []string `json:"contributors,omitempty" validate:"required"`
	Image        string   `json:"image,omitempty"`
	PreviewUrl   string   `json:"preview_url,omitempty" validate:"url"`
}

type ProjectResponse struct {
	Id           int       `json:"id,omitempty"`
	Title        string    `json:"title,omitempty"`
	Description  string    `json:"description,omitempty"`
	TechStack    []string  `json:"tech_stack,omitempty"`
	Contributors []string  `json:"contributors,omitempty"`
	Image        string    `json:"image,omitempty"`
	PreviewUrl   string    `json:"preview_url,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}
