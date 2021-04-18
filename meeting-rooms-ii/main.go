package main

import (
	"container/heap"
	"sort"
)

/*
	1. Paraphrase
		Given a set of intervals, what is the greatest number of overlap. These overlaps are the number rooms needed. How many rooms?
	2. Review constraints
		1 <= intervals.length <= 104
		0 <= starti < endi <= 10^6
	3. Initial Test cases
		[[0,30],[5,10],[15,20]] 2
		[[7,10],[2,4]] 1
		[[0,30], [0,30], [0,30]] 3
	4. Approach(es?)
		brute force (n^2)
			go through and compare each meeting to every other meeting to get overlap count and store it in =
			a slice of the same length
			go through the slice and pick the highest number of overlaps is the greats number of rooms
			(could probably improve it by avoiding reverse comparisons? but it still wouldbe n^2 ordered)
			(however the constraints tell us this would be 10^4 * 10^4 look ups max, so 10^8 is comparisons)
		sort all intervals by start time (n log n)
		put meeting rooms into a priority queue based on end times, pick the room with the min end time n * log m
	4. Pseudocode
		hard to do with this type of question
	5. Walkthrough Pseudocode with a good test case.
		hard to do with this type of question


 */
func minMeetingRooms(intervals [][]int) int {
	sort.Sort(Intervals(intervals))
	mq := NewMinEndMeetingQueue()
	// loop through the intervals
	for _, iv := range intervals {
		//is meeting room queue empty? create a new meeting from this interval and add it then continue
		if mq.Len() == 0 || mq.GetMinMeetingEndTime() > iv[0] {
			m := NewMeetingRoom(iv)
			mq.AddMeetingRoom(m)
		} else
		{
			// mq.GetMin().lastEnd =< iv[0] => is the minimum ending less than the start of the next meeting? If so we can schedule it there
			m := mq.RemoveMeetingRoomWithMinEndTime()
			m.AddMeeting(iv)
			mq.AddMeetingRoom(m)
		}
	}
		//get room with min end time, can we schedule interval into the meeting room?
		//	if so then remove meeting room, add meeting then add back into queue
		//	if not then create new meeting from this interval and add it to the queue
	// the length of the meeting room queue is the min number of meetings
	return mq.Len()
}


// type Interface interface {
//    // Len is the number of elements in the collection.
//    Len() int
//
//    // Less reports whether the element with index i
//    // must sort before the element with index j.
//    //
//    // If both Less(i, j) and Less(j, i) are false,
//    // then the elements at index i and j are considered equal.
//    // Sort may place equal elements in any order in the final result,
//    // while Stable preserves the original input order of equal elements.
//    //
//    // Less must describe a transitive ordering:
//    //  - if both Less(i, j) and Less(j, k) are true, then Less(i, k) must be true as well.
//    //  - if both Less(i, j) and Less(j, k) are false, then Less(i, k) must be false as well.
//    //
//    // Note that floating-point comparison (the < operator on float32 or float64 values)
//    // is not a transitive ordering when not-a-number (NaN) values are involved.
//    // See Float64Slice.Less for a correct implementation for floating-point values.
//    Less(i, j int) bool
//
//    // Swap swaps the elements with indexes i and j.
//    Swap(i, j int)
//}

type Intervals [][]int

func (iv Intervals) Len() int {
	return len(iv)
}

func (iv Intervals) Less(i, j int) bool {
	if iv[i][0] < iv[j][0] {
		return true
	}
	return false
}

func (iv Intervals) Swap(i, j int) {
	tmp := iv[i]
	iv[i] = iv[j]
	iv[j] = tmp
}

type MeetingRoom struct  {
	lastEnd int
}

func NewMeetingRoom(interval []int) *MeetingRoom {
	return &MeetingRoom{lastEnd: interval[1]}
}

type MinEndMeetingQueue struct {
	rooms []*MeetingRoom
}

func NewMinEndMeetingQueue() *MinEndMeetingQueue {
	return &MinEndMeetingQueue{}
}

func (m *MeetingRoom) AddMeeting(interval []int) {
	m.lastEnd = interval[1]
}

// type Interface interface {
//    sort.Interface
//    Push(x interface{}) // add x as element Len()
//    Pop() interface{}   // remove and return element Len() - 1.
//}

func (q *MinEndMeetingQueue) Push(x interface{}) {
	q.rooms = append(q.rooms, x.(*MeetingRoom))
}

func (q *MinEndMeetingQueue) Pop() interface{} {
	var last *MeetingRoom
	q.rooms, last = q.rooms[:len(q.rooms)-1], q.rooms[len(q.rooms)-1]
	return last
}

func (q *MinEndMeetingQueue) Len() int {
	return len(q.rooms)
}

func (q *MinEndMeetingQueue) Less(i, j int) bool {
	if q.rooms[i].lastEnd < q.rooms[j].lastEnd {
		return true
	}
	return false
}

func (q *MinEndMeetingQueue) Swap(i, j int) {
	tmp := q.rooms[i]
	q.rooms[i] = q.rooms[j]
	q.rooms[j] = tmp
}

func (q *MinEndMeetingQueue) GetMinMeetingEndTime() int {
	return q.rooms[0].lastEnd
}

func (q *MinEndMeetingQueue) AddMeetingRoom(m *MeetingRoom)  {
	heap.Push(q, m)
}

func (q *MinEndMeetingQueue) RemoveMeetingRoomWithMinEndTime() *MeetingRoom {
	return heap.Pop(q).(*MeetingRoom)
}