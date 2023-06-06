package main

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/kitex_gen/rpc"
	"github.com/TikTokTechImmersion/assignment_demo_2023/rpc-server/pkg/models"
)

// IMServiceImpl implements the last service interface defined in the IDL.
type IMServiceImpl struct{}

type Message struct {
	Chat     string
	Text     string
	Sender   string
	SendTime string
}

func (s *IMServiceImpl) Send(ctx context.Context, req *rpc.SendRequest) (*rpc.SendResponse, error) {
	SendMessage := &models.Message{}
	SendMessage.Chat = req.Message.Chat
	SendMessage.Text = req.Message.Text
	SendMessage.Sender = req.Message.Sender
	SendMessage.SendTime = uint64(time.Now().UnixMicro())

	_, err := SendMessage.CreateMessage()
	resp := rpc.NewSendResponse()

	if err != nil {
		resp.Code = 500
		return resp, nil
	} else {
		resp := rpc.NewSendResponse()
		resp.Code = 0
		return resp, nil
	}
}

func (s *IMServiceImpl) Pull(ctx context.Context, req *rpc.PullRequest) (*rpc.PullResponse, error) {
	chat := req.Chat
	cursor := req.Cursor
	limit := req.Limit
	reverse := req.Reverse

	resp := rpc.NewPullResponse()

	messages, db := models.GetMessages(chat, strconv.FormatInt(cursor, 10), *reverse, int(limit))
	if db.Error != nil {
		resp.Code = 500
		return resp, nil
	}

	limitPlusOne, db := models.GetMessages(chat, strconv.FormatInt(cursor, 10), *reverse, int(limit)+1)
	if db.Error != nil {
		resp.Code = 500
		return resp, nil
	}

	var respMessages []*rpc.Message
	for _, message := range messages {
		respMessage := rpc.NewMessage()
		respMessage.Chat = message.Chat
		respMessage.Text = message.Text
		respMessage.Sender = message.Sender
		respMessage.SendTime = int64(message.SendTime)
		respMessages = append(respMessages, respMessage)
	}

	var nextCursor int64 = 0
	var hasMore bool = false
	if len(limitPlusOne) > len(messages) {
		hasMore = true
		nextCursor = int64(limitPlusOne[len(limitPlusOne)-1].SendTime)
	}

	resp.Code = 0
	resp.Messages = respMessages
	resp.NextCursor = &nextCursor
	resp.HasMore = &hasMore
	return resp, nil
}

func areYouLucky() (int32, string) {
	if rand.Int31n(2) == 1 {
		return 0, "success"
	} else {
		return 500, "oops"
	}
}
