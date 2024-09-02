package evm

import (
	"crypto/sha256"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
	"os"
	"plasma/config"
	"testing"
)

func TestKeyStore(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("")
	if err != nil {
		t.Fatal(err)
	}

	//ks := keystore.NewKeyStore("keystore", keystore.StandardScryptN, keystore.StandardScryptP)
	//acc, err := ks.ImportECDSA(privateKey, "passphrase")
	//if err != nil {
	//	t.Fatal(err)
	//}
	id, err := uuid.NewRandom()
	if err != nil {
		panic(fmt.Sprintf("Could not create random uuid: %v", err))
	}
	key := &keystore.Key{
		Id:         id,
		Address:    crypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}

	dataEncrypted, err := keystore.EncryptKey(key, "passphrase", keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		t.Fatal(err)
	}

	file, err := os.Create("account")
	if err != nil {
		t.Fatalf("Error creating file: %v", err)
	}
	defer file.Close()

	n, err := file.Write(dataEncrypted)
	if err != nil {
		t.Fatalf("Error writing file: %v", err)
	}
	fmt.Printf("Wrote %d bytes using os.Create\n", n)
}

func TestGetSubmitter(t *testing.T) {
	submitter, err := NewSubmitter(config.DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}
	t.Log(submitter.Transactor.Address().String())

	dataHash := sha256.Sum256([]byte("data hash"))
	da := sha256.Sum256([]byte("da"))
	data, err := submitter.GetSubmitter(submitter.Transactor.Address(), dataHash, da)
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fmt.Sprintf("da: %s - cid: %s", da, data))
}

func TestSubmitData(t *testing.T) {
	submitter, err := NewSubmitter(config.DefaultConfig())
	if err != nil {
		t.Fatal(err)
	}
	dataHash := sha256.Sum256([]byte("data hash"))
	da := sha256.Sum256([]byte("da"))
	cid := []byte("cid")
	tx, err := submitter.SubmitData(dataHash, da, cid)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(tx)

	data, err := submitter.GetSubmitter(submitter.Transactor.Address(), dataHash, da)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(data))

	fmt.Println(fmt.Sprintf("cid: %s", data))
}
