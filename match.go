package btcsubprotos

import (
	"bytes"

	"github.com/conformal/btcscript"
	"github.com/conformal/btcwire"
)

func IsBulletin(tx *btcwire.MsgTx) bool {
	// Takes a TX and determines if it is an ahimsa bulletin.
	magic := Ahimsa.Magic
	return matchFirstOut(tx, magic) && len(tx.TxOut) > 1
}

func IsCounterParty(tx *btcwire.MsgTx) bool {
	// Tries to match pushed data to the counterparty format
	magic := CounterParty.Magic

	for _, txout := range tx.TxOut {
		outdata, err := btcscript.PushedData(txout.PkScript)
		if err != nil {
			return false
		}
		for _, push := range outdata {
			if len(push) < 1+len(magic) {
				continue
			}
			// OP_RETURN
			if bytes.Equal(push[:len(magic)], magic) {
				return true
			}
			// Data encoded in public key (First byte is for faking a ECDSA PK).
			if bytes.Equal(push[1:len(magic)+1], magic) {
				return true
			}
		}
	}
	return false
}

func matchFirstOut(tx *btcwire.MsgTx, magic []byte) bool {
	// Tests to see if the first txout in the tx matches the magic bytes.
	firstOutScript := tx.TxOut[0].PkScript

	outdata, err := btcscript.PushedData(firstOutScript)
	if err != nil {
		return false
	}
	if len(outdata[0]) > len(magic) {
		firstpush := outdata[0]
		if bytes.Equal(firstpush[:len(magic)], magic) {
			return true
		}
	}
	return false
}

func IsDocProof(tx *btcwire.MsgTx) bool {
	// Checks to see if tx is a docproof message by looking at its first output.
	magic := DocProof.Magic
	return matchFirstOut(tx, magic)
}
