const getProjectByIDUseCase = (getProjectByIDRepository) => async (state) => {
  const id = state.route.params.id
  state.project = await getProjectByIDRepository(id)
  state.defaultProject = { ...state.project }
}

export { getProjectByIDUseCase }
