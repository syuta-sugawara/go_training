package sample

import (
	"testing"

	"github.com/golang/mock/gomock"

	mock "./mock_sample"
)

func TestSample1(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSample := mock.NewMockSample(ctrl)
	mockSample.EXPECT().Method("hoge").Return(10)

	t.Log("result:", mockSample.Method("hoge"))
}
