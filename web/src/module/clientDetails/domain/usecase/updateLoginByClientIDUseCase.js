const updateLoginByClientIDUseCase = (updateLoginByClientIDRepository) => (state) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    const { loginConfig } = state
    return updateLoginByClientIDRepository(projectId, clientId, loginConfig)
}

export { updateLoginByClientIDUseCase }
