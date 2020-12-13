package main

import (
	"context"
	"io/ioutil"
	"log"

	pb "github.com/e-conomic/hiring-assignments/machinelearningteam/summary-statistics-service/proto"

	"google.golang.org/grpc"
)

const (
	host = "34.68.0.193:50051"
)

func main() {
	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewDocumentSummarizerClient(conn)

	ctx := context.Background()
	resp, err := client.SummarizeDocument(ctx, &pb.SummarizeDocumentRequest{
		Document: &pb.Document{
			Source: &pb.DocumentSource{
				HttpUri: "https://raw.githubusercontent.com/e-conomic/hiring-assignments/master/machinelearningteam/summary-statistics-service/test.csv",
			},
		},
	})

	ioutil.WriteFile("out_url.csv", resp.GetContent(), 0644)
}
