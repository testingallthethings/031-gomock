package drivinglicence_test

import (
	"drivinglicence"
	mock_drivinglicence "drivinglicence/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

var ctrl *gomock.Controller
var a *mock_drivinglicence.MockApplicant
var l *mock_drivinglicence.MockLogger
var r *mock_drivinglicence.MockRandomNumbersGenerator

var lg drivinglicence.NumberGenerator

func (s *DrivingLicenceSuite) SetupTest() {
	ctrl = gomock.NewController(s.T())

	a = mock_drivinglicence.NewMockApplicant(ctrl)
	l = mock_drivinglicence.NewMockLogger(ctrl)
	r = mock_drivinglicence.NewMockRandomNumbersGenerator(ctrl)

	lg = drivinglicence.NewNumberGenerator(l, r)
}

func (s *DrivingLicenceSuite) TearDownTest() {
	ctrl.Finish()
}

func (s *DrivingLicenceSuite) TestUnderageApplicant() {
	a.EXPECT().IsOver17().Return(false)
	a.EXPECT().HoldsLicence().Return(false)

 	l.EXPECT().LogStuff("Underaged Applicant, you must be 17 for a licence").Times(1)

 	_, err := lg.Generate(a)

 	s.Error(err)
	s.Contains(err.Error(), "Underaged")
}

func (s *DrivingLicenceSuite) TestNoSecondLicence() {
	a.EXPECT().HoldsLicence().Return(true)

	l.EXPECT().LogStuff("Duplicate Applicant, you can only hold one licence").Times(1)

	_, err := lg.Generate(a)

	s.Error(err)
	s.Contains(err.Error(), "Duplicate")
}

func (s *DrivingLicenceSuite) TestLicenceGeneration() {
	a.EXPECT().HoldsLicence().Return(false)
	a.EXPECT().IsOver17().Return(true)
	a.EXPECT().GetInitials().Return("MDB")
	a.EXPECT().GetDOB().Return("23082011")

	r.EXPECT().GetRandomNumbers(5).Return("00000")

	ln, err := lg.Generate(a)

	s.NoError(err)
	s.Equal("MDB2308201100000", ln)
}

func (s *DrivingLicenceSuite) TestLicenceGenerationShorterInitials() {
	a.EXPECT().HoldsLicence().Return(false)
	a.EXPECT().IsOver17().Return(true)
	a.EXPECT().GetInitials().Return("MB")
	a.EXPECT().GetDOB().Return("23082011")

	r.EXPECT().GetRandomNumbers(6).Return("000000")

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
