package types

import (
	log "github.com/sirupsen/logrus"
	"heyuanlong/blockchain-step/protocol"
	"google.golang.org/protobuf/proto"
)

type TransactionMgt struct {
	Tx protocol.Tx
}

func (ts TransactionMgt) Bytes() ([]byte,error) {
	b, err := proto.Marshal(&ts.Tx)
	if err != nil {
		log.Error("to bytes fail",err)
		return []byte{}, err
	}
	return b ,nil
}

func (ts TransactionMgt) SetSign(sign []byte)  {
	ts.Tx.Sign = sign
}
