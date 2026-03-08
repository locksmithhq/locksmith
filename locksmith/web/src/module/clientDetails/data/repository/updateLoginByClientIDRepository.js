const updateLoginByClientIDRepository = (axios) => async (projectId, clientId, login) => {
    return await axios.put(`/projects/${projectId}/clients/${clientId}/login`, login)
}

export { updateLoginByClientIDRepository }
