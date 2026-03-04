const getLoginByClientIDUseCase = (getLoginByClientIDRepository) => async (state) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    state.loginConfig = await getLoginByClientIDRepository(projectId, clientId)
}

export { getLoginByClientIDUseCase }
