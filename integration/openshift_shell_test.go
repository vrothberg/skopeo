// +build openshift_shell

package main

import (
	"os"
	"os/exec"

	"github.com/go-check/check"
)

/*
TestRunShell is not really a test; it is a convenient way to use the registry setup code
in openshift.go and CopySuite to get an interactive environment for experimentation.

To use it, run:
	sudo make shell
to start a container, then within the container:
	SKOPEO_CONTAINER_TESTS=1 PS1='nested> ' go test -tags openshift_shell -timeout=24h ./integration -v -check.v -check.vv -check.f='CopySuite.TestRunShell'

An example of what can be done within the container:
	cd ..; make bin/skopeo install
	./skopeo --tls-verify=false  copy --sign-by=personal@example.com docker://busybox:latest atomic:localhost:5000/myns/personal:personal
	oc get istag personal:personal -o json
	curl -L -v 'http://localhost:5000/v2/'
	cat ~/.docker/config.json
	curl -L -v 'http://localhost:5000/openshift/token&scope=repository:myns/personal:pull' --header 'Authorization: Basic $auth_from_docker'
	curl -L -v 'http://localhost:5000/v2/myns/personal/manifests/personal' --header 'Authorization: Bearer $token_from_oauth'
	curl -L -v 'http://localhost:5000/extensions/v2/myns/personal/signatures/$manifest_digest' --header 'Authorization: Bearer $token_from_oauth'
*/
func (s *CopySuite) TestRunShell(c *check.C) {
	cmd := exec.Command("bash", "-i")
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	c.Assert(err, check.IsNil)
	cmd.Stdin = tty
	cmd.Stdout = tty
	cmd.Stderr = tty
	err = cmd.Run()
	c.Assert(err, check.IsNil)
}
