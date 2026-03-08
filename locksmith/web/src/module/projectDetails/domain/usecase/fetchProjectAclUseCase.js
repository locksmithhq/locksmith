const fetchProjectAclUseCase = (fetchProjectAclRepository) => async (state) => {
  const projectID = state.route.params.id
  const permission = await fetchProjectAclRepository(projectID)
  state.permission.permissions = permission.roles
}

export { fetchProjectAclUseCase }
