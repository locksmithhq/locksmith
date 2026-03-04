const fetchSessionsByProjectIDRepository = (axios) => async (projectID, params) => {
  const response = await axios.get(`/projects/${projectID}/sessions`, { params })
  return response.data
}

export { fetchSessionsByProjectIDRepository }
