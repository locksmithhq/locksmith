const fetchRefreshTokensByAccountIDRepository = (axios) => async (projectID, accountID, params) => {
  const response = await axios.get(`/projects/${projectID}/accounts/${accountID}/refresh-tokens`, { params })
  return response.data
}

export { fetchRefreshTokensByAccountIDRepository }
