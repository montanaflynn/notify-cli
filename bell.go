package main

import "fmt"

type bellNotifier struct{}

func (b *bellNotifier) Notify() error {
	_, err := fmt.Print("\a")
	return err
}
