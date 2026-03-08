const getAccountByProjectIDRepository = (axios) => async (projectID, accountID) => {
  const response = await axios.get(`/projects/${projectID}/accounts/${accountID}`)
  return response.data
}

export { getAccountByProjectIDRepository }
