package endpoints

import (
	"micro_demo/internal"
	"micro_demo/prisma"
)

type GetRequest struct {
}

type GetResponse struct {
	Documents []prisma.DocumentModel `json:"documents"`
	Err       string                 `json:"err,omitempty"`
}

type StatusRequest struct {
	TicketId string `json:"ticketId"`
}

type StatusResponse struct {
	Status internal.Status `json:"status"`
	Err    string          `json:"err,omitempty"`
}

type WatermarkRequest struct {
	TicketId string `json:"ticketId"`
	Mark     string `json:"mark"`
}

type WatermarkResponse struct {
	Code int    `json:"code"`
	Err  string `json:"err"`
}

type AddDocumentRequest struct {
	Document *prisma.DocumentModel `json:"document"`
}

type AddDocumentResponse struct {
	TicketId string `json:"ticketId"`
	Err      string `json:"err,omitempty"`
}

type ServiceStatusRequest struct{}

type ServiceStatusResponse struct {
	Code int    `json:"status"`
	Err  string `json:"err,omitempty"`
}
