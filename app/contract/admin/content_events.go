package admin

import "mime/multipart"

type CreateContentBlocks struct {
	ParentID   int            `form:"parentID" validate:"required,gte=0"`
	ParentType string         `form:"parentType" validate:"required"`
	Type       string         `form:"type" validate:"required"`
	Index      int            `form:"index" validate:"required"`
	ValueStr   string         `form:"valueStr"`
	ValueFile  multipart.File `form:"valueFile"`
}
