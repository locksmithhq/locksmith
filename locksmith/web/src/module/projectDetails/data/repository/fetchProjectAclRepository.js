const fetchProjectAclRepository = (axios) => async (projectID) => {
  const response = await axios.get(`/acl/projects/${projectID}`)
  return response.data
}

export { fetchProjectAclRepository }
