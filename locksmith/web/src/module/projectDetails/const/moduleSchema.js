import * as z from 'zod'

const moduleSchema = z.object({
  title: z
    .string()
    .startsWith('module:', 'Module must start with "module:"')
    .trim()
    .min(2, 'At least 2 chars.'),
})

export { moduleSchema }
