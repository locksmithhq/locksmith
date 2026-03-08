const getLoginByClientIDRepository = (axios) => async (projectId, clientId) => {
    const response = await axios.get(`/projects/${projectId}/clients/${clientId}/login`)
    return response.data
}

export { getLoginByClientIDRepository }