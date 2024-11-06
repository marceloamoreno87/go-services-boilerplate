package entity

import (
	"time"
)

type Category struct {
	ID        int
	UserID    *int
	Name      string
	Icon      string
	Color     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewCategory(
	userID *int,
	name string,
	icon string,
	color string,
) *Category {
	return &Category{
		UserID:    userID,
		Name:      name,
		Icon:      icon,
		Color:     color,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (c *Category) GetID() int {
	return c.ID
}

func (c *Category) GetUserID() *int {
	return c.UserID
}

func (c *Category) GetName() string {
	return c.Name
}

func (c *Category) GetIcon() string {
	return c.Icon
}

func (c *Category) GetColor() string {
	return c.Color
}

func (c *Category) GetCreatedAt() time.Time {
	return c.CreatedAt
}

func (c *Category) GetUpdatedAt() time.Time {
	return c.UpdatedAt
}

func (c *Category) SetID(ID int) *Category {
	c.ID = ID
	return c
}

func (c *Category) SetUserID(userID *int) *Category {
	c.UserID = userID
	return c
}

func (c *Category) SetName(name string) *Category {
	c.Name = name
	return c
}

func (c *Category) SetIcon(icon string) *Category {
	c.Icon = icon
	return c
}

func (c *Category) SetColor(color string) *Category {
	c.Color = color
	return c
}

func (c *Category) SetCreatedAt(createdAt time.Time) *Category {
	c.CreatedAt = createdAt
	return c
}

func (c *Category) SetUpdatedAt(updatedAt time.Time) *Category {
	c.UpdatedAt = updatedAt
	return c
}
