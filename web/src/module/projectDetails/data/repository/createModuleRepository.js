const createModuleRepository = (axios) => async (module) => {
  await axios.post('/acl/module', module)
}

export { createModuleRepository }
