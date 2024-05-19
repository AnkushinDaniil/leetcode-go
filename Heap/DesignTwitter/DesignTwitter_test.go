package designTwitter

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	testTable := []struct {
		funcs  []string
		params [][]int
		output [][]int
	}{
		{
			funcs: []string{
				"Twitter",
				"postTweet",
				"getNewsFeed",
				"follow",
				"postTweet",
				"getNewsFeed",
				"unfollow",
				"getNewsFeed",
			},
			params: [][]int{{}, {1, 5}, {1}, {1, 2}, {2, 6}, {1}, {1, 2}, {1}},
			output: [][]int{{}, {}, {5}, {}, {}, {6, 5}, {}, {5}},
		},
		{
			funcs: []string{
				"Twitter",
				"postTweet",
				"getNewsFeed",
				"follow",
				"getNewsFeed",
				"unfollow",
				"getNewsFeed",
			},
			params: [][]int{{}, {1, 1}, {1}, {2, 1}, {2}, {2, 1}, {2}},
			output: [][]int{{}, {}, {1}, {}, {1}, {}, {}},
		},
	}
	for testCaseNum, testCase := range testTable {
		t.Run(fmt.Sprintf("Test №%v", testCaseNum+1), func(t *testing.T) {
			var obj Twitter
			res := make([][]int, len(testCase.funcs))
			for i := 0; i < len(testCase.funcs); i++ {
				switch testCase.funcs[i] {
				case "Twitter":
					obj = Constructor()
					res[i] = []int{}
				case "postTweet":
					obj.PostTweet(testCase.params[i][0], testCase.params[i][1])
					res[i] = []int{}
				case "getNewsFeed":
					res[i] = obj.GetNewsFeed(testCase.params[i][0])
				case "follow":
					obj.Follow(testCase.params[i][0], testCase.params[i][1])
					res[i] = []int{}
				case "unfollow":
					obj.Unfollow(testCase.params[i][0], testCase.params[i][1])
					res[i] = []int{}
				}
			}
			assert.Equal(t, testCase.output, res)
		})
	}
}

func BenchmarkTwitter(b *testing.B) {
	testTable := []struct {
		funcs  []string
		params [][]int
		output [][]int
	}{
		{
			funcs: []string{
				"Twitter",
				"postTweet",
				"getNewsFeed",
				"follow",
				"postTweet",
				"getNewsFeed",
				"unfollow",
				"getNewsFeed",
			},
			params: [][]int{{}, {1, 5}, {1}, {1, 2}, {2, 6}, {1}, {1, 2}, {1}},
			output: [][]int{{}, {}, {5}, {}, {}, {6, 5}, {}, {5}},
		},
		{
			funcs: []string{
				"Twitter",
				"postTweet",
				"getNewsFeed",
				"follow",
				"getNewsFeed",
				"unfollow",
				"getNewsFeed",
			},
			params: [][]int{{}, {1, 1}, {1}, {2, 1}, {2}, {2, 1}, {2}},
			output: [][]int{{}, {}, {1}, {}, {1}, {}, {}},
		},
	}
	for i := 0; i < b.N; i++ {
		for _, testCase := range testTable {
			var obj Twitter
			res := make([][]int, len(testCase.funcs))
			for i := 0; i < len(testCase.funcs); i++ {
				switch testCase.funcs[i] {
				case "Twitter":
					obj = Constructor()
					res[i] = []int{}
				case "postTweet":
					obj.PostTweet(testCase.params[i][0], testCase.params[i][1])
					res[i] = []int{}
				case "getNewsFeed":
					res[i] = obj.GetNewsFeed(testCase.params[i][0])
				case "follow":
					obj.Follow(testCase.params[i][0], testCase.params[i][1])
					res[i] = []int{}
				case "unfollow":
					obj.Unfollow(testCase.params[i][0], testCase.params[i][1])
					res[i] = []int{}
				}
			}

		}
	}
}
