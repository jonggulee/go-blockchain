package wallet

import (
	"crypto/x509"
	"encoding/hex"
	"os"
	"reflect"
	"testing"
)

const (
	testKey     string = "30770201010420863d34287fc4fc45c86326dbb953bab0a98a4aafc0e7eeb7b9b7c59e4d055cf6a00a06082a8648ce3d030107a144034200041bd4e7c1564f820ac678c3b5b0fa1c78e063aee5d5596396833e01422cd4bea17988ba7cafda35bb343bc44b10eeae2ffc299f147c393d293f75706d90a3dcf9"
	testPayload string = "00c484f521331041a05ab86ac25fbb5a0016b06eef9691247e62a7ef5dcabfd8"
	testSig     string = "49721e0f5b1559cbfc56650e2111f9cdd2964f1b5f149b52a0858ef8e9f080eb2b8b3e9a880535b35c3c6a92d8a14d6e709b7fe4b1be6b34324d4700f2576470"
)

type fakeLayer struct {
	fakeHasWalletFile func() bool
}

func (f fakeLayer) hasWalletFile() bool {
	return f.fakeHasWalletFile()

}

func (fakeLayer) writeFile(name string, data []byte, perm os.FileMode) error {
	return nil
}

func (fakeLayer) readFile(name string) ([]byte, error) {
	return x509.MarshalECPrivateKey(makeTestWallet().privateKey)
}

func TestWallet(t *testing.T) {
	t.Run("New Wallet is created", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool {
				t.Log("I have been called")
				return false
			},
		}
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("New Wallet should return a new wallet instance")
		}
	})
	t.Run("Wallet is restored", func(t *testing.T) {
		files = fakeLayer{
			fakeHasWalletFile: func() bool {
				t.Log("I have been called")
				return true
			},
		}
		w = nil
		tw := Wallet()
		if reflect.TypeOf(tw) != reflect.TypeOf(&wallet{}) {
			t.Error("wallet() should return a new wallet instance")
		}
	})
}

func makeTestWallet() *wallet {
	w := &wallet{}
	b, _ := hex.DecodeString(testKey)
	key, _ := x509.ParseECPrivateKey(b)
	w.privateKey = key
	w.Address = aFromK(key)
	return w
}

func TestSign(t *testing.T) {
	s := Sign(testPayload, makeTestWallet())
	_, err := hex.DecodeString(s)
	if err != nil {
		t.Errorf("Sign should return hex encoded string, got %s", s)
	}
}

func TestVerify(t *testing.T) {
	type test struct {
		input string
		ok    bool
	}
	tests := []test{
		{input: testPayload, ok: true},
		{input: "03c484f521331041a05ab86ac25fbb5a0016b06eef9691247e62a7ef5dcabfd8", ok: false},
	}
	for _, tc := range tests {
		w := makeTestWallet()
		ok := Verify(testSig, tc.input, w.Address)
		if ok != tc.ok {
			t.Error("Verify() colud net verify testSignature and testPayload")
		}
	}

}

func TestRestoreBigInts(t *testing.T) {
	_, _, err := restoreBigInts("xx")
	if err == nil {
		t.Error("restoreBigInts should return error when payload is not hex encoded")
	}
}
