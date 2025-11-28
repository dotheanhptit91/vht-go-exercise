package pubsub

import (
	"context"
	"encoding/json"
	"flag"

	"github.com/nats-io/nats.go"
	"github.com/pkg/errors"
	sctx "github.com/viettranx/service-context"
)

type natsc struct {
	nc     *nats.Conn
	uri    string
	id     string
	logger sctx.Logger
}

func NewNatsComp(id string) *natsc {
	return &natsc{id: id}
}

func (n *natsc) ID() string {
	return n.id
}

func (n *natsc) InitFlags() {
	flag.StringVar(&n.uri, "nats-uri", nats.DefaultURL, "Nats URI")
}

func (n *natsc) Activate(sctx sctx.ServiceContext) error {
	n.logger = sctx.Logger("nats")
	nc, err := nats.Connect(n.uri)

	if err != nil {
		return errors.WithStack(err)
	}

	n.nc = nc
	n.logger.Debugln("Nats connected to", n.uri)

	return nil
}

func (n *natsc) Stop() error {
	if n.nc != nil {
		n.nc.Drain()
		n.nc.Close()
	}

	return nil
}

func (n *natsc) Publish(ctx context.Context, channel Topic, data *Message) error {
	jsonData, err := json.Marshal(data.Data())
	// n.logger.Debugln("Publishing message to channel", string(channel), string(jsonData))

	if err != nil {
		return errors.WithStack(err)
	}

	return n.nc.Publish(string(channel), jsonData)
}

func (n *natsc) Subscribe(ctx context.Context, channel Topic) (ch chan *Message, close func()) {
	ch = make(chan *Message)

	sub, err := n.nc.Subscribe(string(channel), func(m *nats.Msg) {
		var mapData map[string]interface{}

		err := json.Unmarshal(m.Data, &mapData)
		n.logger.Debugln("Received message from channel", string(channel), string(m.Data), mapData)
		if err != nil {
			n.logger.Errorln("Nats connected to", n.uri)
			return
		}

		appMessage := NewMessage(mapData).WithChannel(channel)

		ch <- appMessage
	})

	if err != nil {
		n.logger.Errorln("Error subscribing to channel", err)
		return nil, nil
	}

	return ch, func() {
		sub.Unsubscribe()
	}
}