const fetchModulesRepository = (axios) => async () => {
  const response = await axios.get('/acl/modules')
  return response.data
}

export { fetchModulesRepository }
