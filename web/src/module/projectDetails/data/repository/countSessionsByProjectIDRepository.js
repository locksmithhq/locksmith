const countSessionsByProjectIDRepository = (axios) => async (projectID, search) => {
  const response = await axios.get(`/projects/${projectID}/sessions/count`, {
    params: { search },
  })
  return response.data
}

export { countSessionsByProjectIDRepository }
