package handlers

import (
	"errors"
	"time"

	"github.com/p2064/creator/proto"
	"github.com/p2064/pkg/db"
	"github.com/p2064/pkg/logs"
	kafka "github.com/segmentio/kafka-go"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) CreateEvent(ctx context.Context, in *proto.CreateEventRequest) (
	*proto.CreateEventResponse,
	error,
) {
	logs.InfoLogger.Printf("Receive message body from client: %s", in.String())
	data := db.Event{
		Place:      in.Place,
		EventTime:  in.Time,
		MaxPlayers: in.MaxPlayers,
	}
	res := db.DB.Create(&data)
	if res.Error != nil {
		return nil, errors.New("Event creation unsuccessful")
	}

	topic := "notify"
	partition := 0
	conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", topic, partition)
	if err != nil {
		logs.ErrorLogger.Fatal("failed to dial leader:", err)
	}
	conn.SetWriteDeadline(time.Now().Add(10 * time.Second))
	_, err = conn.WriteMessages(
		kafka.Message{Value: []byte(in.String())},
	)
	if err != nil {
		logs.ErrorLogger.Fatal("failed to write messages:", err)
	}
	if err := conn.Close(); err != nil {
		logs.ErrorLogger.Fatal("failed to close writer:", err)
	}
	return &proto.CreateEventResponse{Status: 200, Error: "No error", Id: data.ID}, nil
}
