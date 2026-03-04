const getClientByIDRepository = (axios) => async (projectId, clientId) => {
    const response = await axios.get(`/projects/${projectId}/clients/${clientId}`)
    return response.data
}

export { getClientByIDRepository }
