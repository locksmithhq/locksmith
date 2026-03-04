const fetchActionsRepository = (axios) => async () => {
  const response = await axios.get('/acl/actions')
  return response.data
}

export { fetchActionsRepository }
