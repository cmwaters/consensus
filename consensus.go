package consensus

import "context"

type Execution interface {
	Execute(tx []byte) uint64
	Verify(tx []byte) error
	Query(query []byte) []byte
	State() uint64
}

type Consensus interface {
	Agree(
		ctx context.Context, 
		height uint64, 
		proposedTxs []byte,
		group Group,
		valid func([]byte) bool,
	) ([]byte, error)
}

type Dissemination interface {
	Broadcast(tx []byte) error
	Get(ctx context.Context, txKey [32]byte) ([]byte, error)
	Scan(ctx context.Context, n int) [][]byte
}

type Group interface {
	Proposer(index int) Member
	Member(index int) Member
}

type Signer interface {
	ID() string
	Sign(nonce []int, msg []byte) ([]byte, error)
}

type Member interface {
	ID() string
	Weight() float64
	Verify(msg, sig []byte) bool
}
