package endpoints

import (
	"context"
	"errors"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
    "micro_demo/internal"
    "micro_demo/pkg/watermark"
	"micro_demo/prisma"
	"os"
)

type Set struct {
	GetEndpoint           endpoint.Endpoint
	AddDocumentEndpoint   endpoint.Endpoint
	StatusEndpoint        endpoint.Endpoint
	ServiceStatusEndpoint endpoint.Endpoint
	WatermarkEndpoint     endpoint.Endpoint
}

func NewEndpointSet(svc watermark.Service) Set {
	return Set{
		GetEndpoint:           MakeGetEndpoint(svc),
		StatusEndpoint:        MakeStatusEndpoint(svc),
		AddDocumentEndpoint:   MakeAddDocumentEndpoint(svc),
		ServiceStatusEndpoint: MakeServiceStatusEndpoint(svc),
		WatermarkEndpoint:     MakeWatermarkEndpoint(svc),
	}
}

func MakeGetEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		//		req := request.(GetRequest)

		docs, err := svc.Get(ctx)

		if err != nil {
			return GetResponse{docs, err.Error()}, nil
		}

		return GetResponse{docs, ""}, nil

	}
}

func MakeStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(StatusRequest)

		status, err := svc.Status(ctx, req.TicketId)

		if err != nil {
			return StatusResponse{Status: status, Err: err.Error()}, nil
		}

		return StatusResponse{Status: status, Err: ""}, nil
	}
}

func MakeAddDocumentEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(AddDocumentRequest)

		document, err := svc.AddDocument(ctx, req)

		if err != nil {
			return AddDocumentResponse{TicketId: document, Err: err.Error()}, nil
		}

		return AddDocumentResponse{TicketId: document, Err: ""}, nil
	}
}

func MakeServiceStatusEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		status, err := svc.ServiceStatus(ctx)

		if err != nil {
			return ServiceStatusResponse{Code: status, Err: err.Error()}, nil
		}

		return ServiceStatusResponse{Code: status, Err: ""}, nil

	}

}

func MakeWatermarkEndpoint(svc watermark.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(WatermarkRequest)

		mark, err := svc.Watermark(ctx, req.Mark, req.TicketId)

		if err != nil {
			return WatermarkResponse{Code: mark, Err: err.Error()}, nil
		}
		return WatermarkResponse{Code: mark, Err: ""}, nil
	}
}

func (s *Set) Get(ctx context.Context) ([]prisma.DocumentModel, error) {
	resp, err := s.GetEndpoint(ctx, GetRequest{})

	if err != nil {
		return []prisma.DocumentModel{}, err
	}

	getResp := resp.(GetResponse)

	if getResp.Err != "" {
		return []prisma.DocumentModel{}, errors.New(getResp.Err)
	}
	return getResp.Documents, nil
}

func (s *Set) ServiceStatus(ctx context.Context) (int, error) {
	resp, err := s.ServiceStatusEndpoint(ctx, ServiceStatusRequest{})

	svcStatusResp := resp.(ServiceStatusResponse)

	if err != nil {

		return svcStatusResp.Code, err
	}

	if svcStatusResp.Err != "" {

		return svcStatusResp.Code, errors.New(svcStatusResp.Err)

	}

	return svcStatusResp.Code, nil
}

func (s *Set) AddDocument(ctx context.Context, doc *prisma.DocumentModel) (string, error)  {
	resp, err := s.AddDocumentEndpoint(ctx, AddDocumentRequest{Document: doc})

	if err != nil {
		return "", err
	}

	adResp := resp.(AddDocumentResponse)
	if adResp.Err != "" {
		return "", errors.New(adResp.Err)
	}

	return adResp.TicketId, nil
}

func (s *Set) Status(ctx context.Context, ticketID string) (internal.Status, error) {
	resp, err := s.StatusEndpoint(ctx, StatusRequest{TicketId: ticketID})
	if err != nil {
		return internal.Failed, err
	}
	stsResp := resp.(StatusResponse)
	if stsResp.Err != "" {
		return internal.Failed, errors.New(stsResp.Err)
	}
	return stsResp.Status, nil
}

func (s *Set) Watermark(ctx context.Context, ticketID, mark string) (int, error) {
	resp, err := s.WatermarkEndpoint(ctx, WatermarkRequest{TicketId: ticketID, Mark: mark})
	wmResp := resp.(WatermarkResponse)
	if err != nil {
		return wmResp.Code, err
	}
	if wmResp.Err != "" {
		return wmResp.Code, errors.New(wmResp.Err)
	}
	return wmResp.Code, nil
}

var logger log.Logger

func init() {
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)
}
