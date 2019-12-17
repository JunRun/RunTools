/**
 *
 * @Description:
 * @Version: 1.0.0
 * @Date: 2019-12-13 14:52
 */
package test

type Option struct {
	num  int
	name string
}

type ModOption func(option *Option)
