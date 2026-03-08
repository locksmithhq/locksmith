const createAccountRepository = (axios) => async (projectId, client) => {
  const response = await axios.post(`/projects/${projectId}/accounts`, client)
  return response.data
}

export { createAccountRepository }
