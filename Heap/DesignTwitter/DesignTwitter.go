package designTwitter

const (
	NUMBER_OF_TWEETS = 10
	NUMBER_OF_USERS  = 8
)

type Node struct {
	count, tweetId, foloweeId, index int
}

type MaxHeap []Node

func parent(i int) int {
	return (i - 1) / 2
}

func left(i int) int {
	return 2*i + 1
}

func right(i int) int {
	return 2*i + 2
}

func NewMaxHeap(cap int) MaxHeap {
	return make(MaxHeap, 0, cap)
}

func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}

func (h *MaxHeap) Insert(node Node) {
	i := len((*h))
	(*h) = append((*h), node)

	for i != 0 && (*h)[parent(i)].count < (*h)[i].count {
		(*h).Swap(i, parent(i))
		i = parent(i)
	}
}

func (h *MaxHeap) Heapify(i int) {
	l := left(i)
	r := right(i)
	biggest := i
	if l < len((*h)) && (*h)[l].count > (*h)[i].count {
		biggest = l
	}
	if r < len((*h)) && (*h)[r].count > (*h)[biggest].count {
		biggest = r
	}
	if biggest != i {
		h.Swap(i, biggest)
		h.Heapify(biggest)
	}
}

func (h *MaxHeap) ExtractMax() Node {
	res := (*h)[0]
	(*h)[0] = (*h)[len((*h))-1]
	(*h) = (*h)[:len((*h))-1]
	(*h).Heapify(0)
	return res
}

type Tweet struct {
	Time, Id int
}

type Twitter struct {
	count     int
	tweetMap  map[int][]Tweet
	followMap map[int]map[int]struct{}
}

func Constructor() Twitter {
	return Twitter{
		count:     0,
		tweetMap:  make(map[int][]Tweet, NUMBER_OF_USERS),
		followMap: make(map[int]map[int]struct{}, NUMBER_OF_USERS),
	}
}

func (t *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := t.tweetMap[userId]; !ok {
		t.tweetMap[userId] = make([]Tweet, 0, NUMBER_OF_TWEETS)
	}
	t.tweetMap[userId] = append(t.tweetMap[userId], Tweet{
		Time: t.count,
		Id:   tweetId,
	})
	t.count++
}

func (t *Twitter) GetNewsFeed(userId int) []int {
	res := make([]int, 0, NUMBER_OF_TWEETS)
	heap := NewMaxHeap(NUMBER_OF_TWEETS + 1)

	if _, ok := t.followMap[userId]; !ok {
		t.followMap[userId] = make(map[int]struct{}, NUMBER_OF_USERS)
	}
	t.followMap[userId][userId] = struct{}{}

	for followeeId := range t.followMap[userId] {
		if _, ok := t.tweetMap[followeeId]; ok {
			index := len(t.tweetMap[followeeId]) - 1
			heap.Insert(Node{
				count:     t.tweetMap[followeeId][index].Time,
				tweetId:   t.tweetMap[followeeId][index].Id,
				foloweeId: followeeId,
				index:     index - 1,
			})
		}
	}

	for len(heap) > 0 && len(res) < 10 {
		node := heap.ExtractMax()
		res = append(res, node.tweetId)
		if node.index >= 0 {
			heap.Insert(Node{
				count:     t.tweetMap[node.foloweeId][node.index].Time,
				tweetId:   t.tweetMap[node.foloweeId][node.index].Id,
				foloweeId: node.foloweeId,
				index:     node.index - 1,
			})
		}
	}

	return res
}

func (t *Twitter) Follow(followerId int, followeeId int) {
	if _, ok := t.followMap[followerId]; !ok {
		t.followMap[followerId] = make(map[int]struct{}, NUMBER_OF_USERS)
	}
	t.followMap[followerId][followeeId] = struct{}{}
}

func (t *Twitter) Unfollow(followerId int, followeeId int) {
	delete(t.followMap[followerId], followeeId)
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
