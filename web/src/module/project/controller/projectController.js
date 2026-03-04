import { useAppForm } from '@/module/core/composables/appForm'
import { onMounted, reactive } from 'vue'
import * as z from 'zod'

const projectController =
  (
    fetchProjectsUseCase,
    createProjectUseCase,
    updateProjectUseCase,
    deleteProjectUseCase,
  ) => () => {
    const schema = z.object({
      name: z.string().trim().min(2, 'At least 2 chars.'),
      domain: z
        .string()
        .trim()
        .startsWith('domain:', 'Must start with domain:'),
    })

    const {
      fields,
      errors,
      handleSubmit,
      handleReset,
      transaction,
      setValues,
    } = useAppForm(schema, {
      id: null,
      name: '',
      description: '',
      domain: '',
    })

    const state = reactive({
      errors,
      projects: [],
      dialog: false,
      form: fields,
      filter: {
        page: 1,
        limit: 25,
        projectName: '',
        projectDescription: '',
      },
      openDialog: () => {
        state.dialog = true
        handleReset()
      },
      openEditDialog: async (p) => {
        setValues(p)
        state.dialog = true
      },
      applyFilter: async () => {
        await fetchProjectsUseCase(state)
      },
      save: handleSubmit(() => {
        transaction(
          async () => {
            if (state.form.id) {
              await updateProjectUseCase(state)
            } else {
              await createProjectUseCase(state)
            }
          },
          {
            onSuccess: () => {
              state.dialog = false
              handleReset()
              state.applyFilter()
            },
          },
        )
      }),
      deleteDialog: false,
      deleteId: null,
      openDeleteDialog: (project) => {
        state.deleteId = project.id
        state.deleteDialog = true
      },
      confirmDelete: async () => {
        await deleteProjectUseCase(state.deleteId)
        state.deleteDialog = false
        state.deleteId = null
        state.applyFilter()
      },
    })

    onMounted(async () => {
      await fetchProjectsUseCase(state)
    })

    return state
  }

export { projectController }
