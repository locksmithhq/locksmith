const createOAuthClientRepository = (axios) => async (projectId, client) => {
  const response = await axios.post(`/projects/${projectId}/clients`, client)
  return response.data
}

export { createOAuthClientRepository }
