package main

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestBreed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	duck := NewMockDuck(ctrl)
	quack := "quak quak"
	duck.EXPECT().Quack().Return(quack)

	farmer := NewFarmer(duck)
	fmt.Println(farmer)
	breed := farmer.Breed()

	if fmt.Sprintf("#0 %s\n", quack) != breed {
		t.Error("breeding is failed [" + breed + "]")
	}
}
