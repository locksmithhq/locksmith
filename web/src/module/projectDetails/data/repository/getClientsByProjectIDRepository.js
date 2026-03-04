const getClientsByProjectIDRepository = (axios) => async (id) => {
  const response = await axios.get(`/projects/${id}/clients`)
  return response.data
}

export { getClientsByProjectIDRepository }
