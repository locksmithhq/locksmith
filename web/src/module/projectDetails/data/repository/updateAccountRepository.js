const updateAccountRepository =
  (axios) => async (projectId, accountId, account) => {
    const response = await axios.put(
      `/projects/${projectId}/accounts/${accountId}`,
      account,
    )
    return response.data
  }

export { updateAccountRepository }
