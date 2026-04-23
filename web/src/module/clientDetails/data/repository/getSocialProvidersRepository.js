const getSocialProvidersRepository = (axios) => async (projectId, clientId) => {
    const response = await axios.get(`/projects/${projectId}/clients/${clientId}/social-providers`)
    return response.data
}

export { getSocialProvidersRepository }
