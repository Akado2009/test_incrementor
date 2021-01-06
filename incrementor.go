package incrementor

import (
	"fmt"
	"math"
)

type Incrementor struct {
	Number int32
}

/* 
NewIncrementor возвращает объект со значением num
*/ 
func NewIncrementor(num int32) *Incrementor {
	return &Incrementor{
		Number: num,
	}
}

/**
	* Возвращает текущее число. В самом начале это ноль.
*/
func (i *Incrementor) GetNumber() int32 {
	return i.Number
}

/**
* Увеличивает текущее число на один. После каждого вызова этого
* метода getNumber() будет возвращать число на один больше.
*/
func (i *Incrementor) IncrementNumber() {
	if i.Number < math.MaxInt32 - 1 {
		i.Number++
	} else {
		i.Number = 0
	}
}

/**
   * Устанавливает максимальное значение текущего числа.
   * Когда при вызове incrementNumber() текущее число достигает
   * этого значения, оно обнуляется, т.е. getNumber() начинает
   * снова возвращать ноль, и снова один после следующего
   * вызова incrementNumber() и так далее.
   * По умолчанию максимум -- максимальное значение int.
   * Если при смене максимального значения число начинает
   * превышать максимальное значение, то число надо обнулить.
   * Нельзя позволять установить тут число меньше нуля.
*/
func (i *Incrementor) SetMaximumValue(maxValue ...int32) error {
	if len(maxValue) > 1 {
		return fmt.Errorf("too many arguments, expected 1, got: %d.", len(maxValue))
	}
	if len(maxValue) == 1 {
		i.Number = maxValue[0]
		return nil
	}
	i.Number = math.MaxInt32
	return nil
}
