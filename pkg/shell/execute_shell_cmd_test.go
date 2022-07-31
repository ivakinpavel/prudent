package shell

import "testing"

func Test_no_error(t *testing.T) {
	err, stdout, stderr := ExecuteShellCmd("ls -l")
	println(err)
	println(stdout)
	println(stderr)
}
