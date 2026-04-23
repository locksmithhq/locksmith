const getSocialProvidersUseCase = (getSocialProvidersRepository) => async (state) => {
    const projectId = state.route.params.project_id
    const clientId = state.route.params.client_id
    const list = await getSocialProvidersRepository(projectId, clientId)

    state.socialProviders = { google: { enabled: false, client_key: '', client_secret: '' } }

    if (Array.isArray(list)) {
        list.forEach((p) => {
            state.socialProviders[p.provider] = {
                enabled: p.enabled,
                client_key: p.client_key,
                client_secret: p.client_secret,
            }
        })
    }
}

export { getSocialProvidersUseCase }
