const fetchSessionsByAccountIDRepository = (axios) => async (projectID, accountID, params) => {
  const response = await axios.get(`/projects/${projectID}/accounts/${accountID}/sessions`, { params })
  return response.data
}

export { fetchSessionsByAccountIDRepository }
