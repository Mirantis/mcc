package mke

import (
	"archive/zip"
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/Mirantis/mcc/pkg/api"
	"github.com/Mirantis/mcc/pkg/exec"

	log "github.com/sirupsen/logrus"
)

// AuthToken represents a session token
type AuthToken struct {
	ID        string `json:"token_id,omitempty"`
	Token     string `json:"auth_token,omitempty"`
	UserAgent string `json:"user_agent,omitempty"`
}

// Credentials represents a username/password pair for mke login
type Credentials struct {
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// CollectFacts gathers the current status of installed mke setup
// Currently we only need to know the existing version and whether mke is installed or not.
// In future we probably need more.
func CollectFacts(swarmLeader *api.Host, mkeMeta *api.MKEMetadata) error {
	output, err := swarmLeader.ExecWithOutput(swarmLeader.Configurer.DockerCommandf(`inspect --format '{{.Config.Image}}' ucp-proxy`))
	if err != nil {
		if strings.Contains(output, "No such object") {
			mkeMeta.Installed = false
			mkeMeta.InstalledVersion = ""
			return nil
		}
		return err
	}

	vparts := strings.Split(output, ":")
	if len(vparts) != 2 {
		return fmt.Errorf("malformed version output: %s", output)
	}
	repo := vparts[0][:strings.LastIndexByte(vparts[0], '/')]

	mkeMeta.Installed = true
	mkeMeta.InstalledVersion = vparts[1]
	mkeMeta.InstalledBootstrapImage = fmt.Sprintf("%s:/mke:%s", repo, vparts[1])

	// Find out calico data plane by inspecting the calico container's env variables
	cmd := swarmLeader.Configurer.DockerCommandf(`ps --filter label=name="Calico node" --format {{.ID}}`)
	calicoContainer, err := swarmLeader.ExecWithOutput(cmd)

	if calicoContainer != "" && err != nil {
		log.Debugf("%s: calico container found: %s", swarmLeader, calicoContainer)
		cmd := swarmLeader.Configurer.DockerCommandf(`inspect %s --format {{.Config.Env}}`, calicoContainer)
		err := swarmLeader.Exec(fmt.Sprintf("%s | tr ' ' '\n' | grep FELIX_VXLAN= | grep -q true", cmd))
		if err != nil {
			mkeMeta.VXLAN = true
			log.Debugf("%s: has calico VXLAN enabled", swarmLeader)
		}
	}

	return nil
}

// GetClientBundle fetches the client bundle from mke
// It returns a *zip.Reader with the contents of the bundle
// or an error if such occurred.
func GetClientBundle(mkeURL *url.URL, tlsConfig *tls.Config, username, password string) (*zip.Reader, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig,
		},
	}

	// Login and get a token for the user
	token, err := GetToken(client, mkeURL, username, password)
	if err != nil {
		return nil, fmt.Errorf("Failed to get token for (%s:%s) : %s", username, password, err)
	}

	mkeURL.Path = "/api/clientbundle"

	// Now download the bundle
	req, err := http.NewRequest(http.MethodGet, mkeURL.String(), nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", token))
	resp, err := client.Do(req)
	if err != nil {
		log.Debugf("Failed to get bundle: %v", err)
		return nil, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		if err == nil {
			return nil, fmt.Errorf("Failed to get client bundle (%d): %s", resp.StatusCode, string(body))
		}
		return nil, err
	}
	return zip.NewReader(bytes.NewReader(body), int64(len(body)))
}

// GetToken gets a mke Authtoken from the given mkeURL
func GetToken(client *http.Client, mkeURL *url.URL, username, password string) (string, error) {
	mkeURL.Path = "/auth/login"
	creds := Credentials{
		Username: username,
		Password: password,
	}

	reqJSON, err := json.Marshal(creds)
	if err != nil {
		return "", err
	}
	resp, err := client.Post(mkeURL.String(), "application/json", bytes.NewBuffer(reqJSON))
	if err != nil {
		log.Debugf("Failed to POST %s: %v", mkeURL.String(), err)
		return "", err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode == 200 {
		var authToken AuthToken
		if err := json.Unmarshal(body, &authToken); err != nil {
			return "", err
		}
		return authToken.Token, nil
	}
	return "", fmt.Errorf("Unexpected error logging in to mke: %s", string(body))
}

// GetTLSConfigFrom retrieves the valid tlsConfig from the given mke manager
func GetTLSConfigFrom(manager *api.Host, imageRepo, mkeVersion string) (*tls.Config, error) {
	runFlags := []string{"--rm", "-v /var/run/docker.sock:/var/run/docker.sock"}
	if manager.Configurer.SELinuxEnabled() {
		runFlags = append(runFlags, "--security-opt label=disable")
	}
	output, err := manager.ExecWithOutput(fmt.Sprintf(`sudo docker run %s %s/ucp:%s dump-certs --ca`, strings.Join(runFlags, " "), imageRepo, mkeVersion), exec.Redact(`[A-Za-z0-9+/=_\-]{64}`))
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