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
    if (error.response.status === 401) {
      if (tries >= maxTries) {
        return Promise.reject(error)
      }
      tries++
      const response = await axiosInstance.post('/locksmith/r', null, {
        withCredentials: true,
      })
      if (response.status === 200) {
        return axiosInstance(error.config)
      }
    }
    return Promise.reject(error)
  },
)

export { axiosInstance }
