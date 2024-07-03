package bip32

import "github.com/pkg/errors"

// BitcoinMainnetPrivate is the version that is used for
// bitcoin mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var BitcoinMainnetPrivate = [4]byte{
	0x04,
	0x88,
	0xad,
	0xe4,
}

// BitcoinMainnetPublic is the version that is used for
// bitcoin mainnet bip32 public extended keys.
// Ecnodes to xpub in base58.
var BitcoinMainnetPublic = [4]byte{
	0x04,
	0x88,
	0xb2,
	0x1e,
}

// SadMainnetPrivate is the version that is used for
// Sad mainnet bip32 private extended keys.
// Ecnodes to xprv in base58.
var SadMainnetPrivate = [4]byte{
	0x03,
	0x8f,
	0x2e,
	0xf4,
}

// SadMainnetPublic is the version that is used for
// Sad mainnet bip32 public extended keys.
// Ecnodes to kpub in base58.
var SadMainnetPublic = [4]byte{
	0x03,
	0x8f,
	0x33,
	0x2e,
}

// SadTestnetPrivate is the version that is used for
// Sad testnet bip32 public extended keys.
// Ecnodes to ktrv in base58.
var SadTestnetPrivate = [4]byte{
	0x03,
	0x90,
	0x9e,
	0x07,
}

// SadTestnetPublic is the version that is used for
// Sad testnet bip32 public extended keys.
// Ecnodes to ktub in base58.
var SadTestnetPublic = [4]byte{
	0x03,
	0x90,
	0xa2,
	0x41,
}

// SadDevnetPrivate is the version that is used for
// Sad devnet bip32 public extended keys.
// Ecnodes to kdrv in base58.
var SadDevnetPrivate = [4]byte{
	0x03,
	0x8b,
	0x3d,
	0x80,
}

// SadDevnetPublic is the version that is used for
// Sad devnet bip32 public extended keys.
// Ecnodes to xdub in base58.
var SadDevnetPublic = [4]byte{
	0x03,
	0x8b,
	0x41,
	0xba,
}

// SadSimnetPrivate is the version that is used for
// Sad simnet bip32 public extended keys.
// Ecnodes to ksrv in base58.
var SadSimnetPrivate = [4]byte{
	0x03,
	0x90,
	0x42,
	0x42,
}

// SadSimnetPublic is the version that is used for
// Sad simnet bip32 public extended keys.
// Ecnodes to xsub in base58.
var SadSimnetPublic = [4]byte{
	0x03,
	0x90,
	0x46,
	0x7d,
}

func toPublicVersion(version [4]byte) ([4]byte, error) {
	switch version {
	case BitcoinMainnetPrivate:
		return BitcoinMainnetPublic, nil
	case SadMainnetPrivate:
		return SadMainnetPublic, nil
	case SadTestnetPrivate:
		return SadTestnetPublic, nil
	case SadDevnetPrivate:
		return SadDevnetPublic, nil
	case SadSimnetPrivate:
		return SadSimnetPublic, nil
	}

	return [4]byte{}, errors.Errorf("unknown version %x", version)
}

func isPrivateVersion(version [4]byte) bool {
	switch version {
	case BitcoinMainnetPrivate:
		return true
	case SadMainnetPrivate:
		return true
	case SadTestnetPrivate:
		return true
	case SadDevnetPrivate:
		return true
	case SadSimnetPrivate:
		return true
	}

	return false
}
