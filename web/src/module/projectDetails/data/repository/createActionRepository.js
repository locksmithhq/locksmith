const createActionRepository = (axios) => async (action) => {
  await axios.post('/acl/action', action)
}

export { createActionRepository }
