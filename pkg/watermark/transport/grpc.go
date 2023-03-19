package transport

import (
	"context"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	"micro_demo/api/v1/watermark"
	"micro_demo/internal"
	"micro_demo/pkg/watermark/endpoints"
	"micro_demo/prisma"
)

type grpcServer struct {
	get           grpctransport.Handler
	status        grpctransport.Handler
	addDocument   grpctransport.Handler
	watermark     grpctransport.Handler
	serviceStatus grpctransport.Handler
	watermark.UnimplementedWatermarkServer
}



func NewGrpcServer(ep endpoints.Set) watermark.WatermarkServer {
	return &grpcServer{
		get: grpctransport.NewServer(
			ep.GetEndpoint,
			decodeGRPCGetRequest,
			decodeGRPCGetResponse,
		),
		status: grpctransport.NewServer(
			ep.StatusEndpoint,
			decodeGRPCStatusRequest,
			decodeGRPCStatusResponse,
		),
		addDocument: grpctransport.NewServer(
			ep.AddDocumentEndpoint,
			decodeGRPCAddDocumentRequest,
			decodeGRPCAddDocumentResponse,
		),
		watermark: grpctransport.NewServer(
			ep.WatermarkEndpoint,
			decodeGRPCWatermarkRequest,
			decodeGRPCWatermarkResponse,
		),
		serviceStatus: grpctransport.NewServer(
			ep.ServiceStatusEndpoint,
			decodeGRPCServiceStatusRequest,
			decodeGRPCServiceStatusResponse,
		),
	}
}


func (g *grpcServer) Get(ctx context.Context, c *watermark.GetRequest) (*watermark.GetReply, error) {
	_, reply, err := g.get.ServeGRPC(ctx, c)

	if err != nil {
		return nil, err
	}

	return reply.(*watermark.GetReply), err
}

func (g *grpcServer) Watermark(ctx context.Context, c *watermark.WatermarkRequest) (*watermark.WatermarkReply, error) {
	_, reply, err := g.watermark.ServeGRPC(ctx, c)

	if err != nil {
		return nil, err
	}

	return reply.(*watermark.WatermarkReply), nil
}

func (g *grpcServer) Status(ctx context.Context, c *watermark.StatusRequest) (*watermark.StatusReply, error) {
	_, reply, err := g.status.ServeGRPC(ctx, c)

	if err != nil {
		return nil, err
	}

	return reply.(*watermark.StatusReply), nil
}

func (g *grpcServer) AddDocument(ctx context.Context, c *watermark.AddDocumentRequest) (*watermark.AddDocumentReply, error) {
	_, reply, err := g.addDocument.ServeGRPC(ctx, c)

	if err != nil {
		return nil, err
	}

	return reply.(*watermark.AddDocumentReply), nil
}

func (g *grpcServer) ServiceStatus(ctx context.Context, c *watermark.ServiceStatusRequest) (*watermark.ServiceStatusReply, error) {
	_, reply, err := g.serviceStatus.ServeGRPC(ctx, c)

	if err != nil {
		return nil, err
	}

	return reply.(*watermark.ServiceStatusReply), nil
}

func decodeGRPCGetRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {

	return endpoints.GetRequest{}, nil
}
func decodeGRPCStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*watermark.StatusRequest)
	return endpoints.StatusRequest{TicketId: req.TicketID}, nil
}

func decodeGRPCWatermarkRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*watermark.WatermarkRequest)
	return endpoints.WatermarkRequest{TicketId: req.TicketID, Mark: req.Mark}, nil
}

func decodeGRPCAddDocumentRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	req := grpcReq.(*watermark.AddDocumentRequest)
	doc := &prisma.DocumentModel{
		InnerDocument: prisma.InnerDocument{
			Content: req.Document.Content,
			Title:   req.Document.Title,
			Author:  req.Document.Author,
			Topic:   req.Document.Topic,
		},
	}
	return endpoints.AddDocumentRequest{Document: doc}, nil
}

func decodeGRPCServiceStatusRequest(_ context.Context, grpcReq interface{}) (interface{}, error) {
	return endpoints.ServiceStatusRequest{}, nil
}
func decodeGRPCGetResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*watermark.GetReply)

	var docs []prisma.DocumentModel

	for _, d := range reply.Documents {
		doc := prisma.DocumentModel{
			InnerDocument: prisma.InnerDocument{
				Content: d.Content,
				Title:   d.Title,
				Author:  d.Author,
				Topic:   d.Topic,
			},
			RelationsDocument: prisma.RelationsDocument{},
		}

		docs = append(docs, doc)
	}

	return endpoints.GetResponse{Documents: docs, Err: reply.Err}, nil
}

func decodeGRPCStatusResponse(_ context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*watermark.StatusReply)
	return endpoints.StatusResponse{Status: internal.Status(reply.Status), Err: reply.Err}, nil
}

func decodeGRPCWatermarkResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*watermark.WatermarkReply)
	return endpoints.WatermarkResponse{Code: int(reply.Code), Err: reply.Err}, nil
}

func decodeGRPCAddDocumentResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*watermark.AddDocumentReply)
	return endpoints.AddDocumentResponse{TicketId: reply.TicketID, Err: reply.Err}, nil
}

func decodeGRPCServiceStatusResponse(ctx context.Context, grpcReply interface{}) (interface{}, error) {
	reply := grpcReply.(*watermark.ServiceStatusReply)
	return endpoints.ServiceStatusResponse{Code: int(reply.Code), Err: reply.Err}, nil
}
