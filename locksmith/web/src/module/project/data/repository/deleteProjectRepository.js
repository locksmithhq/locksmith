const deleteProjectRepository = (axios) => async (id) => {
  await axios.delete(`/projects/${id}`)
}

export { deleteProjectRepository }
