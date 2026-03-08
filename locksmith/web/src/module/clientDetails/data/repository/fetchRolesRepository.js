const fetchRolesRepository = (axios) => async () => {
  const response = await axios.get('/acl/roles')
  return response.data
}

export { fetchRolesRepository }
