const upsertSocialProviderRepository = (axios) => async (projectId, clientId, provider, data) => {
    const response = await axios.put(`/projects/${projectId}/clients/${clientId}/social-providers/${provider}`, data)
    return response.data
}

export { upsertSocialProviderRepository }
