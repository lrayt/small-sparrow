package kit

import "fmt"

type LicenseChecker struct {
}

func NewLicenseChecker() *LicenseChecker {
	return &LicenseChecker{}
}
func (l LicenseChecker) Verify() error {
	return fmt.Errorf("license err:%s", "--->")
}
