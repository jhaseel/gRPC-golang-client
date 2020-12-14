package main

import (
	"github.com/jhaseel/GRPC/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial("10.186.6.2:8980", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewGeneratePdfClient(conn)

	response, err := c.CreatePdf1(context.Background(), &api.CreatePdf1Request{Parametro: &api.CortePDFAgentes{
		Host: "",
		DireccionPlataforma: "",
		CorreoPlataforma: "",
		NumeroPlataforma: "",
		CorteId: "",
		FechaGenercion: "",
		Periodo: "",
	}})

	if err != nil {
		log.Fatalf("Error when calling CreatePdf1: %s", err)
	}
	log.Printf("Response from server: %s", response.Message)

}
