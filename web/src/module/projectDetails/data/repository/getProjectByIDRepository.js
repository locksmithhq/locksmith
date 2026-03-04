const getProjectByIDRepository = (axios) => async (id) => {
  const response = await axios.get(`/projects/${id}`)
  return response.data
}

export { getProjectByIDRepository }
