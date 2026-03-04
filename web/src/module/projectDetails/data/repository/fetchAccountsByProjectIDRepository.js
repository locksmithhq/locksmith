const fetchAccountsByProjectIDRepository = (axios) => async (projectID, params) => {
    const response = await axios.get(`/projects/${projectID}/accounts`, params.buildAxiosConfig())
    return response.data
}

export { fetchAccountsByProjectIDRepository }