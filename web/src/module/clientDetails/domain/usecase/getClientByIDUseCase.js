const getClientByIDUseCase = (getClientByIDRepository) => async (state) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    state.client = await getClientByIDRepository(projectId, clientId)
    state.defaultClient = { ...state.client }
}

export { getClientByIDUseCase }
