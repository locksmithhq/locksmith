import axios from 'axios'
import { createPaginate } from 'js-query-pagination'

const maxTries = 2
let tries = 0

const axiosInstance = axios.create({
  baseURL:
    (import.meta.env.VITE_LOCKSMITH_API_BASE_URL || window.location.origin) +
    '/api',
  timeout: 1000,
  headers: {
    'Content-Type': 'application/json',
  },
  paramsSerializer: (params) => {
    return createPaginate().merge(params).buildQueryString()
  },
})

axiosInstance.interceptors.response.use(
  (response) => response,
  async (error) => {
    const url = error.config?.url || ''
    const isLocksmithEndpoint = url.includes('/locksmith/')

    if (error.response?.status === 401 && !isLocksmithEndpoint) {
      if (tries >= maxTries) {
        tries = 0
        window.location.replace('/')
        return Promise.reject(error)
      }
      tries++
      try {
        const response = await axiosInstance.post('/locksmith/r', null, {
          withCredentials: true,
        })
        if (response.status === 200) {
          tries = 0
          return axiosInstance(error.config)
        }
      } catch {
        tries = 0
        window.location.replace('/')
        return Promise.reject(error)
      }
    }
    return Promise.reject(error)
  },
)

export { axiosInstance }
