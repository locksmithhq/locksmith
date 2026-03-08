const fetchSessionsByProjectIDUseCase =
  (fetchSessionsByProjectIDRepository) => async (state) => {
    const projectID = state.route.params.id
    const params = {
      page: state.session.filter.page,
      limit: state.session.filter.limit,
      search: state.session.filter.search,
    }
    state.sessions = await fetchSessionsByProjectIDRepository(projectID, params)
  }

export { fetchSessionsByProjectIDUseCase }
