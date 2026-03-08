<template>
  <v-breadcrumbs class="ma-0 pa-0" density="compact" :items="items" rounded>
    <template #item="props">
      <v-chip
        density="compact"
        :to="props.item.to"
        color="primary"
        readonly
        size="small"
        :variant="props.item.disabled ? 'elevated' : 'outlined'"
      >
        <p class="text-caption font-weight-semibold">{{ props.item.title }}</p>
      </v-chip>
    </template>
  </v-breadcrumbs>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const items = ref([])

const generateBreadcrumbs = (currentRoute) => {
  const allSegments = currentRoute.path.split('/').filter((segment) => segment)
  let localePrefix = ''
  let pathSegments = [...allSegments]

  const possibleLocales = ['en', 'pt-br']
  if (allSegments.length > 0 && possibleLocales.includes(allSegments[0])) {
    localePrefix = `/${allSegments[0]}`
    pathSegments.shift()
  }

  // Build a map from param value → human label
  // e.g. project_id → "Project", client_id → "Client", id → resolved from context
  const paramValues = new Set(Object.values(currentRoute.params || {}))
  const paramLabels = {}
  Object.entries(currentRoute.params || {}).forEach(([key, value]) => {
    if (key === 'id') return // resolved from previous segment below
    const label = key
      .replace(/_id$/, '')
      .replace(/_/g, ' ')
      .replace(/\b\w/g, (c) => c.toUpperCase())
    paramLabels[value] = label
  })

  let currentPath = ''
  let prevSegmentLabel = ''

  const breadcrumbItems = pathSegments.map((segment, index) => {
    currentPath += `/${segment}`

    let title
    if (paramValues.has(segment)) {
      // Use explicit label from param name, or fall back to singularising the previous segment
      title =
        paramLabels[segment] ||
        prevSegmentLabel.replace(/s$/i, '').replace(/\b\w/, (c) => c.toUpperCase())
    } else {
      title = segment.charAt(0).toUpperCase() + segment.slice(1)
    }

    prevSegmentLabel = title

    return {
      title,
      disabled: index === pathSegments.length - 1,
      to: localePrefix + currentPath,
    }
  })

  breadcrumbItems.unshift({
    title: 'Home',
    disabled:
      currentRoute.path === localePrefix ||
      currentRoute.path === localePrefix + '/',
    to: localePrefix || '/',
  })

  items.value = breadcrumbItems
}

generateBreadcrumbs(route)

watch(route, (newRoute) => {
  generateBreadcrumbs(newRoute)
})
</script>
