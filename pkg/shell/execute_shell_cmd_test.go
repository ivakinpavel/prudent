package shell

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_no_error(t *testing.T) {
	err, stdout, stderr := ExecuteShellCmd("echo -n test")
	assert.NoError(t, err)
	assert.Equal(t, "test", stdout)
	assert.Equal(t, "", stderr)
}

func Test_an_error(t *testing.T) {
	err, stdout, stderr := ExecuteShellCmd("cat not_existing_file.txt")
	assert.Error(t, err)
	assert.Equal(t, "", stdout)
	assert.Equal(t, "cat: not_existing_file.txt: No such file or directory\n", stderr)
}
