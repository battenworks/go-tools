package main

import (
	"os/exec"

	"github.com/battenworks/go-tools/common/v2/console"
)

func main() {
	result, _ := getImpersonationToken("atlantis@eng-infrastructure.iam.gserviceaccount.com")
	console.Yellow(string(result))
}

func getImpersonationToken(accountToImpersonate string) (string, error) {
	cmdArgs := []string{"auth", "print-access-token", "--impersonate-service-account=" + accountToImpersonate}
	cmd := exec.Command("gcloud", cmdArgs...)
	token, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return string(token), nil
}
