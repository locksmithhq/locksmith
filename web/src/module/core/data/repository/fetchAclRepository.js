const fetchAclRepository = (axios) => async (params) => {
  const response = await axios.get('/acl', params.buildAxiosConfig())
  return response.data
}

export { fetchAclRepository }
