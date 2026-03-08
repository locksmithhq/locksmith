const getConfigRepository = (axiosInstance) => async () => {
  const response = await axiosInstance.get('/config')
  return response.data
}

export { getConfigRepository }
