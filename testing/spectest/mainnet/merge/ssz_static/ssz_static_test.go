package ssz_static

import (
	"testing"

	"github.com/prysmaticlabs/prysm/testing/spectest/shared/merge/ssz_static"
)

func TestMainnet_Merge_SSZStatic(t *testing.T) {
	ssz_static.RunSSZStaticTests(t, "mainnet")
}