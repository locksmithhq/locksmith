const createProjectAclRepository = (axios) => async (projectId, form) => {
  const response = await axios.post(`/acl/projects/${projectId}`, {
    roles: form,
  })
  return response.data
}

export { createProjectAclRepository }
