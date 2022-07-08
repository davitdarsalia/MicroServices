package service

import (
	"regexp"
	"testing"
)

func TestPasswordRegex(t *testing.T) {
	pattern := passwordRegex

	testTable := []struct {
		input    string
		expected bool
	}{
		{
			input:    "Password.1998!",
			expected: true,
		},
		{
			input:    "RandomPassword",
			expected: true,
		},
		{
			input:    "Sdadsadsadsadqwd1343fast23g34tg",
			expected: true,
		},
		{
			input:    "Ad32d.32..43.6y.64.3r.31,e1fr!@#$%^&*%^$^#",
			expected: true,
		},
		//
		{
			input:    "Asdksdsdd",
			expected: true,
		},
		{
			input:    "randomPassword",
			expected: false,
		},
		{
			input:    "sdadsadsadsadqwd1343fast23g34tg",
			expected: false,
		},
		{
			input:    "pass",
			expected: false,
		},
		{
			input:    "passwordRandomGenerated",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		regex, err := regexp.Compile(pattern)

		if err != nil {
			t.Logf("PasswordRegex Failed: %s", err.Error())
			t.Fail()
		}

		match := regex.MatchString(testCase.input)

		if testCase.expected != match {
			t.Logf("Input: %s Failed", testCase.input)
			t.Fail()
		}
	}
}

func TestNameRegex(t *testing.T) {
	pattern := nameRegex

	testTable := []struct {
		input    string
		expected bool
	}{
		{
			input:    "Davit",
			expected: false,
		},
		{
			input:    "Darsalia",
			expected: false,
		},
		{
			input:    "Davit Darsalia",
			expected: true,
		},
		{
			input:    "Luke Stoltman",
			expected: true,
		},
		//
		{
			input:    "Rhoshandiatellyneshiaunneveshenk Koyaanisquatsiuth Williams",
			expected: true,
		},
		{
			input:    "davit",
			expected: false,
		},
		{
			input:    "Luke",
			expected: false,
		},
		{
			input:    "Wolfeschlegelsteinhausenbergerdorff",
			expected: false,
		},
		{
			input:    "wolfeschlegelsteinhausenbergerdorff",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		regex, err := regexp.Compile(pattern)

		if err != nil {
			t.Logf("NameRegex Failed: %s", err.Error())
			t.Fail()
		}

		match := regex.MatchString(testCase.input)

		if testCase.expected != match {
			t.Logf("Input: %s Failed", testCase.input)
			t.Fail()
		}
	}
}

func TestEmailRegex(t *testing.T) {
	pattern := emailRegex

	testTable := []struct {
		input    string
		expected bool
	}{
		{
			input:    "lkrout4@gmail.com",
			expected: true,
		},
		{
			input:    "rtuminelli5@gmail.com",
			expected: true,
		},
		{
			input:    "wcrowth6@gmail.com",
			expected: true,
		},
		{
			input:    "kknee7@gmail.com",
			expected: true,
		},
		//
		{
			input:    "rletford8@gmail.com",
			expected: true,
		},
		{
			input:    "ehacketc@gmail.com",
			expected: true,
		},
		{
			input:    "dlowdesd@gmail.org",
			expected: false,
		},
		{
			input:    "fodesone@gmail.com",
			expected: true,
		},
		{
			input:    "ehacketc@comsenz.com",
			expected: false,
		},
		{
			input:    "dlowdesd@edublogs.org",
			expected: false,
		},
		{
			input:    "fodesone@usatoday.com",
			expected: false,
		},
		{
			input:    "sheminsleym@cmu.edu",
			expected: false,
		},
		{
			input:    "ecowderayn@vk.com",
			expected: false,
		},
		{
			input:    "wiredello@qq.com",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		regex, err := regexp.Compile(pattern)

		if err != nil {
			t.Logf("EmailRegex Failed: %s", err.Error())
			t.Fail()
		}

		match := regex.MatchString(testCase.input)

		if testCase.expected != match {
			t.Logf("Input: %s, Expected: %v  Failed, Got: %v", testCase.input, testCase.expected, testCase.expected != match)
			t.Fail()
		}
	}
}

func TestUsernameRegex(t *testing.T) {
	pattern := usernameRegex

	testTable := []struct {
		input    string
		expected bool
	}{
		{
			input:    "lbeagang",
			expected: true,
		},
		{
			input:    "tjenckeni",
			expected: true,
		},
		{
			input:    "mwybrewo",
			expected: true,
		},
		{
			input:    "lcardenp",
			expected: true,
		},
		//
		{
			input:    "ktrouelq",
			expected: true,
		},
		{
			input:    "mmarielr",
			expected: true,
		},
		{
			input:    "wrushmeres",
			expected: true,
		},
		{
			input:    "tcrickmoort",
			expected: true,
		},
		{
			input:    "dbee",
			expected: false,
		},
		{
			input:    "Rians",
			expected: false,
		},
		{
			input:    "Gb2ni",
			expected: false,
		},
		{
			input:    "cevis",
			expected: false,
		},
		{
			input:    "wupwn",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		regex, err := regexp.Compile(pattern)

		if err != nil {
			t.Logf("UserNameRegex Failed: %s", err.Error())
			t.Fail()
		}

		match := regex.MatchString(testCase.input)

		if testCase.expected != match {
			t.Logf("Input: %s, Expected: %v  Failed, Got: %v", testCase.input, testCase.expected, testCase.expected != match)
			t.Fail()
		}
	}
}

func TestIpAddressRegex(t *testing.T) {
	pattern := ipAddressRegex

	testTable := []struct {
		input    string
		expected bool
	}{
		{
			input:    "187.255.224.43",
			expected: true,
		},
		{
			input:    "48.137.68.246",
			expected: true,
		},
		{
			input:    "216.239.114.7",
			expected: true,
		},
		{
			input:    "93.96.69.138",
			expected: true,
		},
		{
			input:    "85.1.3.127",
			expected: true,
		},
		{
			input:    "21.0.12.162",
			expected: true,
		},
		{
			input:    "207.22.7.154",
			expected: true,
		},
		{
			input:    "102.204.81.99",
			expected: true,
		},
		{
			input:    "221.69.106.220",
			expected: true,
		},
		{
			input:    "40.165.65.217",
			expected: true,
		},
	}

	for _, testCase := range testTable {
		regex, err := regexp.Compile(pattern)

		if err != nil {
			t.Logf("IpAddressRegex Failed: %s", err.Error())
			t.Fail()
		}

		match := regex.MatchString(testCase.input)

		if testCase.expected != match {
			t.Logf("Input: %s, Expected: %v  Failed, Got: %v", testCase.input, testCase.expected, testCase.expected != match)
			t.Fail()
		}
	}
}

func TestPersonalNumberRegex(t *testing.T) {
	pattern := personalNumberRegex

	testTable := []struct {
		input    string
		expected bool
	}{
		{
			input:    "29136229008",
			expected: true,
		},
		{
			input:    "87264344191",
			expected: true,
		},
		{
			input:    "76623462339",
			expected: true,
		},
		{
			input:    "40133397164",
			expected: true,
		},
		{
			input:    "28036304759",
			expected: true,
		},
		{
			input:    "28036304759",
			expected: true,
		},
		{
			input:    "77149295242",
			expected: true,
		},
		{
			input:    "94223917967",
			expected: true,
		},
		{
			input:    "40591650272",
			expected: true,
		},
		{
			input:    "34909544480",
			expected: true,
		},
		{
			input:    "34909544480",
			expected: true,
		},
		{
			input:    "340",
			expected: false,
		},
		{
			input:    "213",
			expected: false,
		},
	}

	for _, testCase := range testTable {
		regex, err := regexp.Compile(pattern)

		if err != nil {
			t.Logf("IpAddressRegex Failed: %s", err.Error())
			t.Fail()
		}

		match := regex.MatchString(testCase.input)

		if testCase.expected != match {
			t.Logf("Input: %s, Expected: %v  Failed, Got: %v", testCase.input, testCase.expected, testCase.expected != match)
			t.Fail()
		}
	}
}
