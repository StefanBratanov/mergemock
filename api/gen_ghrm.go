// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package api

import (
	"encoding/json"
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

var _ = (*getHeaderResponseMessageMarshalling)(nil)

// MarshalJSON marshals as JSON.
func (g GetHeaderResponseMessage) MarshalJSON() ([]byte, error) {
	type GetHeaderResponseMessage struct {
		Header ExecutionPayloadHeaderV1 `json:"header"`
		Value  *hexutil.Big             `json:"value"`
	}
	var enc GetHeaderResponseMessage
	enc.Header = g.Header
	enc.Value = (*hexutil.Big)(g.Value)
	return json.Marshal(&enc)
}

// UnmarshalJSON unmarshals from JSON.
func (g *GetHeaderResponseMessage) UnmarshalJSON(input []byte) error {
	type GetHeaderResponseMessage struct {
		Header *ExecutionPayloadHeaderV1 `json:"header"`
		Value  *hexutil.Big              `json:"value"`
	}
	var dec GetHeaderResponseMessage
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Header != nil {
		g.Header = *dec.Header
	}
	if dec.Value != nil {
		g.Value = (*big.Int)(dec.Value)
	}
	return nil
}
