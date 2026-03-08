const getSignupByClientIDUseCase = (getSignupByClientIDRepository) => async (state) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    state.registerConfig = await getSignupByClientIDRepository(projectId, clientId)
}

export { getSignupByClientIDUseCase }
