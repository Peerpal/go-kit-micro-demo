package watermark

import (
    "context"
    "github.com/go-kit/kit/log"
    "github.com/lithammer/shortuuid/v3"
    "micro_demo/internal"
    "micro_demo/prisma"
    "net/http"
    "os"
)

type watermarkService struct {
	client *prisma.PrismaClient
}

func NewService(client *prisma.PrismaClient) Service {
	return &watermarkService{
		client: client,
	}
}

func (receiver *watermarkService) Get(ctx context.Context) ([]prisma.DocumentModel, error) {
	documents, err := receiver.client.Document.FindMany().Exec(ctx)
    if err != nil {
		logger.Log("error Fetching Documents")
    }
	return documents, nil
}

func (receiver *watermarkService) Status(ctx context.Context, ticketId string) (internal.Status, error) {
	return internal.InProgress, nil
}

func (receiver *watermarkService) Watermark(ctx context.Context, ticketID, mark string) (int, error) {
	// update the database entry with watermark field as non empty
	// first check if the watermark status is not already in InProgress, Started or Finished state
	// If yes, then return invalid request
	// return error if no item found using the ticketID
	return http.StatusOK, nil
}

func (receiver *watermarkService) AddDocument(ctx context.Context, doc interface{}) (string, error) {
	// add the document entry in the database by calling the database service
	// return error if the doc is invalid and/or the database invalid entry error
	newTicket := shortuuid.New()

	return newTicket, nil
}

func (receiver *watermarkService) ServiceStatus(ctx context.Context) (int, error) {
	logger.Log("Checking the service health")
	return http.StatusOK, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}