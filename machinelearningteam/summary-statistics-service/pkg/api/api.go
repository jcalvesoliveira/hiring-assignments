package api

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	api "github.com/e-conomic/hiring-assignments/machinelearningteam/summary-statistics-service/proto"
)

const (
	host = "localhost:50050"
)

// Server is a server implementing the proto API
type Server struct {
	api.UnimplementedDocumentSummarizerServer
}

// SummarizeDocument echoes the document provided in the request
func (s *Server) SummarizeDocument(
	ctx context.Context,
	req *api.SummarizeDocumentRequest,
) (*api.SummarizeDocumentReply, error) {
	// Echo
	fmt.Println("Received document...")

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := api.NewStatisticsProcesserClient(conn)

	doc_content := req.Document.GetContent()
	resp, err := client.ProcessDocument(context.Background(), &api.ProcessDocumentRequest{
		Content: doc_content,
	},
	)

	return &api.SummarizeDocumentReply{
		Content: resp.Summary,
	}, nil
}
