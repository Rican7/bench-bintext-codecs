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

func BenchmarkEncodeDefaultString(b *testing.B) {
	for range b.N {
		for _, id := range ids {
			id.String()
		}
	}
}

func BenchmarkEncodeBase64StdString(b *testing.B) {
	for range b.N {
		for _, id := range ids {
			base64.StdEncoding.EncodeToString(id[:])
		}
	}
}

func BenchmarkEncodeBase32StdString(b *testing.B) {
	for range b.N {
		for _, id := range ids {
			base32.StdEncoding.EncodeToString(id[:])
		}
	}
}

func BenchmarkEncodeShortUUIDV3String(b *testing.B) {
	for range b.N {
		for _, id := range ids {
			shortuuidv3.DefaultEncoder.Encode(id)
		}
	}
}

func BenchmarkEncodeShortUUIDV4String(b *testing.B) {
	for range b.N {
		for _, id := range ids {
			shortuuidv4.DefaultEncoder.Encode(id)
		}
	}
}

func BenchmarkEncodeULIDV2CrockfordBase32String(b *testing.B) {
	for range b.N {
		for _, id := range ids {
			// ULID and UUIDs are compatible 16 byte arrays
			ulid.ULID(id).String()
		}
	}
}

func BenchmarkEncodeBTCBase58String(b *testing.B) {
	for range b.N {
		for _, id := range ids {
			base58.Encode(id[:])
		}
	}
}
