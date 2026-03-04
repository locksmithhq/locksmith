export const changePasswordRepository =
  (axios) =>
  async ({ jwt, password, confirmPassword, clientID }) => {
    await axios.post(
      '/accounts/change-password',
      {
        password,
        confirm_password: confirmPassword,
        client_id: clientID,
      },
      {
        headers: {
          Authorization: `Bearer ${jwt}`,
        },
      },
    )
  }
