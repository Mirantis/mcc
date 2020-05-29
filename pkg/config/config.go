package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"

	api "github.com/Mirantis/mcc/pkg/apis/v1beta1"
	validator "github.com/go-playground/validator/v10"
	log "github.com/sirupsen/logrus"
)

// FromYaml loads the cluster config from given yaml data
func FromYaml(data []byte) (api.ClusterConfig, error) {
	c := api.ClusterConfig{}

	err := yaml.Unmarshal(data, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

// Validate validates that everything in the config makes sense
// Currently we do only very "static" validation using https://github.com/go-playground/validator
func Validate(c *api.ClusterConfig) error {
	validator := validator.New()
	return validator.Struct(c)
}

// ResolveClusterFile looks for the cluster.yaml file, based on the value.
// It returns the contents of this file as []byte if found,
// or error if it didn't.
func ResolveClusterFile(clusterFile string) ([]byte, error) {
	file, err := openClusterFile(clusterFile)
	defer file.Close()

	buf, err := ioutil.ReadAll(file)
	if err != nil {
		return []byte{}, fmt.Errorf("failed to read file: %v", err)
	}
	return buf, nil
}

func openClusterFile(clusterFile string) (*os.File, error) {
	// the first option always is the file name provided by the user
	possibleOptions := []string{clusterFile}
	if strings.HasSuffix(clusterFile, ".yaml") {
		possibleOptions = append(possibleOptions, strings.ReplaceAll(clusterFile, ".yaml", ".yml"))
	}
	if strings.HasSuffix(clusterFile, ".yml") {
		possibleOptions = append(possibleOptions, strings.ReplaceAll(clusterFile, ".yml", ".yaml"))
	}

	var (
		file *os.File
		fp   string
		err  error
	)
	// iterate over all possible options
	for _, option := range possibleOptions {
		if _, err := os.Stat(option); err != nil {
			continue
		}

		file, fp, err = openFileWithName(option)
		if err != nil {
			return nil, fmt.Errorf("error while opening cluster file %s: %w", option, err)
		}
		log.Debugf("opened config file from %s", fp)
		return file, nil
	}

	return nil, fmt.Errorf("can not find cluster configuration file %s: %v", clusterFile, err)

}

func openFileWithName(fileName string) (file *os.File, path string, err error) {
	fp, err := filepath.Abs(fileName)
	if err != nil {
		return nil, "", fmt.Errorf("failed to lookup current directory name: %v", err)
	}
	file, err = os.Open(fp)
	if err != nil {
		return nil, fp, fmt.Errorf("can not find cluster configuration file: %v", err)
	}
	return file, fp, nil
}
