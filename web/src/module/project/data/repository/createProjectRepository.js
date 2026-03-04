const createProjectRepository = (axios) => async (project) => {
  const response = await axios.post('/projects', project)
  return response.data
}

export { createProjectRepository }
