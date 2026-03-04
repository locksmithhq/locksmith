/**
 * plugins/vuetify.js
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import { createVuetify } from 'vuetify'

import colors from 'vuetify/util/colors'

// Themes
const lightTheme = {
  dark: false,
  colors: {
    primary: '#2563eb',
    secondary: '#953DF4',
    error: colors.red.accent3,
  },
}

const darkTheme = {
  dark: true,
  colors: {
    primary: '#2563eb',
    secondary: '#953DF4',
    error: colors.red.accent2,
  },
}

// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides
export default createVuetify({
  theme: {
    defaultTheme: 'light',
    themes: {
      light: lightTheme,
      dark: darkTheme,
    },
  },
  defaults: {
    VTextField: {
      variant: 'outlined',
      color: 'primary',
      rounded: 'lg',
      density: 'compact',
    },
    VTextarea: {
      variant: 'outlined',
      color: 'primary',
      rounded: 'lg',
      density: 'compact',
    },
    VSelect: {
      variant: 'outlined',
      color: 'primary',
      rounded: 'lg',
      density: 'compact',
    },
    VCard: { rounded: 'lg' },
    VBtn: { color: 'primary', rounded: 'lg' },
    VAppBar: { density: 'compact' },
    VTabs: { density: 'compact' },
    VTab: { density: 'compact' },
    VChip: { density: 'compact' },
    VList: { density: 'compact' },
    VListItem: { density: 'compact' },
    VPagination: { density: 'compact' },
  },
})
