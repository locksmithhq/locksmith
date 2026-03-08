const debounce = (func, delay) => {
  let timeout
  return (...args) => {
    clearTimeout(timeout)
    return new Promise((resolve) => {
      timeout = setTimeout(() => resolve(func(...args)), delay)
    })
  }
}

export { debounce }
