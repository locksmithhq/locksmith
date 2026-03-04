const updateClientUseCase = (updateClientRepository) => async (state) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    const updatedClient = await updateClientRepository(
        projectId,
        clientId,
        state.client,
    )
    state.client = updatedClient
    state.defaultClient = { ...updatedClient }
}

export { updateClientUseCase }
