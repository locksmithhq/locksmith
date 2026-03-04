const updateSignupByClientIDUseCase = (updateSignupByClientIDRepository) => (state) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    const { registerConfig } = state
    return updateSignupByClientIDRepository(projectId, clientId, registerConfig)
}

export { updateSignupByClientIDUseCase }
