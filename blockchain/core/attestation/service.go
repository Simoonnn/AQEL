package attestation

import (
	"context"
	"github.com/raidoNetwork/RDO_v2/blockchain/consensus"
	"github.com/raidoNetwork/RDO_v2/blockchain/consensus/attestation"
	"github.com/raidoNetwork/RDO_v2/blockchain/consensus/staking"
	"github.com/raidoNetwork/RDO_v2/blockchain/core/rdochain"
	"github.com/raidoNetwork/RDO_v2/blockchain/state"
	"github.com/raidoNetwork/RDO_v2/events"
	"github.com/raidoNetwork/RDO_v2/proto/prototype"
	"github.com/raidoNetwork/RDO_v2/shared/params"
	"time"
)

func NewService(ctx context.Context, bc *rdochain.Service, txFeed events.Feed, enableStats bool) (*Service, error) {
	cfg := params.RaidoConfig()

	stakeAmount := cfg.StakeSlotUnit * cfg.RoiPerRdo
	slotTime := time.Duration(cfg.SlotTime) * time.Second

	// create new staking pool
	stakePool, err := staking.NewPool(bc, cfg.ValidatorRegistryLimit, cfg.RewardBase, stakeAmount)
	if err != nil {
		return nil, err
	}

	validatorCfg := attestation.CryspValidatorConfig{
		SlotTime:               slotTime,
		MinFee:                 cfg.MinimalFee,
		StakeUnit:              stakeAmount,
		ValidatorRegistryLimit: cfg.ValidatorRegistryLimit,
		LogStat: enableStats,
	}

	// new block and tx validator
	validator := attestation.NewCryspValidator(bc, stakePool, &validatorCfg)

	// new tx pool
	txPool := NewTxPool(ctx, validator, &PoolConfig{
		MinimalFee: cfg.MinimalFee,
		BlockSize:  cfg.BlockSize,
		TxFeed: txFeed,
	})

	srv := &Service{
		txPool: txPool,
		stakePool: stakePool,
		validator: validator,
		stateEvent: make(chan state.State, 1),
	}

	return srv, nil
}

type Service struct{
	txPool *TxPool
	stakePool consensus.StakePool
	validator consensus.Validator
	stateEvent chan state.State
}

func (s *Service) Start(){
	s.waitNodeSync()

	// listen events for forge error
	go s.lookForgeError()

	// start tx pool work
	go s.txPool.ReadingLoop()
}

func (s *Service) Status() error {
	return nil
}

func (s *Service) Stop() error {
	s.txPool.StopWriting()
	log.Info("Stop tx pool")

	return nil
}


// SendRawTx implements PoolAPI for gRPC gateway
func (s *Service) SendRawTx(tx *prototype.Transaction) error {
	return s.txPool.SendRawTx(tx)
}

// GetFee returns minimal fee needed to insert transaction in the block.
func (s *Service) GetFee() uint64 {
	return s.txPool.GetFee()
}

// GetPendingTransactions returns list of pending transactions.
func (s *Service) GetPendingTransactions() ([]*prototype.Transaction, error) {
	return s.txPool.GetPendingTransactions()
}

func (s *Service) GetServiceStatus() (string, error) {
	return "Unimplemented", nil
}

func (s *Service) StakePool() consensus.StakePool {
	return s.stakePool
}

func (s *Service) TxPool() consensus.TxPool {
	return s.txPool
}

func (s *Service) Validator() consensus.Validator {
	return s.validator
}

func (s *Service) waitNodeSync() {
	for st := range s.stateEvent {
		if st == state.Synced {
			return
		}
	}
}

func (s *Service) lookForgeError() {
	for st := range s.stateEvent {
		if st == state.ForgeFailed {
			s.txPool.catchForgeError()
			return
		}
	}
}