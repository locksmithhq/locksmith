const getClientsByProjectIDUseCase =
  (getClientsByProjectIDRepository) => async (state) => {
    const id = state.route.params.id
    state.clients = await getClientsByProjectIDRepository(id)
  }

export { getClientsByProjectIDUseCase }
