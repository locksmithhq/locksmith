const countSessionsByAccountIDRepository = (axios) => async (projectID, accountID) => {
  const response = await axios.get(`/projects/${projectID}/accounts/${accountID}/sessions/count`)
  return response.data
}

export { countSessionsByAccountIDRepository }
