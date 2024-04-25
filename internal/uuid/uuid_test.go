package uuid_test

import (
	"encoding/base32"
	"encoding/base64"
	"flag"
	"log"
	"os"
	"testing"

	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
	shortuuidv3 "github.com/lithammer/shortuuid/v3"
	shortuuidv4 "github.com/lithammer/shortuuid/v4"
	"github.com/oklog/ulid/v2"
	typeidbase32 "go.jetpack.io/typeid/base32"
)

const (
	crockford32Alphabet = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"
)

var (
	ids []uuid.UUID
)

func TestMain(m *testing.M) {
	flag.Parse()

	testUUIDs, err := Generate(100)
	if err != nil {
		log.Fatal(err)
	}

	ids = testUUIDs

	os.Exit(m.Run())
}

// Generate takes a number and generates the given amount of UUIDs, returning an
// error (and an empty result slice) if any error occurs.
func Generate(num int) ([]uuid.UUID, error) {
	generated := make([]uuid.UUID, 0, num)

	for range num {
		newID, err := uuid.NewRandom()
		if err != nil {
			return nil, err
		}

		generated = append(generated, newID)
	}

	return generated, nil
}

type encodeFunc func(id uuid.UUID) (string, error)
type decodeFunc func(encoded string) (uuid.UUID, error)

func benchCodec(b *testing.B, encode encodeFunc, decode decodeFunc) {
	b.Helper()

	encodedIDs := make([]string, len(ids))

	b.Run("Encode", func(b *testing.B) {
		for range b.N {
			for i, id := range ids {
				encoded, err := encode(id)
				if err != nil {
					b.Error(err)
				}

				encodedIDs[i] = encoded
			}
		}
	})

	b.Run("Decode", func(b *testing.B) {
		for range b.N {
			for _, encodedID := range encodedIDs {
				_, err := decode(encodedID)
				if err != nil {
					b.Error(err)
				}
			}
		}
	})
}

func BenchmarkDefaultString(b *testing.B) {
	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			return id.String(), nil
		},
		func(encoded string) (uuid.UUID, error) {
			return uuid.Parse(encoded)
		},
	)
}

func BenchmarkBase64Std(b *testing.B) {
	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			return base64.StdEncoding.EncodeToString(id[:]), nil
		},
		func(encoded string) (uuid.UUID, error) {
			raw, err := base64.StdEncoding.DecodeString(encoded)
			if err != nil {
				return uuid.Nil, err
			}

			return uuid.FromBytes(raw)
		},
	)
}

func BenchmarkBase64RawURLPrePadded(b *testing.B) {
	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			frontPadded := append([]byte{0, 0}, id[:]...)

			return base64.RawURLEncoding.EncodeToString(frontPadded)[2:], nil
		},
		func(encoded string) (uuid.UUID, error) {
			raw, err := base64.RawURLEncoding.DecodeString(encoded)
			if err != nil {
				return uuid.Nil, err
			}

			return uuid.FromBytes(raw)
		},
	)
}

func BenchmarkBase32Std(b *testing.B) {
	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			return base32.StdEncoding.EncodeToString(id[:]), nil
		},
		func(encoded string) (uuid.UUID, error) {
			raw, err := base32.StdEncoding.DecodeString(encoded)
			if err != nil {
				return uuid.Nil, err
			}

			return uuid.FromBytes(raw)
		},
	)
}

func BenchmarkBase32HexPrePadded(b *testing.B) {
	encoding := base32.HexEncoding.WithPadding(base32.NoPadding)

	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			frontPadded := append([]byte{0, 0, 0, 0}, id[:]...)

			return encoding.EncodeToString(frontPadded)[6:], nil
		},
		func(encoded string) (uuid.UUID, error) {
			raw, err := encoding.DecodeString(encoded)
			if err != nil {
				return uuid.Nil, err
			}

			return uuid.FromBytes(raw)
		},
	)
}

func BenchmarkStdLibCrockfordBase32(b *testing.B) {
	encoding := base32.NewEncoding(crockford32Alphabet).WithPadding(base32.NoPadding)

	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			frontPadded := append([]byte{0, 0, 0, 0}, id[:]...)

			return encoding.EncodeToString(frontPadded)[6:], nil
		},
		func(encoded string) (uuid.UUID, error) {
			raw, err := encoding.DecodeString(encoded)
			if err != nil {
				return uuid.Nil, err
			}

			return uuid.FromBytes(raw)
		},
	)
}

func BenchmarkULIDV2CrockfordBase32(b *testing.B) {
	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			// ULID and UUIDs are compatible 16 byte arrays
			return ulid.ULID(id).String(), nil
		},
		func(encoded string) (uuid.UUID, error) {
			// ULID and UUIDs are compatible 16 byte arrays
			ulid := &ulid.ULID{}
			err := ulid.Scan(encoded)

			return uuid.UUID(*ulid), err
		},
	)
}

func BenchmarkTypeIDCrockfordBase32(b *testing.B) {
	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			return typeidbase32.Encode(id), nil
		},
		func(encoded string) (uuid.UUID, error) {
			raw, err := typeidbase32.Decode(encoded)
			if err != nil {
				return uuid.Nil, err
			}

			return uuid.FromBytes(raw)
		},
	)
}

func BenchmarkShortUUIDV3(b *testing.B) {
	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			return shortuuidv3.DefaultEncoder.Encode(id), nil
		},
		func(encoded string) (uuid.UUID, error) {
			return shortuuidv3.DefaultEncoder.Decode(encoded)
		},
	)
}

func BenchmarkShortUUIDV4(b *testing.B) {
	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			return shortuuidv4.DefaultEncoder.Encode(id), nil
		},
		func(encoded string) (uuid.UUID, error) {
			return shortuuidv4.DefaultEncoder.Decode(encoded)
		},
	)
}

func BenchmarkBTCBase58(b *testing.B) {
	benchCodec(
		b,
		func(id uuid.UUID) (string, error) {
			return base58.Encode(id[:]), nil
		},
		func(encoded string) (uuid.UUID, error) {
			raw := base58.Decode(encoded)

			return uuid.FromBytes(raw)
		},
	)
}
