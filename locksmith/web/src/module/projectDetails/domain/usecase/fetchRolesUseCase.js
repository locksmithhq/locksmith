const fetchRolesUseCase = (fetchRolesRepository) => async (state) => {
  state.role.roles = await fetchRolesRepository()
}

export { fetchRolesUseCase }
