//go:build unittest

package sinovate

import (
	"encoding/hex"
	"math/big"
	"os"
	"reflect"
	"testing"

	"github.com/martinboehm/btcutil/chaincfg"
	"github.com/trezor/blockbook/bchain"
	"github.com/trezor/blockbook/bchain/coins/btc"
)

func TestMain(m *testing.M) {
	c := m.Run()
	chaincfg.ResetParams()
	os.Exit(c)
}

func Test_GetAddrDescFromAddress_Testnet(t *testing.T) {
	type args struct {
		address string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "pubkeyhash",
			args:    args{address: "SiS23eTK7RCJmemuhkfjqt6Jtv4vKJmNwH"},
			want:    "76a914e7ded0c4339c31a0eb8b557bb04bb25a9468d01588ac",
			wantErr: false,
		},
		{
			name:    "p2sh-segwit",
			args:    args{address: "3GV9e2AA9G3jtnzZNQQF9iGzgTzBgxnhf3"},
			want:    "a914a24b55e19fb5433856965330a314e80b64fa239d87",
			wantErr: false,
		},
		{
			name:    "bech32-witness_v0_keyhash",
			args:    args{address: "tsin1qc0d29m2ku9flug57wgwx8lah34d9q9hfhd7cg6"},
			want:    "0014c3daa2ed56e153fe229e721c63ffb78d5a5016e9",
			wantErr: false,
		},
		{
			name:    "witness_v1_taproot",
			args:    args{address: "tsin1pk0g0qrup6e3khr9ny2g5h8ws97545fh6ksf4uzw7xsg54e95h4lqpsa484"},
			want:    "5120b3d0f00f81d6636b8cb322914b9dd02fa95a26fab4135e09de34114ae4b4bd7e",
			wantErr: false,
		},
	}
	parser := NewSinovateParser(GetChainParams("test"), &btc.Configuration{})

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parser.GetAddrDescFromAddress(tt.args.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAddrDescFromAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("GetAddrDescFromAddress() = %v, want %v", h, tt.want)
			}
		})
	}
}

var (
	testTx1 bchain.Tx

	testTxPacked1 = "000000788c9e81f66e0200000001d1eb47850efc0fdcfd354333fe0d9917815031a04ae3d7b9a79b9f0f9da0f736000000006a473044022012b0928541e460f3e6a3f224830d2fc8bcb5f7e257f81b244233d65c84ddb85d022028660e5f3b9a3d5e4871b7f5681915f9e18cd9695aff9f4d798a44b765aabc590121038269d5c2c6d5fd99fd5e4f369c2eecace27a2e5647b2df4400c467b9fbdad1b6feffffff0200ab9041000000001976a914e7ded0c4339c31a0eb8b557bb04bb25a9468d01588ac6c0aa7c5860400001976a9146441543c593a50d80f97ad14fae04c5f6858d34d88ac77000000"
)

func init() {
	testTx1 = bchain.Tx{
		Hex:       "0200000001d1eb47850efc0fdcfd354333fe0d9917815031a04ae3d7b9a79b9f0f9da0f736000000006a473044022012b0928541e460f3e6a3f224830d2fc8bcb5f7e257f81b244233d65c84ddb85d022028660e5f3b9a3d5e4871b7f5681915f9e18cd9695aff9f4d798a44b765aabc590121038269d5c2c6d5fd99fd5e4f369c2eecace27a2e5647b2df4400c467b9fbdad1b6feffffff0200ab9041000000001976a914e7ded0c4339c31a0eb8b557bb04bb25a9468d01588ac6c0aa7c5860400001976a9146441543c593a50d80f97ad14fae04c5f6858d34d88ac77000000",
		Blocktime: 1642085815,
		Txid:      "b1d5ad4e857d6660101f10baa0b0673b51b8a7a77a69278e21d16e42466e2af2",
		LockTime:  119,
		Version:   2,
		Vin: []bchain.Vin{
			{
				ScriptSig: bchain.ScriptSig{
					Hex: "473044022012b0928541e460f3e6a3f224830d2fc8bcb5f7e257f81b244233d65c84ddb85d022028660e5f3b9a3d5e4871b7f5681915f9e18cd9695aff9f4d798a44b765aabc590121038269d5c2c6d5fd99fd5e4f369c2eecace27a2e5647b2df4400c467b9fbdad1b6",
				},
				Txid:     "36f7a09d0f9f9ba7b9d7e34aa031508117990dfe334335fddc0ffc0e8547ebd1",
				Vout:     0,
				Sequence: 4294967294,
			},
		},
		Vout: []bchain.Vout{
			{
				ValueSat: *big.NewInt(1100000000),
				N:        0,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a914e7ded0c4339c31a0eb8b557bb04bb25a9468d01588ac",
					Addresses: []string{
						"SiS23eTK7RCJmemuhkfjqt6Jtv4vKJmNwH",
					},
				},
			},
			{
				ValueSat: *big.NewInt(4976888187500),
				N:        1,
				ScriptPubKey: bchain.ScriptPubKey{
					Hex: "76a9146441543c593a50d80f97ad14fae04c5f6858d34d88ac",
					Addresses: []string{
						"SWS6sAcXFE3tzDqknZv6HxHdSuX4D7AeT5",
					},
				},
			},
		},
	}
}

func Test_PackTx(t *testing.T) {
	type args struct {
		tx        bchain.Tx
		height    uint32
		blockTime int64
		parser    *SinovateParser
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "sinovate-1",
			args: args{
				tx:        testTx1,
				height:    120,
				blockTime: 1642085815,
				parser:    NewSinovateParser(GetChainParams("test"), &btc.Configuration{}),
			},
			want:    testTxPacked1,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.args.parser.PackTx(&tt.args.tx, tt.args.height, tt.args.blockTime)
			if (err != nil) != tt.wantErr {
				t.Errorf("packTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			h := hex.EncodeToString(got)
			if !reflect.DeepEqual(h, tt.want) {
				t.Errorf("packTx() = %v, want %v", h, tt.want)
			}
		})
	}
}

func Test_UnpackTx(t *testing.T) {
	type args struct {
		packedTx string
		parser   *SinovateParser
	}
	tests := []struct {
		name    string
		args    args
		want    *bchain.Tx
		want1   uint32
		wantErr bool
	}{
		{
			name: "sinovate-1",
			args: args{
				packedTx: testTxPacked1,
				parser:   NewSinovateParser(GetChainParams("test"), &btc.Configuration{}),
			},
			want:    &testTx1,
			want1:   120,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, _ := hex.DecodeString(tt.args.packedTx)
			got, got1, err := tt.args.parser.UnpackTx(b)
			if (err != nil) != tt.wantErr {
				t.Errorf("unpackTx() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackTx() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("unpackTx() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
