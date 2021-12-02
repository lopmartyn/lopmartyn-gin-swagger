package model

import "context"

var ctx context.Context

func init() {
	ctx = context.Background()
}
