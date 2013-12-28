package main

import "errors"
import "github.com/jameseb7/roguelike/types"
import "github.com/jameseb7/roguelike/actions"


type command int
const(
	NONE command = iota
	MOVE
)

/*
runCommand(n, c, args...) performs the command, c, n times 
with args as arguments. Commands arguments are taken as integers
and are expected to be correct for the given command
*/
func runCommand(n int, c command, args ...interface{}) (quit bool, err error) {
	defer func() {
		//catch panics so all errors are signalled by err
		if e := recover(); e != nil {
			err = e.(error)
		}
	}()
	for i := 0; i < n; i++ {
		switch c {
		case NONE:
			return
		case MOVE:
			if len(args) < 1 {
				err = errors.New("Too few arguments for MOVE, require 1")
				return
			}
			move := new(actions.Move)
			move.Dir = args[0].(types.Direction)
			actionChannel <- move
		} 
	}
	return
}
	
	
