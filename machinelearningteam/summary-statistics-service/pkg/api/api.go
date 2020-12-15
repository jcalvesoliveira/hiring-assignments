package api

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"google.golang.org/grpc"

	api "github.com/e-conomic/hiring-assignments/machinelearningteam/summary-statistics-service/proto"
)

const (
	host = "35.239.102.102:50050"
)

// Server is a server implementing the proto API
type Server struct {
	api.UnimplementedDocumentSummarizerServer
}

// SummarizeDocument send the document provided in the request
// to calculate in the statistics-processing micro-service
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

	// If no document bytes is provided, look for url to download
	docContent := req.Document.GetContent()
	if docContent == nil {
		httpURI := req.Document.GetSource().GetHttpUri()
		resp, err := http.Get(httpURI)
		if err != nil {
			log.Fatalf("not able to get file from url: %v", err)
		}
		defer resp.Body.Close()
		docContent, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("not able to convert request body into bytes: %v", err)
		}
	}

	resp, err := client.ProcessDocument(context.Background(), &api.ProcessDocumentRequest{
		Content: docContent,
	},
	)

	return &api.SummarizeDocumentReply{
		Content: resp.Summary,
	}, nil
}
