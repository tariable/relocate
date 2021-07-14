package repository

import "context"

type Announcements interface {
	Create(ctx context.Context)
}

type Repositories struct {
	Announcements Announcements
}
