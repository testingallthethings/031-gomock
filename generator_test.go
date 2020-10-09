package drivinglicence_test

import (
	"drivinglicence"
	"github.com/stretchr/testify/suite"
	"strings"
	"testing"
)

func (s *DrivingLicenceSuite) TestUnderageApplicant() {
 	a := UnderageApplicant{}
 	l := &SpyLogger{}
	r := FakeRand{}

 	lg := drivinglicence.NewNumberGenerator(l, r)
 	_, err := lg.Generate(a)

 	s.Error(err)
	s.Contains(err.Error(), "Underaged")

 	s.Equal(1, l.callCount)
 	s.Contains(l.lastMessage, "Underaged")
}

func (s *DrivingLicenceSuite) TestNoSecondLicence() {
	a := LicenceHolderApplicant{}
	l := &SpyLogger{}
	r := FakeRand{}

	lg := drivinglicence.NewNumberGenerator(l, r)
	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Duplicate")

	s.Equal(1, l.callCount)
	s.Contains(l.lastMessage, "Duplicate")
}

func (s *DrivingLicenceSuite) TestLicenceGeneration() {
	a := ValidApplicant{"MDB", "23082011"}
	l := &SpyLogger{}
	r := FakeRand{}

	lg := drivinglicence.NewNumberGenerator(l, r)
	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("MDB2308201100000", ln)
}

func (s *DrivingLicenceSuite) TestLicenceGenerationShorterInitials() {
	a := ValidApplicant{"MB", "23082011"}
	l := &SpyLogger{}
	r := FakeRand{}

	lg := drivinglicence.NewNumberGenerator(l, r)
	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("MB23082011000000", ln)
}

type DrivingLicenceSuite struct {
	suite.Suite
}

func TestDrivingLicenceSuite(t *testing.T) {
	suite.Run(t, new(DrivingLicenceSuite))
}

type UnderageApplicant struct {}

func (u UnderageApplicant) GetDOB() string {
	return ""
}

func (u UnderageApplicant) GetInitials() string {
	return ""
}

func (u UnderageApplicant) IsOver17() bool {
	return false
}

func (u UnderageApplicant) HoldsLicence() bool {
	return false
}

type LicenceHolderApplicant struct {}

func (l LicenceHolderApplicant) IsOver17() bool {
	return true
}

func (l LicenceHolderApplicant) HoldsLicence() bool {
	return true
}

func (l LicenceHolderApplicant) GetInitials() string {
	return ""
}

func (l LicenceHolderApplicant) GetDOB() string {
	return ""
}

type ValidApplicant struct {
	initials string
	dob string
}

func (v ValidApplicant) IsOver17() bool {
	return true
}

func (v ValidApplicant) HoldsLicence() bool {
	return false
}

func (v ValidApplicant) GetInitials() string {
	return v.initials
}

func (v ValidApplicant) GetDOB() string {
	return v.dob
}

type FakeRand struct {}

func (f FakeRand) GetRandomNumbers(len int) string {
	return strings.Repeat("0", len)
}

type SpyLogger struct {
	callCount int
	lastMessage string
}

func (s *SpyLogger) LogStuff(v string) {
	s.callCount++
	s.lastMessage = v
}



