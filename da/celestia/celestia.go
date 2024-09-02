package celestia

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/rollkit/go-da"
	"github.com/rollkit/go-da/proxy"
	"plasma/common"
	"plasma/evm"
	"time"
)

// ErrBlobTooLarge is returned when the blob is too large.
var ErrBlobTooLarge = errors.New("blob too large")

// ErrFailedToSubmit is returned when the blob submission fails.
var ErrFailedToSubmit = errors.New("failed to submit blob")

var ErrBlobNotFound = errors.New("blob: not found")

const DaCelestia = "celestia"

type MappingCommitment struct {
	// key is the key of the blob
	Key []da.ID `json:"key"`
}

type Store struct {
	Client     da.DA
	GetTimeout time.Duration
	Namespace  da.Namespace
	cfg        Config
	daId       [32]byte
	submitter  *evm.Submitter
}

// NewCelestiaStore creates a new CelestiaStore.
func NewCelestiaStore(cfg Config, daId [32]byte, submitter *evm.Submitter) (*Store, error) {
	client, err := proxy.NewClient(cfg.Rpc, cfg.AuthToken)
	if err != nil {
		return nil, err
	}

	ns, err := hex.DecodeString(cfg.Namespace)
	if err != nil {
		return nil, err
	}

	return &Store{
		Client:     client,
		GetTimeout: time.Minute,
		Namespace:  ns,
		cfg:        cfg,
		daId:       daId,
		submitter:  submitter,
	}, nil
}

func (c *Store) Get(ctx context.Context, key []byte) ([]byte, error) {
	// get ids from plasma hub contract
	dataRead, err := c.submitter.GetSubmitter(c.submitter.Transactor.Address(), sha256.Sum256(key), c.daId)
	if err != nil {
		return nil, err
	}
	if len(dataRead) == 0 {
		return nil, common.ErrDataNotFound
	}

	var dataMap MappingCommitment
	if err := json.Unmarshal(dataRead, &dataMap); err != nil {
		return nil, err
	}
	dataBlob, err := c.Client.Get(ctx, dataMap.Key, c.Namespace)
	if err != nil {
		return nil, err
	}
	if len(dataBlob) == 0 {
		return nil, ErrBlobNotFound
	}
	return dataBlob[0], nil
}

func (c *Store) Put(ctx context.Context, key []byte, value []byte) error {
	if uint64(len(value)) > c.cfg.MaxBlobSize {
		return ErrBlobTooLarge
	}

	var blobs [][]byte
	blobs = append(blobs, value)

	ids, err := c.Client.Submit(ctx, blobs, c.cfg.GasPrice, c.Namespace)
	if err != nil {
		return err
	}

	if len(ids) == 0 {
		return ErrFailedToSubmit
	}

	dataMap := MappingCommitment{
		Key: ids,
	}
	dataWrite, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	_, err = c.submitter.SubmitData(sha256.Sum256(key), c.daId, dataWrite)
	if err != nil {
		return err
	}
	return nil
}
