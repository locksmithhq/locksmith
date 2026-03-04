import { useForm, useField } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/zod'

const uow = {
  WithTransaction: async (fn, options = {}) => {
    const {
      onError = (err) => console.error(err),
      onSuccess = () => {},
      onFinally = () => {},
      veeValidate,
      showSnackbar = null, // Pode ser uma função de callback para notificações
      showAlert = null, // Pode ser uma função de callback para alertas
    } = options

    try {
      await fn()
      onSuccess()
    } catch (error) {
      // Se houver erros de campo vindos do backend e tivermos o setErrors
      if (veeValidate?.setErrors && error?.response?.data?.fields?.length > 0) {
        const formErrors = {}
        error.response.data.fields.forEach((err) => {
          formErrors[err.field] = err.message
        })
        veeValidate.setErrors(formErrors)
      }

      // Exemplo de integração com feedbacks visuais globais
      if (showSnackbar && typeof showSnackbar === 'function') {
        showSnackbar(error.message || 'Ocorreu um erro inesperado')
      }

      onError(error)
    } finally {
      onFinally()
    }
  },
}

/**
 * Hook genérico para qualquer transação (API, processos complexos, etc)
 * Desvinculado de qualquer formulário.
 */
export const useTransaction = () => {
  return (fn, options = {}) => uow.WithTransaction(fn, options)
}

/**
 * Super Hook especializado para formulários
 */
export const useAppForm = (zodSchema, modelInstance) => {
  const { handleSubmit, handleReset, errors, setErrors, setValues } = useForm({
    validationSchema: toTypedSchema(zodSchema),
    initialValues: modelInstance,
  })

  const fields = {}
  const register = (obj, path = '') => {
    Object.keys(obj).forEach((key) => {
      const fieldPath = path ? `${path}.${key}` : key
      if (
        typeof obj[key] === 'object' &&
        obj[key] !== null &&
        !Array.isArray(obj[key])
      ) {
        register(obj[key], fieldPath)
      } else {
        // Retornamos o objeto do useField para permitir o uso de fields.name.value no template
        fields[key] = useField(fieldPath).value
      }
    })
  }
  register(modelInstance)

  const transaction = (fn, options = {}) =>
    uow.WithTransaction(fn, {
      veeValidate: { setErrors },
      ...options,
    })

  return {
    fields,
    errors,
    handleSubmit,
    handleReset,
    transaction,
    setErrors,
    setValues,
  }
}

export { uow }
