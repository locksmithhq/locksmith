const countAccountsByProjectIDRepository = (axios) => async (projectID, params) => {
  const response = await axios.get(`/projects/${projectID}/accounts/count`, params.buildAxiosConfig())
  return response.data
}

export { countAccountsByProjectIDRepository }
