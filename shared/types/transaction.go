package types

import (
	"crypto/ecdsa"
	"github.com/raidoNetwork/RDO_v2/proto/prototype"
	"github.com/raidoNetwork/RDO_v2/shared/crypto"
	"github.com/raidoNetwork/RDO_v2/shared/hasher"
	"github.com/sirupsen/logrus"
	"time"
)

type TxOptions struct {
	Inputs  []*prototype.TxInput
	Outputs []*prototype.TxOutput
	Fee     uint64
	Data    []byte
	Num     uint64
	Type    uint32
}

var log = logrus.WithField("prefix", "types")

// NewTx creates new transaction with given options
func NewTx(opts TxOptions, key *ecdsa.PrivateKey) (*prototype.Transaction, error) {
	tx := new(prototype.Transaction)

	tx.Num = opts.Num
	tx.Type = opts.Type
	tx.Timestamp = uint64(time.Now().UnixNano())
	tx.Fee = opts.Fee
	tx.Data = opts.Data
	tx.Inputs = opts.Inputs
	tx.Outputs = opts.Outputs

	hash, err := hasher.TxHash(tx)
	if err != nil {
		log.Errorf("NewTx: Error generating tx hash. Error: %s", err)
		return nil, err
	}

	tx.Hash = hash[:]

	if key != nil {
		signer := MakeTxSigner("keccak256")

		// signature digest = Keccak256(hash)
		dgst := GetTxDomain(tx.Hash)
		sign, err := signer.Sign(dgst, key)
		if err != nil {
			return nil, err
		}

		tx.Signature = sign
	} else {
		tx.Signature = make([]byte, crypto.SignatureLength)
	}

	return tx, nil
}

// SignTx create transaction signature with given private key
func SignTx(tx *prototype.Transaction, key *ecdsa.PrivateKey) error {
	signer := MakeTxSigner("keccak256")

	// signature digest = Keccak256(hash)
	dgst := GetTxDomain(tx.Hash)
	sign, err := signer.Sign(dgst, key)
	if err != nil {
		return err
	}

	tx.Signature = sign

	return nil
}


func NewTxData(tx *prototype.Transaction) *TransactionData {
	td := TransactionData{
		tx:    tx,
		size:  tx.SizeSSZ(),
		fee:   tx.Fee,
		num:   tx.Num,
		timestamp: tx.Timestamp,
		alias: make([]string, 0),
	}

	return &td
}

type TransactionData struct {
	tx    *prototype.Transaction
	size  int
	fee   uint64
	num   uint64
	alias []string
	timestamp uint64
	status int
}

func (td *TransactionData) GetTx() *prototype.Transaction {
	return td.tx
}

func (td *TransactionData) Size() int {
	return td.size
}

func (td *TransactionData) Fee() uint64 {
	return td.fee
}

func (td *TransactionData) Num() uint64 {
	return td.num
}

func (td *TransactionData) Timestamp() uint64 {
	return td.timestamp
}

func (td *TransactionData) AddAlias(hash string) {
	td.alias = append(td.alias, hash)
}

func (td *TransactionData) GetAlias() []string {
	return td.alias
}

func (td *TransactionData) SetStatus(s int){
	td.status = s
}