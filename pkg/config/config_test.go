package config

import (
	"testing"

	validator "github.com/go-playground/validator/v10"

	"github.com/stretchr/testify/require"
)

func TestNonExistingHostsFails(t *testing.T) {
	data := `
hosts:
`
	c := loadYaml(t, data)
	err := c.Validate()
	require.Error(t, err)

	validateErrorField(t, err, "Hosts")
}

func TestHostAddressValidationWithInvalidIP(t *testing.T) {
	data := `
hosts:
- address: "512.1.2.3"
`
	c := loadYaml(t, data)

	err := c.Validate()
	require.Error(t, err)
	validateErrorField(t, err, "Address")
}

func TestHostAddressValidationWithValidIP(t *testing.T) {
	data := `
hosts:
- address: "10.10.10.10"
`
	c := loadYaml(t, data)

	err := c.Validate()
	require.NotContains(t, getAllErrorFields(err), "Address")
}

func TestHostAddressValidationWithInvalidHostname(t *testing.T) {
	data := `
hosts:
- address: "1-2-foo"
`
	c := loadYaml(t, data)

	err := c.Validate()
	require.Error(t, err)
	validateErrorField(t, err, "Address")
}

func TestHostAddressValidationWithValidHostname(t *testing.T) {
	data := `
hosts:
- address: "foo.bar.com"
`
	c := loadYaml(t, data)

	err := c.Validate()
	require.NotContains(t, getAllErrorFields(err), "Address")

}

func TestHostSshPortValidation(t *testing.T) {
	data := `
hosts:
- address: "1.2.3.4"
  sshPort: 0
`
	c := loadYaml(t, data)

	err := c.Validate()
	require.Error(t, err)
	validateErrorField(t, err, "SSHPort")
}

func TestHostSshKeyValidation(t *testing.T) {
	data := `
hosts:
- address: "1.2.3.4"
  sshPort: 22
  sshKeyPath: /path/to/nonexisting/key
`
	c := loadYaml(t, data)

	err := c.Validate()
	require.Error(t, err)
	validateErrorField(t, err, "SSHKeyPath")
}

func TestHostRoleValidation(t *testing.T) {
	data := `
hosts:
- address: "1.2.3.4"
  sshPort: 22
  role: foobar
`
	c := loadYaml(t, data)
	err := c.Validate()
	require.Error(t, err)
	validateErrorField(t, err, "Role")
}

// Just a small helper to load the config struct from yaml to get defaults etc. in place
func loadYaml(t *testing.T, data string) *ClusterConfig {
	c, err := FromYaml([]byte(data))
	if err != nil {
		t.Error(err)
	}
	return &c
}

// checks that the validation errors contains error for the expected field
func validateErrorField(t *testing.T, err error, field string) {
	fields := getAllErrorFields(err)

	require.Contains(t, fields, field)
}

func getAllErrorFields(err error) []string {
	validationErrors := err.(validator.ValidationErrors)
	fields := make([]string, len(validationErrors))

	// Collect all fields that failed validation
	// Also "store" the validation error for the expected field so that we can return it
	// and the correcponding test can further validate it if needed
	for _, fieldError := range validationErrors {
		fields = append(fields, fieldError.Field())
	}

	return fields
}