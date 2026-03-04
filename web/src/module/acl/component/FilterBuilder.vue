<template>
  <v-card
    elevation="0"
    border
    class="pa-6 bg-white rounded-lg filter-builder-card"
    width="800"
  >
    <div class="d-flex align-center justify-space-between mb-6">
      <h3 class="text-h6 font-weight-bold text-grey-darken-3">
        Include filters
      </h3>
      <v-btn
        variant="text"
        size="small"
        color="grey-darken-1"
        prepend-icon="mdi-trash-can-outline"
        @click="$emit('clear')"
        class="text-none"
      >
        Clear all
      </v-btn>
    </div>

    <div
      v-for="(rule, index) in rules"
      :key="index"
      class="d-flex align-center gap-4 mb-4"
    >
      <!-- Prefix Label (Where / AND / OR) -->
      <div class="prefix-container">
        <span
          v-if="index === 0"
          class="text-body-2 font-weight-medium text-grey-darken-1"
          >Where</span
        >
        <v-select
          v-else
          :model-value="logic"
          @update:model-value="$emit('update:logic', $event)"
          :items="['AND', 'OR']"
          density="compact"
          variant="outlined"
          hide-details
          class="logic-select"
        ></v-select>
      </div>

      <!-- Column Selector -->
      <v-select
        v-model="rule.column"
        :items="columns"
        density="compact"
        variant="outlined"
        hide-details
        class="flex-grow-1"
        :disabled="rule.condition === 'search'"
      ></v-select>

      <!-- Condition Selector -->
      <v-select
        v-model="rule.condition"
        :items="conditions"
        density="compact"
        variant="outlined"
        hide-details
        class="flex-grow-1"
        @update:model-value="handleConditionChange(rule)"
      ></v-select>

      <!-- Value Input(s) -->
      <div v-if="rule.condition === 'between'" class="d-flex gap-2 flex-grow-1">
        <v-text-field
          v-model="rule.value[0]"
          placeholder="From"
          density="compact"
          variant="outlined"
          hide-details
          @keyup.enter="$emit('apply')"
        ></v-text-field>
        <v-text-field
          v-model="rule.value[1]"
          placeholder="To"
          density="compact"
          variant="outlined"
          hide-details
          @keyup.enter="$emit('apply')"
        ></v-text-field>
      </div>
      <v-text-field
        v-else
        v-model="rule.value"
        :placeholder="
          rule.condition === 'whereIn'
            ? 'Val1, Val2...'
            : rule.condition === 'search'
            ? 'Global search...'
            : 'Value'
        "
        density="compact"
        variant="outlined"
        hide-details
        class="flex-grow-1"
        @keyup.enter="$emit('apply')"
      ></v-text-field>

      <!-- Remove Button -->
      <v-btn
        icon="mdi-trash-can-outline"
        variant="text"
        size="small"
        color="grey-lighten-1"
        @click="removeRule(index)"
      ></v-btn>
    </div>

    <div class="d-flex align-center gap-4 mt-6">
      <v-btn
        variant="text"
        color="primary"
        prepend-icon="mdi-plus"
        @click="addRule"
        class="text-none font-weight-bold"
      >
        Add filter
      </v-btn>
      <v-spacer></v-spacer>

      <v-btn
        color="primary"
        variant="flat"
        @click="$emit('apply')"
        class="text-none px-6"
        rounded="lg"
      >
        Apply filters
      </v-btn>
    </div>
  </v-card>
</template>

<script setup>
import { ref } from 'vue'

const props = defineProps({
  rules: {
    type: Array,
    required: true,
  },
  columns: {
    type: Array,
    required: true,
  },
  logic: {
    type: String,
    default: 'AND',
  },
  autoSort: {
    type: Boolean,
    default: true,
  },
})

const emit = defineEmits([
  'update:rules',
  'apply',
  'update:logic',
  'clear',
  'update:autoSort',
])

const conditions = [
  { title: 'Search (Table Wide)', value: 'search' },
  { title: 'Contains (AND)', value: 'likeAnd' },
  { title: 'Contains (OR)', value: 'likeOr' },
  { title: 'Equal', value: 'equals' },
  { title: 'Where In', value: 'whereIn' },
  { title: 'Greater Than', value: 'greaterThan' },
  { title: 'Less Than or Equal', value: 'lessThanOrEqual' },
  { title: 'Between', value: 'between' },
]

const handleConditionChange = (rule) => {
  if (rule.condition === 'between') {
    rule.value = ['', '']
  } else {
    rule.value = ''
  }
}

const addRule = () => {
  const newRules = [
    ...props.rules,
    { column: props.columns[0]?.value, condition: 'likeAnd', value: '' },
  ]
  emit('update:rules', newRules)
}

const removeRule = (index) => {
  if (props.rules.length === 1) {
    const newRules = [
      { column: props.columns[0]?.value, condition: 'likeAnd', value: '' },
    ]
    emit('update:rules', newRules)
    return
  }
  const newRules = props.rules.filter((_, i) => i !== index)
  emit('update:rules', newRules)
}
</script>

<style scoped>
.filter-builder-card {
  box-shadow: 0 4px 24px rgba(0, 0, 0, 0.06) !important;
  border: 1px solid rgba(0, 0, 0, 0.05) !important;
}

.prefix-container {
  width: 100px;
  display: flex;
  justify-content: center;
}

.logic-select :deep(.v-field__input) {
  padding-top: 0;
  padding-bottom: 0;
  min-height: 40px;
  font-weight: bold;
}

.gap-4 {
  gap: 16px;
}

:deep(.v-field--variant-outlined) {
  border-radius: 8px !important;
  background-color: #fcfcfc !important;
  transition: all 0.2s;
}

:deep(.v-field--variant-outlined:hover) {
  background-color: #fff !important;
  border-color: rgba(var(--v-theme-primary), 0.5) !important;
}

:deep(.v-field--focused) {
  background-color: #fff !important;
  box-shadow: 0 0 0 4px rgba(var(--v-theme-primary), 0.1) !important;
}
</style>
