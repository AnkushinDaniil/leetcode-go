package longestPalindrome

import "testing"

type test struct {
	s      string
	output int
}

var tests = []test{
	{
		s:      "abccccdd",
		output: 7,
	},
	{
		s:      "a",
		output: 1,
	},
	{
		s:      "civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth",
		output: 983,
	},
}

func Test(t *testing.T) {
	for _, testCase := range tests {
		runTest(t, testCase)
	}
}

func runTest(t *testing.T, testCase test) {
	res := longestPalindrome(testCase.s)
	if testCase.output != res {
		t.Errorf("expected: %v, \ngot: %v", testCase.output, res)
	}
}
