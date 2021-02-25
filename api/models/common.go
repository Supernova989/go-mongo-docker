package models

type RequestUpdate struct {
	ModifiedCount int64 `json:"modifiedCount"`
	Result        interface{}
}

type RequestDelete struct {
	DeletedCount int64 `json:"deletedCount"`
}
