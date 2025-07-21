// Copyright 2025 Buzz Space Labs
//
// Adapted from the Aura project © 2023 Aura Contributors
// Licensed under the Apache License, Version 2.0
// See LICENSE file in the project root for full license text.

package types

import (
	"encoding/json"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type AccountMsg struct {
	PreExecuteTx   *PreExecuteTx   `json:"pre_execute,omitempty"`
	AfterExecuteTx *AfterExecuteTx `json:"after_execute,omitempty"`
	RecoverTx      *RecoverTx      `json:"recover,omitempty"`
}

type RecoverTx struct {
	Caller      string `json:"caller"`
	PubKey      []byte `json:"pub_key"`
	Credentials []byte `json:"credentials"`
}

type TxInfor struct {
	Gas     uint64    `json:"gas"`
	Fee     sdk.Coins `json:"fee"`
	Payer   string    `json:"payer"`
	Granter string    `json:"granter"`
}

type PreExecuteTx struct {
	Msgs    []MsgData `json:"msgs"`
	TxInfor TxInfor   `json:"tx_infor"`
}

type AfterExecuteTx struct {
}

type MsgData struct {
	TypeURL string `json:"type_url"`
	Value   string `json:"value"`
}

func ParseMessagesString(msgs []sdk.Msg) ([]MsgData, error) {
	msgsStr := make([]MsgData, 0)

	for _, msg := range msgs {
		msgData, err := json.Marshal(msg)
		if err != nil {
			return nil, err
		}

		data := MsgData{
			TypeURL: sdk.MsgTypeURL(msg),
			Value:   string(msgData),
		}

		msgsStr = append(msgsStr, data)
	}
	return msgsStr, nil
}
