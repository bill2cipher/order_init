package order

import "testing"

func makefunc(s string, t *testing.T) func() {
  return func() {
    t.Log(s)
  }
}

func TestOrder(t *testing.T) {
  RegisteFunc("A", makefunc("A", t), "E", "D", "G", "M")
  RegisteFunc("B", makefunc("B", t), "D", "F")
  RegisteFunc("C", makefunc("C", t), "H")
  RegisteFunc("D", makefunc("D", t), "C", "I", "L")
  RegisteFunc("E", makefunc("E", t), "G", "B", "C", "B")
  RegisteFunc("F", makefunc("F", t))
  RegisteFunc("G", makefunc("G", t), "K")
  RegisteFunc("H", makefunc("H", t), "I", "J")
  RegisteFunc("I", makefunc("I", t), "J")
  RegisteFunc("J", makefunc("J", t), "G")
  RegisteFunc("K", makefunc("K", t))
  RegisteFunc("L", makefunc("L", t), "M", "K")
  RegisteFunc("M", makefunc("M", t))
  if err := ExecFunc(); err != nil {
    t.Error(err.Error())
  }
}
