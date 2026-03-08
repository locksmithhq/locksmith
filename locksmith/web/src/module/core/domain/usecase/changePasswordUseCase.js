export const changePasswordUseCase =
  (changePasswordRepository) => async (state) => {
    try {
      state.loading = true
      state.error = null

      if (state.newPassword !== state.confirmPassword) {
        throw new Error('Passwords do not match')
      }

      if (state.newPassword.length < 8) {
        throw new Error('Password must be at least 8 characters')
      }

      await changePasswordRepository({
        jwt: state.changePasswordJWT,
        password: state.newPassword,
        confirmPassword: state.confirmPassword,
        clientID: state.clientId,
      })

      // After success, it should probably clear the flag or redirect
      // But since the loginRepository.js usually redirects,
      // we might need to handle the redirection after password change or just allow them to login again.
      // The user didn't specify what happens AFTER password change.
      // Usually, we redirect back to the authorize flow if it was a successful login.

      // For now, let's just reload to let them login with the new password,
      // or as it's a "must change password" during login, they should be redirected to the original redirect_uri.
      // Wait, the original login request had a redirect_uri.

      window.location.reload()
    } catch (error) {
      state.loading = false
      state.error = error.message || 'Change password failed'
    }
  }
