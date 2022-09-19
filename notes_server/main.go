package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/vaino-online/paper/notes"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "The notes server port number")
)

// Implement the notes service (notes.NotesServer interface)
type notesServer struct {
	notes.UnimplementedNotesServer
}

func (s *notesServer) Save(ctx context.Context, n *notes.Note) (*notes.NoteSaveReply, error) {
	log.Printf("Received a note to save: %v", n.Title)
	err := notes.SaveToDisk(n, "testdata")

	if err != nil {
		return &notes.NoteSaveReply{Saved: false}, err
	}

	return &notes.NoteSaveReply{Saved: true}, nil
}

func (s *notesServer) Load(ctx context.Context, search *notes.NoteSearch) (*notes.Note, error) {
	log.Printf("Searching for notes with keyword: %v", search.Keyword)
	note, err := notes.LoadFromDisk(search.Keyword, "testdata")

	if err != nil {
		return &notes.Note{}, err
	}

	return note, nil
}

func main() {
	// Parse the command line arguments and start a new TCP listener
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to start notes server: %v", err)
	}
	// Instantiate the gRPC server
	server := grpc.NewServer()

	// Register the server implementation
	notes.RegisterNotesServer(server, &notesServer{})

	log.Printf("Started notes server at %v", listener.Addr())

	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start notes server: %v", err)
	}
}
