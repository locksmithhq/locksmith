const countRefreshTokensByAccountIDRepository = (axios) => async (projectID, accountID, params) => {
  const response = await axios.get(`/projects/${projectID}/accounts/${accountID}/refresh-tokens/count`, { params })
  return response.data
}

export { countRefreshTokensByAccountIDRepository }
