import * as z from 'zod'

const actionSchema = z.object({
  title: z
    .string()
    .trim()
    .startsWith('action:', 'Action must start with "action:"')
    .min(2, 'At least 2 chars.'),
})

export { actionSchema }
