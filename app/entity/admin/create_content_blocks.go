package entity

import "mime/multipart"

type CreateContentBlocksEntity struct {
	ParentID   int
	ParentType string
	Type       string
	Index      int
	ValueStr   string
	ValueFile  *multipart.FileHeader
}
