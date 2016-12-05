package order

import (
  "errors"
)

type orderInfo struct {
  name string
  deps map[string]bool
  exec func()
}

var orders map[string]*orderInfo

func init() {
  orders = make(map[string]*orderInfo)
}

func RegisteFunc(name string, exec func(), deps ...string) {
  info := &orderInfo{name : name, exec: exec}
  info.deps = make(map[string]bool)
  for _, d := range deps {
    info.deps[d] = true
  }
  orders[name] = info
}

func ExecFunc() error {
  fs, err := sort()
  if err != nil {
    return err
  }

  for i := len(fs) - 1; i >= 0; i-- {
    fs[i]()
  }
  return nil
}

func sort() ([]func(), error) {
  var rslt []func()
  for {
    if k, f := next(orders); k == "" && len(orders) != 0 {
      return nil, errors.New("cycle depends")
    } else if k == "" {
      return rslt, nil
    } else {
      delete(orders, k)
      rslt = append(rslt, f)
    }
  }
}

func next(deps map[string]*orderInfo) (string, func()) {
  loop:
  for k, ki := range deps {
    for _, i := range deps {
      if depend(k, i) {
        continue loop
      }
    }
    return k, ki.exec
  }
  return "", nil
}

func depend(name string, info *orderInfo) bool {
  _, ok := info.deps[name]
  return ok
}