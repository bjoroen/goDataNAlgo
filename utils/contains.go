package utils

func Contains[T comparable](s []T, value T) bool{

  for _, v := range s {
    if v == value {
      return true
    }
  }

  return false

}
