const fetchRolesUseCase = (fetchRolesRepository) => async (state) => {
  state.roles = await fetchRolesRepository()
}

export { fetchRolesUseCase }
