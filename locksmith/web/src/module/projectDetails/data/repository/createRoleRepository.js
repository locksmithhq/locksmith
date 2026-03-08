const createRoleRepository = (axios) => async (role) => {
  await axios.post('/acl/role', role)
}

export { createRoleRepository }
