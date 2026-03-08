<template>
  <v-layout>
    <v-app-bar app class="rounded-b-lg" style="position: fixed">
      <Logo class="ml-2" @click="drawer = !drawer" style="cursor: pointer" />
      <v-spacer></v-spacer>
      <v-menu>
        <template v-slot:activator="{ props }">
          <v-btn v-bind="props">
            {{ locale }}
          </v-btn>
        </template>
        <v-list>
          <v-list-item @click="switchLocale('en')">
            <v-list-item-title>English</v-list-item-title>
          </v-list-item>
          <v-list-item @click="switchLocale('pt-br')">
            <v-list-item-title>Português</v-list-item-title>
          </v-list-item>
        </v-list>
      </v-menu>

      <v-btn @click="logout">{{ $t('common.logout') }}</v-btn>
    </v-app-bar>

    <v-main>
      <router-view />
    </v-main>
  </v-layout>
</template>

<script setup>
import Logo from '@/module/core/component/logo.vue'
import { ref } from 'vue'
import { useI18n } from 'vue-i18n'
import { useRouter, useRoute } from 'vue-router'

const { locale } = useI18n()
const router = useRouter()
const route = useRoute()
const drawer = ref(false)

const switchLocale = (newLocale) => {
  router.push({
    name: route.name,
    params: { ...route.params, locale: newLocale },
    query: route.query,
    replace: true,
  })
  locale.value = newLocale
}

const logout = async () => {
  await fetch('/api/locksmith/logout', { method: 'POST' })
  router.push({ name: 'login' })
}
</script>
