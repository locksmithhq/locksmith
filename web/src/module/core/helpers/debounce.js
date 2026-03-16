const DEBOUNCE_CANCELLED = Symbol('debounce-cancelled')

const debounce = (func, delay) => {
  let timeout
  let rejectPrev = null
  return (...args) => {
    clearTimeout(timeout)
    if (rejectPrev) rejectPrev(DEBOUNCE_CANCELLED)
    return new Promise((resolve, reject) => {
      rejectPrev = reject
      timeout = setTimeout(() => {
        rejectPrev = null
        resolve(func(...args))
      }, delay)
    })
  }
}

const isDebounceCancel = (e) => e === DEBOUNCE_CANCELLED

export { debounce, isDebounceCancel }
