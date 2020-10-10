package lineofsight

import "errors"

func BrilliantFunction() (*Thing, error) {
	something, err := GetSomething()
	if err != nil {
		return nil, err
	}
	defer something.Close()
	if !something.Ok() {
		return nil, errors.New("something went wrong")
	}
	another, err := something.Else()
	if err != nil {
		return nil, &customErr{err: err, location: "BrilliantFunction"}
	}
	another.Lock()
	defer another.Unlock()
	err = another.Update(1)
	if err != nil {
		return nil, err
	}
	return another.Thing(), nil
}

func NotSoBrilliantFunction() (*Thing, error) {
	something, err := GetSomething()
	if err != nil {
		return nil, err
	}
	defer something.Close()
	if something.Ok() {
		another, err := something.Else()
		if err != nil {
			return nil, &customErr{err: err, location: "NotSoBrilliantFunction"}
		}
		another.Lock()
		defer another.Unlock()
		err = another.Update(1)
		if err == nil {
			return another.Thing(), nil
		}
		return nil, err
	} else {
		return nil, errors.New("something went wrong")
	}
}
