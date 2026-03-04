const updateProjectRepository = (axios) => async (id, project) => {
  const response = await axios.put(`/projects/${id}`, project)
  return response.data
}

export { updateProjectRepository }
