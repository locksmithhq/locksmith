const createProjectAclUseCase = (repository) => async (state) => {
  const projectID = state.route.params.id
  return repository(projectID, state.permission.permissions)
}

export { createProjectAclUseCase }
