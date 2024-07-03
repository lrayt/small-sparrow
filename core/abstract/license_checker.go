package abstract

type LicenseChecker interface {
	Verify() error
}
