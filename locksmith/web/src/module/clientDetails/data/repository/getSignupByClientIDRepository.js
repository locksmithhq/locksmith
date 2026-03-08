const getSignupByClientIDRepository = (axios) => async (projectId, clientId) => {
    const response = await axios.get(`/projects/${projectId}/clients/${clientId}/signup`)
    return response.data
}

export { getSignupByClientIDRepository }
