package ipfs

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/json"
	shell "github.com/ipfs/go-ipfs-api"
	"plasma/common"
	"plasma/evm"
)

const DaIpfs = "ipfs"

type MappingCID struct {
	Path string `json:"path"`
}

type Store struct {
	Shel      *shell.Shell
	DaId      [32]byte
	submitter *evm.Submitter
}

func NewIpfsStore(cfg Config, daId [32]byte, submitter *evm.Submitter) (*Store, error) {
	sh := shell.NewShell(cfg.Url)

	return &Store{
		Shel:      sh,
		DaId:      daId,
		submitter: submitter,
	}, nil
}

func (s *Store) Get(_ context.Context, key []byte) ([]byte, error) {
	// get path from plasma hub contract
	dataRead, err := s.submitter.GetSubmitter(s.submitter.Transactor.Address(), sha256.Sum256(key))
	if err != nil {
		return nil, err
	}

	if len(dataRead) == 0 {
		return nil, common.ErrDataNotFound
	}
	data := dataRead[0]

	var dataMap MappingCID
	if err := json.Unmarshal(data.Cid, &dataMap); err != nil {
		return nil, err
	}

	dataPath, err := s.Shel.Cat(dataMap.Path)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(dataPath)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func (s *Store) Put(_ context.Context, key []byte, value []byte) error {
	cid, err := s.Shel.Add(bytes.NewReader(value))
	if err != nil {
		return err
	}
	println(cid)

	dataMap := MappingCID{
		Path: cid,
	}

	// save path to data map
	dataWrite, err := json.Marshal(dataMap)
	if err != nil {
		return err
	}

	_, err = s.submitter.SubmitData(sha256.Sum256(key), s.DaId, dataWrite)
	if err != nil {
		return err
	}
	return nil
}
