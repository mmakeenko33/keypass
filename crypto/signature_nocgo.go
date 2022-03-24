"crypto/elliptic"
	"errors"
	"fmt"
	"math/big"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/btcec/v2"
	btc_ecdsa "github.com/btcsuite/btcd/btcec/v2/ecdsa"
)

// Ecrecover returns the uncompressed public key that created the given signature.
func Ecrecover(hash, sig []byte) ([]byte, error) {
	pub, err := SigToPub(hash, sig)
	pub, err := sigToPub(hash, sig)
	if err != nil {
		return nil, err
	}
	bytes := (*btcec.PublicKey)(pub).SerializeUncompressed()
	bytes := pub.SerializeUncompressed()
	return bytes, err
}

// SigToPub returns the public key that created the given signature.
func SigToPub(hash, sig []byte) (*ecdsa.PublicKey, error) {
func sigToPub(hash, sig []byte) (*btcec.PublicKey, error) {
	if len(sig) != SignatureLength {
		return nil, errors.New("invalid signature")
	}
