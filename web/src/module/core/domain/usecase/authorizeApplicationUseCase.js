export const authorizeApplicationUseCase =
  (authorizeApplicationRepository) => async (state) => {
    if (!state.clientId) {
      throw new Error('client_id is required')
    }

    if (!state.redirectUri) {
      throw new Error('redirect_uri is required')
    }

    return await authorizeApplicationRepository({
      client_id: state.clientId,
      redirect_uri: state.redirectUri,
      response_type: state.responseType,
      state: state.stateParam,
      code_challenge: state.codeChallenge || '',
      code_challenge_method: state.codeChallengeMethod || '',
    })
  }
