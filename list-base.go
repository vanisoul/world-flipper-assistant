package main

func clickBase(strs []string, str string, funcs []func(x int, y int)) (resStrs []string, resFuncs []func(x int, y int)) {
	strs = append(strs, getSystemImg(str))
	funcs = append(funcs, func(x, y int) {
		mouseClick(x, y)
	})
	resStrs = strs
	resFuncs = funcs
	return
}
