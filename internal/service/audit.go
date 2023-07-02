package service

import (
	"context"
	audit "github.com/Asqar95/crud-audit-log/pkg/domain"
	"strconv"
)

type Repository interface {
	Insert(ctx context.Context, item audit.LogItem) error
}

type Audit struct {
	repo Repository
}

func NewAudit(repo Repository) *Audit {
	return &Audit{
		repo: repo,
	}
}

func (s *Audit) Insert(ctx context.Context, req *audit.LogRequest) error {
	item := audit.LogItem{
		Action:    req.GetAction().String(),
		Entity:    req.GetEntity().String(),
		EntityID:  strconv.FormatInt(req.GetEntityId(), 10),
		Timestamp: req.GetTimestamp().AsTime(),
	}

	return s.repo.Insert(ctx, item)
}
