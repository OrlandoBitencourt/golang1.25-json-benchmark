package benchmarks

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/OrlandoBitencourt/golang1.25-json-benchmark/models"
)

var (
	smallUsers  []models.User
	mediumUsers []models.User
	largeUsers  []models.User
	smallData   []byte
	mediumData  []byte
	largeData   []byte
)

func init() {
	var err error

	// Carregar dados para marshal
	smallData, err = os.ReadFile("../testdata/small.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(smallData, &smallUsers)

	mediumData, err = os.ReadFile("../testdata/medium.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(mediumData, &mediumUsers)

	largeData, err = os.ReadFile("../testdata/large.json")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(largeData, &largeUsers)
}

func BenchmarkMarshalSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(smallUsers)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(mediumUsers)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkMarshalLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(largeUsers)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalSmall(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var users []models.User
		err := json.Unmarshal(smallData, &users)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalMedium(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var users []models.User
		err := json.Unmarshal(mediumData, &users)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkUnmarshalLarge(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var users []models.User
		err := json.Unmarshal(largeData, &users)
		if err != nil {
			b.Fatal(err)
		}
	}
}
