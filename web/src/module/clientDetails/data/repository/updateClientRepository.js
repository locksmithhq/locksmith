const updateClientRepository = (axios) => async (projectId, clientId, data) => {
    const response = await axios.put(
        `/projects/${projectId}/clients/${clientId}`,
        data,
    )
    return response.data
}

export { updateClientRepository }
