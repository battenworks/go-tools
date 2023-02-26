package tfcmd

import "os/exec"

func getImpersonationToken(accountToImpersonate string) (string, error) {
	cmdArgs := []string{"auth", "print-access-token", "--impersonate-service-account=" + accountToImpersonate}
	result, err := exec.Command("gcloud", cmdArgs...).CombinedOutput()

	return string(result), err
}
