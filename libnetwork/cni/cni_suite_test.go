//go:build linux

package cni_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/containers/common/internal/attributedstring"
	"github.com/containers/common/libnetwork/cni"
	"github.com/containers/common/libnetwork/types"
	"github.com/containers/common/pkg/config"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var cniPluginDirs = []string{
	"/usr/libexec/cni",
	"/usr/lib/cni",
	"/usr/local/lib/cni",
	"/opt/cni/bin",
}

func TestCni(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CNI Suite")
}

func getNetworkInterface(cniConfDir string) (types.ContainerNetwork, error) {
	return cni.NewCNINetworkInterface(&cni.InitConfig{
		CNIConfigDir: cniConfDir,
		Config: &config.Config{
			Network: config.NetworkConfig{
				CNIPluginDirs: attributedstring.NewSlice(cniPluginDirs),
			},
		},
	})
}

func SkipIfNoDnsname() {
	for _, path := range cniPluginDirs {
		f, err := os.Stat(filepath.Join(path, "dnsname"))
		if err == nil && f.Mode().IsRegular() {
			return
		}
	}
	Skip("dnsname cni plugin needs to be installed for this test")
}
