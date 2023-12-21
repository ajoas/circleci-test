package numbers

func CheckZero(input int) bool {
  isZero := input == 0
  return isZero
}

func CheckLarger(input int) string {
  if input > 8 {
    return "XLarge"
  }
  if input > 6 {
    return "Large"
  }
  if input > 4 {
    return "Medium"
  }
  if input > 2 {
    return "Smedium"
  }
  if CheckZero(input) {
    return "Zero"
  }
  return "Small"
}