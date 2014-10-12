package btcsubprotos

import (
	"bytes"

	"github.com/conformal/btcscript"
	"github.com/conformal/btcwire"
)

// Takes a TX and determines if it is an ahimsa bulletin.
// This method only looks for the leading bytes, it does not
// assert anything about the protocol buffer within.
func IsBulletin(tx *btcwire.MsgTx) bool {
	magic := Ahimsa.Magic
	return matchFirstOut(tx, magic) && len(tx.TxOut) > 1
}

// Tries to match pushed data to the counterparty format
func IsCounterParty(tx *btcwire.MsgTx) bool {
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

// Tests to see if the first txout in the tx matches the magic bytes.
func matchFirstOut(tx *btcwire.MsgTx, magic []byte) bool {
	if len(tx.TxOut) == 0 {
		return false
	}
	firstOutScript := tx.TxOut[0].PkScript

	outdata, err := btcscript.PushedData(firstOutScript)
	if err != nil {
		return false
	}
	if len(outdata) > 0 && len(outdata[0]) > len(magic) {
		firstpush := outdata[0]
		if bytes.Equal(firstpush[:len(magic)], magic) {
			return true
		}
	}
	return false
}

// Checks to see if tx is a docproof message by looking at its first output.
func IsDocProof(tx *btcwire.MsgTx) bool {
	magic := DocProof.Magic
	return matchFirstOut(tx, magic)
}
