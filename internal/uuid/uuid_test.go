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
	encodedIDs := make([]string, len(ids))

	b.Run("Encode", func(b *testing.B) {
		b.ResetTimer() // TODO?
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
		b.ResetTimer() // TODO?
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

func BenchmarkBase64StdString(b *testing.B) {
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

func BenchmarkBase32StdString(b *testing.B) {
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

func BenchmarkShortUUIDV3String(b *testing.B) {
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

func BenchmarkShortUUIDV4String(b *testing.B) {
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

func BenchmarkULIDV2CrockfordBase32String(b *testing.B) {
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

func BenchmarkBTCBase58String(b *testing.B) {
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
