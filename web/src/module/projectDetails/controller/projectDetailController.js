import { onMounted, reactive } from 'vue'
import { useRoute } from 'vue-router'
import { accountSchema } from '../const/accountSchema'
import { oauthClientSchema } from '../const/oauthClientSchema'
import { roleSchema } from '../const/roleSchema'
import { moduleSchema } from '../const/moduleSchema'
import { actionSchema } from '../const/actionSchema'
import { useAppForm } from '@/module/core/composables/appForm'

const projectDetailController =
  (
    getProjectByIDUseCase,
    getClientsByProjectIDUseCase,
    createOAuthClientUseCase,
    updateOAuthClientUseCase,
    createAccountUseCase,
    updateAccountUseCase,
    fetchAccountsByProjectIDUseCase,
    countAccountsByProjectIDUseCase,
    createRoleUseCase,
    createModuleUseCase,
    createActionUseCase,
    fetchRolesUseCase,
    fetchModulesUseCase,
    fetchActionsUseCase,
    fetchProjectAclUseCase,
    createProjectAclUseCase,
    fetchSessionsByProjectIDUseCase,
    countSessionsByProjectIDUseCase,
  ) =>
  () => {
    const route = useRoute()

    const account = useAppForm(accountSchema, {
      id: null,
      name: '',
      email: '',
      username: '',
      password: '',
      role_name: '',
      must_change_password: false,
    })

    const oauthClient = useAppForm(oauthClientSchema, {
      id: null,
      name: '',
      client_id: null,
      client_secret: null,
      grant_types: ['authorization_code'],
      redirect_uris: '',
      domain: '',
    })

    const role = useAppForm(roleSchema, {
      id: null,
      title: '',
    })

    const module = useAppForm(moduleSchema, {
      id: null,
      title: '',
    })

    const action = useAppForm(actionSchema, {
      id: null,
      title: '',
    })

    const state = reactive({
      route,
      permission: {
        permissions: [],
        save: () => {
          createProjectAclUseCase(state)
        },
      },
      account: {
        form: account.fields,
        errors: account.errors,
        filter: {
          page: 1,
          limit: 10,
          search: '',
          totalPages: 1,
        },
        dialog: false,
        openDialog: () => {
          state.account.dialog = true
        },
        openEditDialog: (a) => {
          state.account.dialog = true
          account.setValues(a)
        },
        cancelDialog: () => {
          state.account.dialog = false
          account.handleReset()
        },
        save: account.handleSubmit(() => {
          account.transaction(
            async () => {
              if (state.account.form.id) {
                await updateAccountUseCase(state)
              } else {
                await createAccountUseCase(state)
              }
              await fetchAccountsByProjectIDUseCase(state)
              state.account.dialog = false
            },
            {
              onSuccess: () => {
                state.account.dialog = false
                account.handleReset()
                state.fetchAccountsByProjectID()
              },
            },
          )
        }),
      },
      oauthClient: {
        form: oauthClient.fields,
        errors: oauthClient.errors,
        dialog: false,

        openDialog: () => {
          state.oauthClient.dialog = true
          oauthClient.handleReset()
        },
        openEditDialog: (oa) => {
          state.oauthClient.dialog = true
          const clientData = {
            ...oa,
            grant_types: oa.grant_types
              ? typeof oa.grant_types === 'string'
                ? oa.grant_types.split(' ')
                : oa.grant_types
              : ['authorization_code'],
          }
          oauthClient.setValues(clientData)
          if (oa.domain) {
            state.fetchRoles(oa.domain)
          }
        },
        cancelDialog: () => {
          state.oauthClient.dialog = false
          oauthClient.handleReset()
        },
        save: oauthClient.handleSubmit(() => {
          oauthClient.transaction(
            async () => {
              const clientData = {
                ...state.oauthClient.form,
                grant_types: Array.isArray(state.oauthClient.form.grant_types)
                  ? state.oauthClient.form.grant_types.join(' ')
                  : state.oauthClient.form.grant_types,
              }
              state.oauthClient.form = clientData

              if (state.oauthClient.form.id) {
                await updateOAuthClientUseCase(state)
              } else {
                await createOAuthClientUseCase(state)
              }
              await getClientsByProjectIDUseCase(state)
              state.oauthClient.dialog = false
            },
            {
              onSuccess: () => {
                state.oauthClient.dialog = false
                oauthClient.handleReset()
                state.getClientsByProjectID()
              },
            },
          )
        }),
      },
      role: {
        roles: [],
        fetchRoles: async () => {
          await fetchRolesUseCase(state)
        },
        form: role.fields,
        errors: role.errors,
        save: role.handleSubmit(() => {
          role.transaction(
            async () => {
              await createRoleUseCase(state)
            },
            {
              onSuccess: () => {
                role.handleReset()
                state.role.fetchRoles()
              },
            },
          )
        }),
      },
      module: {
        form: module.fields,
        errors: module.errors,
        modules: [],
        fetchModules: async () => {
          await fetchModulesUseCase(state)
        },
        save: module.handleSubmit(() => {
          module.transaction(
            async () => {
              await createModuleUseCase(state)
            },
            {
              onSuccess: () => {
                module.handleReset()
                state.module.fetchModules()
              },
            },
          )
        }),
      },
      action: {
        form: action.fields,
        errors: action.errors,
        actions: [],
        fetchActions: async () => {
          await fetchActionsUseCase(state)
        },
        save: action.handleSubmit(() => {
          action.transaction(
            async () => {
              await createActionUseCase(state)
            },
            {
              onSuccess: () => {
                action.handleReset()
                state.action.fetchActions()
              },
            },
          )
        }),
      },
      session: {
        filter: {
          page: 1,
          limit: 20,
          search: '',
          totalPages: 1,
        },
      },
      activeTab: route.query.tab || 'config',
      defaultProject: {},
      project: {},
      clients: [],
      accounts: [],
      sessions: [],
      isEdit: false,
      editProject: () => {
        state.defaultProject = { ...state.project }
        state.isEdit = true
      },
      resetProject: () => {
        state.project = state.defaultProject
        state.isEdit = false
      },
      saveProject: () => {
        state.isEdit = false
      },
      fetchAccountsByProjectID: async () => {
        await Promise.all([
          fetchAccountsByProjectIDUseCase(state),
          countAccountsByProjectIDUseCase(state),
        ])
      },
      fetchSessionsByProjectID: async () => {
        await Promise.all([
          fetchSessionsByProjectIDUseCase(state),
          countSessionsByProjectIDUseCase(state),
        ])
      },
      cancelAccountDialog: () => {
        state.accountDialog = false
        account.handleReset()
      },
    })

    onMounted(async () => {
      await getProjectByIDUseCase(state)
      await getClientsByProjectIDUseCase(state)
      await Promise.all([
        fetchAccountsByProjectIDUseCase(state),
        countAccountsByProjectIDUseCase(state),
      ])
      await fetchRolesUseCase(state)
      await fetchModulesUseCase(state)
      await fetchActionsUseCase(state)
      await fetchProjectAclUseCase(state)
      await Promise.all([
        fetchSessionsByProjectIDUseCase(state),
        countSessionsByProjectIDUseCase(state),
      ])
    })

    return state
  }

export { projectDetailController }
