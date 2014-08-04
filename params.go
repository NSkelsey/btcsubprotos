package btcsubprotos

import "github.com/NSkelsey/protocol/ahimsa"

type Params struct {
	Magic []byte // Magic bytes associated with some protocol
}

var (
	Ahimsa Params = Params{
		Magic: ahimsa.Magic[:],
	}
	CounterParty Params = Params{
		Magic: []byte{
			0x43, 0x4e, 0x54, 0x52, 0x50, 0x52, 0x54, 0x59, /* | .CNTRPRTY | */
		},
	}
	CounterPartyTestnet Params = Params{
		Magic: []byte{
			0x58, 0x58, /* | XX | */
		},
	}
	DocProof Params = Params{
		Magic: []byte{
			0x44, 0x4f, 0x43, 0x50, 0x52, 0x4f, 0x4f, 0x46, /* | DOCPROOF | */
		},
	}
)
