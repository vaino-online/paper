package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/vaino-online/paper/notes"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	addr = flag.String("addr", "localhost:50051", "The notes server address")
)

func main() {
	flag.Parse()

	// Set up the connection to the gRPC server
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to the notes server: %v", err)
	}
	defer conn.Close()

	// Create a new notes client
	client := notes.NewNotesClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Save command definition
	save := flag.NewFlagSet("save", flag.ExitOnError)
	saveTitle := save.String("title", "", "The title of your note")
	saveBody := save.String("content", "", "The contents of your note")

	// Load command definition
	load := flag.NewFlagSet("load", flag.ExitOnError)
	loadKeyword := load.String("keyword", "", "The keyword to search for")

	// If we try to run without a subcommand
	if len(os.Args) < 2 {
		fmt.Println("Error: Expected 'save' or 'load' subcommands")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "save":
		save.Parse(os.Args[2:])
		_, err := client.Save(ctx, &notes.Note{
			Title: *saveTitle,
			Body:  []byte(*saveBody),
		})

		if err != nil {
			log.Fatalf("Failed to save note: %v", err)
		}

		fmt.Printf("Note created successfully: %v\n", *saveTitle)
	case "load":
		load.Parse(os.Args[2:])
		note, err := client.Load(ctx, &notes.NoteSearch{
			Keyword: *loadKeyword,
		})

		if err != nil {
			log.Fatalf("Failed to load note: %v", err)
		}

		fmt.Printf("%vn", note)
	default:
		fmt.Println("Error: Expected 'save' or 'load' subcommands")
		os.Exit(1)
	}
}
