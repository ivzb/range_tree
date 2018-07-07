package range_tree

import (
	"fmt"
	"log"
	"testing"

	"github.com/ivzb/tick/tick/assert"
)

func Test_RangeTree_Search_InvalidRange(t *testing.T) {
	rt := &Rt{}

	low := int64(5)
	high := int64(5)

	var expectedNodes []interface{}
	expectedErr := fmt.Errorf("low(%d) should be less than high(%d)", low, high)
	actualNodes, actualErr := rt.Search(low, high)

	assert.ErrorEqual(t, expectedErr, actualErr)
	assert.Equal(t, expectedNodes, actualNodes)
}

func Test_RangeTree_SearchRange_1element(t *testing.T) {
	rt := &Rt{}

	low := int64(5)
	high := int64(10)

	rt.Insert(4, 4)
	rt.Insert(7, 7)
	rt.Insert(11, 11)

	expectedNodes := []interface{}{
		7,
	}
	var expectedErr error
	actualNodes, actualErr := rt.Search(low, high)

	assert.ErrorEqual(t, expectedErr, actualErr)
	assert.Equal(t, len(expectedNodes), len(actualNodes))

	for i := 0; i < len(expectedNodes); i++ {
		assert.Equal(t, expectedNodes[i], actualNodes[i].(int))
	}
}

func Test_RangeTree_SearchRange_3elements(t *testing.T) {
	rt := &Rt{}

	low := int64(5)
	high := int64(10)

	rt.Insert(4, 4)
	rt.Insert(5, 5)
	rt.Insert(7, 7)
	rt.Insert(10, 10)
	rt.Insert(11, 11)

	expectedNodes := []interface{}{
		5,
		10,
		7,
	}
	var expectedErr error
	actualNodes, actualErr := rt.Search(low, high)

	log.Println(actualNodes)
	assert.ErrorEqual(t, expectedErr, actualErr)
	assert.Equal(t, len(expectedNodes), len(actualNodes))

	for i := 0; i < len(expectedNodes); i++ {
		assert.Equal(t, expectedNodes[i], actualNodes[i].(int))
	}
}

func Test_RangeTree_SearchRange_3elements_right(t *testing.T) {
	rt := &Rt{}

	low := int64(10)
	high := int64(17)

	rt.Insert(4, 4)
	rt.Insert(5, 5)
	rt.Insert(7, 7)
	rt.Insert(10, 10)
	rt.Insert(11, 11)
	rt.Insert(15, 15)
	rt.Insert(16, 16)
	rt.Insert(17, 17)
	rt.Insert(19, 19)
	rt.Insert(20, 20)
	rt.Insert(23, 23)

	expectedNodes := []interface{}{
		10,
		15,
		11,
		17,
		16,
	}
	var expectedErr error
	actualNodes, actualErr := rt.Search(low, high)

	assert.ErrorEqual(t, expectedErr, actualErr)
	assert.Equal(t, len(expectedNodes), len(actualNodes))

	for i := 0; i < len(expectedNodes); i++ {
		assert.Equal(t, expectedNodes[i], actualNodes[i].(int))
	}
}
