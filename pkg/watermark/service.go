package watermark

import (
    "context"
	"micro_demo/internal"
	"micro_demo/prisma"
)

type Service interface {
	Get(ctx context.Context) ([]prisma.DocumentModel, error)
	Status(ctx context.Context, ticketId string) (internal.Status, error)
	Watermark(ctx context.Context, ticketID, mark string) (int, error)
	AddDocument(ctx context.Context, doc interface{}) (string, error)
	ServiceStatus(ctx context.Context) (int, error)
}
