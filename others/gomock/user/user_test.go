package user_test

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/sgreben/testing-with-gomock/mocks"
)

func TestUse(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockDoer := mocks.NewMockDoer(mockCtrl)

	mockDoer.EXPECT().DoSomething(123, "Hello GoMock").Return(nil).Times(1)

}
