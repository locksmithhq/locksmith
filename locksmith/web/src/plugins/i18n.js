import { createI18n } from 'vue-i18n'
import en from '@/locales/en'
import ptBr from '@/locales/pt-br'

const messages = {
  en,
  'pt-br': ptBr,
}

const i18n = createI18n({
  legacy: false, // you must set `false`, to use Composition API
  locale: 'pt-br', // set locale
  fallbackLocale: 'en', // set fallback locale
  messages, // set locale messages
})

export default i18n
