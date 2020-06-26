package phase

import (
	"archive/zip"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	api "github.com/Mirantis/mcc/pkg/apis/v1beta2"
	"github.com/Mirantis/mcc/pkg/constant"
	"github.com/Mirantis/mcc/pkg/ucp"
	"github.com/Mirantis/mcc/pkg/util"
	"github.com/mitchellh/go-homedir"

	log "github.com/sirupsen/logrus"
)

// DownloadBundle phase downloads the client bundle to local storage
type DownloadBundle struct {
	Analytics
	Username string
	Password string
}

// Title for the phase
func (p *DownloadBundle) Title() string {
	return "Download Client Bundle"
}

// Run collect all the facts from hosts in parallel
func (p *DownloadBundle) Run(conf *api.ClusterConfig) error {
	m := conf.Spec.Managers()[0]

	tlsConfig, err := p.getTLSConfigFrom(m, conf.Spec.Ucp.ImageRepo, conf.Spec.Ucp.Version)
	if err != nil {
		return fmt.Errorf("error getting TLS config: %w", err)
	}

	url, err := p.resolveURL(conf.Spec.WebURL())
	if err != nil {
		return fmt.Errorf("error while parsing URL: %w", err)
	}

	bundle, err := ucp.GetClientBundle(url, tlsConfig, p.Username, p.Password)
	if err != nil {
		return fmt.Errorf("failed to download admin bundle: %s", err)
	}

	bundleDir, err := p.getBundleDir(conf.Metadata.Name, p.Username)
	if err != nil {
		return err
	}
	err = p.writeBundle(bundleDir, bundle)
	if err != nil {
		return fmt.Errorf("failed to write admin bundle: %s", err)
	}

	return nil
}

func (p *DownloadBundle) getBundleDir(clusterName, username string) (string, error) {
	home, err := homedir.Dir()
	if err != nil {
		return "", err
	}
	return path.Join(home, constant.StateBaseDir, "cluster", clusterName, "bundle", username), nil
}

func (p *DownloadBundle) resolveURL(serverURL string) (*url.URL, error) {
	if !strings.HasPrefix(serverURL, "https://") {
		serverURL = fmt.Sprintf("https://%s", serverURL)
	}
	url, err := url.Parse(serverURL)
	if err != nil {
		return nil, err
	}
	return url, nil
}

func (p *DownloadBundle) getTLSConfigFrom(manager *api.Host, imageRepo, ucpVersion string) (*tls.Config, error) {
	runFlags := []string{"--rm", "-v /var/run/docker.sock:/var/run/docker.sock"}
	if manager.Configurer.SELinuxEnabled() {
		runFlags = append(runFlags, "--security-opt label=disable")
	}
	output, err := manager.ExecWithOutput(fmt.Sprintf(`sudo docker run %s %s/ucp:%s dump-certs --ca`, strings.Join(runFlags, " "), imageRepo, ucpVersion))
	if err != nil {
		return nil, fmt.Errorf("error while exec-ing into the container: %w", err)
	}
	i := strings.Index(output, "-----BEGIN CERTIFICATE-----")
	if i < 0 {
		return nil, fmt.Errorf("malformed certificate")
	}

	cert := []byte(output[i:])
	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(cert)
	if !ok {
		return nil, fmt.Errorf("error while appending certs to PEM")
	}
	return &tls.Config{
		RootCAs: caCertPool,
	}, nil
}

func (p *DownloadBundle) writeBundle(bundleDir string, bundle *zip.Reader) error {
	if err := util.EnsureDir(bundleDir); err != nil {
		return fmt.Errorf("error while creating directory: %w", err)
	}
	log.Debugf("Writing out bundle to %s", bundleDir)
	for _, zf := range bundle.File {
		src, err := zf.Open()
		if err != nil {
			return err
		}
		defer src.Close()
		var data []byte
		data, err = ioutil.ReadAll(src)
		if err != nil {
			return err
		}
		mode := int64(0644)
		if strings.Contains(zf.Name, "key.pem") {
			mode = 0600
		}

		// UCP bundle will contain folders as well as files, if folder exists fd will not be empty
		dir, _ := filepath.Split(zf.Name)
		if dir != "" {
			if err := os.MkdirAll(filepath.Join(bundleDir, dir), 0700); err != nil {
				return err
			}
		}

		err = ioutil.WriteFile(filepath.Join(bundleDir, zf.Name), data, os.FileMode(mode))
		if err != nil {
			return err
		}
	}
	log.Infof("Successfully wrote client bundle to %s", bundleDir)
	return nil
}