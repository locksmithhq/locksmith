const createLoginByClientIDRepository = (axios) => async (projectId, clientId, login) => {
    return await axios.post(`/projects/${projectId}/clients/${clientId}/login`, login)
}

export { createLoginByClientIDRepository }