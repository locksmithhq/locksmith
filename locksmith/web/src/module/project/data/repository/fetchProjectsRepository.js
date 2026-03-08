const fetchProjectsRepository = (axios) => async (params) => {
  const reseponse = await axios.get('/projects', params.buildAxiosConfig())
  return reseponse.data
}

export { fetchProjectsRepository }
