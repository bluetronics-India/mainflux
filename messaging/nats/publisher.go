// Copyright (c) Mainflux
// SPDX-License-Identifier: Apache-2.0

package nats

import (
	"github.com/gogo/protobuf/proto"
	"github.com/mainflux/mainflux/messaging"
	broker "github.com/nats-io/nats.go"
)

var _ messaging.Publisher = (*Publisher)(nil)

type Publisher struct {
	conn *broker.Conn
}

// NewPublisher returns NATS message Publisher.
func NewPublisher(conn *broker.Conn) messaging.Publisher {
	return Publisher{
		conn: conn,
	}
}

func (pub Publisher) Publish(topic string, msg messaging.Message) error {
	data, err := proto.Marshal(&msg)
	if err != nil {
		return err
	}

	if err := pub.conn.Publish(topic, data); err != nil {
		return err
	}

	return nil
}