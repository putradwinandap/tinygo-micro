package iface

import (
	"analytics/internal/domain/entity"
)

type URLVisitLogRepository interface {
	Save(log entity.URLVisitLog) error
	FindByShortCode(shortCode string) ([]entity.URLVisitLog, error)
}
