const updateOAuthClientRepository =
  (axios) => async (projectId, clientId, client) => {
    const response = await axios.put(
      `/projects/${projectId}/clients/${clientId}`,
      client,
    )
    return response.data
  }

export { updateOAuthClientRepository }
