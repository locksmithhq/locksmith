const fetchActionsUseCase = (fetchActionsRepository) => async (state) => {
  state.action.actions = await fetchActionsRepository()
}

export { fetchActionsUseCase }
