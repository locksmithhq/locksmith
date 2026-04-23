<template>
  <v-row no-gutters>
    <v-col cols="12" lg="4" md="5" class="pa-0 pr-md-6">
      <div class="mb-6">
        <h3 class="text-h6 font-weight-bold mb-1">
          {{ $t('clientDetails.registerSettings') }}
        </h3>
        <p class="text-caption text-grey-darken-1">
          {{ $t('clientDetails.registerSettingsDescription') }}
        </p>
      </div>

      <v-card elevation="0" border rounded="lg" class="pa-4 mb-4">
        <div class="d-flex align-center justify-space-between">
          <div>
            <div class="text-subtitle-2 font-weight-bold">
              {{ $t('clientDetails.registerEnabled') }}
            </div>
            <div class="text-caption text-grey-darken-1">
              {{ $t('clientDetails.registerEnabledDescription') }}
            </div>
          </div>
          <v-switch
            v-model="controller.registerConfig.enabled"
            color="secondary"
            density="compact"
            hide-details
            inset
          ></v-switch>
        </div>
      </v-card>

      <v-expansion-panels
        v-model="activePanels"
        multiple
        class="custom-panels mb-4"
      >
        <!-- Visual & Layout -->
        <v-expansion-panel elevation="0" border rounded="lg" class="mb-3">
          <v-expansion-panel-title class="py-3">
            <v-icon color="secondary" class="mr-3"
              >mdi-view-quilt-outline</v-icon
            >
            <span class="font-weight-bold">{{
              $t('clientDetails.layoutAndStyle')
            }}</span>
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-row dense>
              <v-col cols="12">
                <v-select
                  v-model="controller.registerConfig.layout"
                  :label="$t('clientDetails.pageLayout')"
                  :items="[
                    { title: 'Split (Right)', value: 'split-right' },
                    { title: 'Split (Left)', value: 'split-left' },
                    { title: 'Centered', value: 'centered' },
                  ]"
                  variant="outlined"
                  density="comfortable"
                  hide-details
                  class="mb-4"
                ></v-select>
              </v-col>
              <v-col cols="12">
                <v-select
                  v-model="controller.registerConfig.input_variant"
                  :label="$t('clientDetails.inputStyle')"
                  :items="[
                    { title: 'Outlined', value: 'outlined' },
                    { title: 'Filled', value: 'filled' },
                    { title: 'Underlined', value: 'underlined' },
                  ]"
                  variant="outlined"
                  density="comfortable"
                  hide-details
                ></v-select>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>

        <!-- Branding -->
        <v-expansion-panel elevation="0" border rounded="lg" class="mb-3">
          <v-expansion-panel-title class="py-3">
            <v-icon color="secondary" class="mr-3">mdi-palette-outline</v-icon>
            <span class="font-weight-bold">{{
              $t('clientDetails.brandingAndColors')
            }}</span>
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-row dense>
              <v-col cols="12">
                <v-label class="text-caption font-weight-bold text-grey-darken-1 mb-2"
                  >{{ $t('clientDetails.logoUpload') }}</v-label
                >
                <input
                  ref="logoFileInput"
                  type="file"
                  accept="image/*"
                  style="display: none"
                  @change="handleLogoUpload"
                />
                
                <div v-if="!controller.registerConfig.logo_url">
                  <v-btn
                    color="secondary"
                    variant="outlined"
                    block
                    height="48"
                    @click="triggerLogoUpload"
                  >
                    <v-icon start>mdi-upload</v-icon>
                    {{ $t('clientDetails.logoUpload') }}
                  </v-btn>
                  <div class="text-caption text-grey mt-1">
                    {{ $t('clientDetails.logoSizeHint') }}
                  </div>
                </div>
                
                <div v-else>
                  <v-card variant="outlined" rounded="lg" class="pa-3">
                    <div class="d-flex align-center ga-3">
                      <v-img
                        :src="controller.registerConfig.logo_url"
                        max-height="60"
                        max-width="120"
                        contain
                        class="border rounded"
                      />
                      <div class="flex-grow-1">
                        <div class="text-caption text-grey mb-2">
                          {{ $t('clientDetails.logoCurrent') }}
                        </div>
                        <div class="d-flex ga-2">
                          <v-btn
                            size="small"
                            color="secondary"
                            variant="tonal"
                            @click="triggerLogoUpload"
                          >
                            <v-icon start size="16">mdi-upload</v-icon>
                            {{ $t('clientDetails.logoChange') }}
                          </v-btn>
                          <v-btn
                            size="small"
                            color="error"
                            variant="text"
                            @click="controller.registerConfig.logo_url = ''"
                          >
                            <v-icon start size="16">mdi-delete</v-icon>
                            {{ $t('clientDetails.removeLogo') }}
                          </v-btn>
                        </div>
                      </div>
                    </div>
                  </v-card>
                </div>
              </v-col>

              <v-col cols="12">
                <v-label class="text-caption font-weight-bold text-grey-darken-1 mb-2"
                  >Favicon</v-label
                >
                <input
                  ref="faviconFileInput"
                  type="file"
                  accept="image/*"
                  style="display: none"
                  @change="handleFaviconUpload"
                />
                <div v-if="!controller.registerConfig.favicon_url">
                  <v-btn
                    color="secondary"
                    variant="outlined"
                    block
                    height="48"
                    @click="triggerFaviconUpload"
                  >
                    <v-icon start>mdi-upload</v-icon>
                    Upload Favicon
                  </v-btn>
                  <div class="text-caption text-grey mt-1">
                    Recomendado: PNG, JPG ou SVG — será salvo em 512x512px
                  </div>
                </div>
                <div v-else>
                  <v-card variant="outlined" rounded="lg" class="pa-3">
                    <div class="d-flex align-center ga-3">
                      <v-img
                        :src="controller.registerConfig.favicon_url"
                        max-height="32"
                        max-width="32"
                        contain
                        class="border rounded"
                      />
                      <div class="flex-grow-1">
                        <div class="text-caption text-grey mb-2">Favicon atual</div>
                        <div class="d-flex ga-2">
                          <v-btn
                            size="small"
                            color="secondary"
                            variant="tonal"
                            @click="triggerFaviconUpload"
                          >
                            <v-icon start size="16">mdi-upload</v-icon>
                            Alterar
                          </v-btn>
                          <v-btn
                            size="small"
                            color="error"
                            variant="text"
                            @click="controller.registerConfig.favicon_url = ''"
                          >
                            <v-icon start size="16">mdi-delete</v-icon>
                            Remover
                          </v-btn>
                        </div>
                      </div>
                    </div>
                  </v-card>
                </div>
              </v-col>

              <v-col cols="12">
                <v-menu :close-on-content-click="false" location="bottom start">
                  <template v-slot:activator="{ props }">
                    <v-text-field
                      v-model="controller.registerConfig.primary_color"
                      :label="$t('clientDetails.primaryColor')"
                      variant="outlined"
                      density="comfortable"
                      hide-details
                      class="mb-4"
                      v-bind="props"
                      readonly
                    >
                      <template v-slot:prepend-inner>
                        <div
                          :style="{
                            backgroundColor:
                              controller.registerConfig.primary_color || '#f5576c',
                            width: '24px',
                            height: '24px',
                            borderRadius: '4px',
                            border: '1px solid rgba(0,0,0,0.1)',
                          }"
                        ></div>
                      </template>
                    </v-text-field>
                  </template>
                  <v-color-picker
                    v-model="controller.registerConfig.primary_color"
                    mode="hex"
                    hide-inputs
                    elevation="4"
                  ></v-color-picker>
                </v-menu>
              </v-col>

              <v-col cols="12">
                <v-select
                  v-model="controller.registerConfig.background_type"
                  :label="$t('clientDetails.backgroundType')"
                  :items="[
                    { title: 'Color', value: 'color' },
                    { title: 'Image', value: 'image' },
                  ]"
                  variant="outlined"
                  density="comfortable"
                  hide-details
                  class="mb-4"
                ></v-select>
              </v-col>

              <v-col
                cols="12"
                v-if="controller.registerConfig.background_type === 'image'"
                key="bg-image-input"
              >
                <v-text-field
                  v-model="controller.registerConfig.background_image"
                  :label="$t('clientDetails.backgroundImageUrl')"
                  variant="outlined"
                  density="comfortable"
                  hide-details
                  prepend-inner-icon="mdi-image"
                ></v-text-field>
              </v-col>

              <v-col cols="12" v-else key="bg-color-input">
                <v-menu :close-on-content-click="false" location="bottom start">
                  <template v-slot:activator="{ props }">
                    <v-text-field
                      v-model="controller.registerConfig.background_color"
                      :label="$t('clientDetails.backgroundColor')"
                      variant="outlined"
                      density="comfortable"
                      hide-details
                      v-bind="props"
                      readonly
                    >
                      <template v-slot:prepend-inner>
                        <div
                          :style="{
                            backgroundColor:
                              controller.registerConfig.background_color ||
                              '#ffffff',
                            width: '24px',
                            height: '24px',
                            borderRadius: '4px',
                            border: '1px solid rgba(0,0,0,0.1)',
                          }"
                        ></div>
                      </template>
                    </v-text-field>
                  </template>
                  <v-color-picker
                    v-model="controller.registerConfig.background_color"
                    mode="hex"
                    hide-inputs
                    elevation="4"
                  ></v-color-picker>
                </v-menu>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>

        <!-- Functionalities -->
        <v-expansion-panel elevation="0" border rounded="lg" class="mb-3">
          <v-expansion-panel-title class="py-3">
            <v-icon color="secondary" class="mr-3">mdi-tune-variant</v-icon>
            <span class="font-weight-bold">{{
              $t('clientDetails.functionalities')
            }}</span>
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-row dense>
              <v-col cols="12">
                <v-select
                  v-model="controller.registerConfig.default_role_name"
                  :items="controller.roles.map((r) => ({ title: r.title, value: r.title }))"
                  :label="$t('clientDetails.defaultRole')"
                  variant="outlined"
                  density="comfortable"
                  hide-details
                  class="mb-4"
                ></v-select>
              </v-col>
              <v-col cols="12">
                <v-switch
                  v-model="controller.registerConfig.show_social"
                  :label="$t('clientDetails.socialRegister')"
                  color="secondary"
                  density="compact"
                  hide-details
                ></v-switch>
              </v-col>
            </v-row>
          </v-expansion-panel-text>
        </v-expansion-panel>

        <!-- Custom HTML/CSS -->
        <v-expansion-panel elevation="0" border rounded="lg">
          <v-expansion-panel-title class="py-3">
            <v-icon color="secondary" class="mr-3">mdi-code-braces</v-icon>
            <span class="font-weight-bold">{{
              $t('clientDetails.advancedHtmlCss')
            }}</span>
          </v-expansion-panel-title>
          <v-expansion-panel-text>
            <v-switch
              v-model="controller.registerConfig.use_custom_html"
              :label="$t('clientDetails.enableCustomHtml')"
              color="secondary"
              density="compact"
              inset
              class="mb-2"
            ></v-switch>

            <template v-if="controller.registerConfig.use_custom_html">
              <v-textarea
                v-model="controller.registerConfig.custom_html"
                :label="$t('clientDetails.htmlCode')"
                variant="outlined"
                density="comfortable"
                rows="6"
                placeholder="<div id='custom-register'>...</div>"
                class="code-textarea mb-3"
              ></v-textarea>

              <v-textarea
                v-model="controller.registerConfig.custom_css"
                :label="$t('clientDetails.cssCode')"
                variant="outlined"
                density="comfortable"
                rows="6"
                placeholder="#custom-register { ... }"
                class="code-textarea"
              ></v-textarea>

              <v-alert
                type="info"
                variant="tonal"
                density="compact"
                class="mt-2 text-caption"
              >
                {{ $t('clientDetails.customHtmlHintRegister') }}
              </v-alert>
            </template>
          </v-expansion-panel-text>
        </v-expansion-panel>
      </v-expansion-panels>

      <v-btn
        color="secondary"
        block
        size="large"
        class="text-none mb-6"
        rounded="lg"
        elevation="2"
        :loading="controller.saving"
        @click="controller.saveRegisterConfig"
      >
        <v-icon start>mdi-content-save-outline</v-icon>
        {{ $t('clientDetails.saveChanges') }}
      </v-btn>

      <v-alert
        v-if="controller.saveError"
        type="error"
        variant="tonal"
        density="compact"
        class="mb-4"
        closable
        @click:close="controller.saveError = null"
      >
        {{ controller.saveError }}
      </v-alert>

      <v-alert
        v-if="controller.saveSuccess"
        type="success"
        variant="tonal"
        density="compact"
        class="mb-4"
        closable
        @click:close="controller.saveSuccess = false"
      >
        {{ $t('clientDetails.savedSuccessfully') }}
      </v-alert>

      <!-- Test Links -->
      <v-card elevation="0" border rounded="lg" class="pa-4 bg-grey-lighten-5">
        <div class="d-flex align-center mb-4">
          <v-icon color="grey-darken-2" class="mr-2">mdi-link-variant</v-icon>
          <span class="text-subtitle-2 font-weight-bold">{{
            $t('clientDetails.testLinksTitle')
          }}</span>
        </div>
        <div v-for="link in testLinks" :key="link.url" class="mb-2">
          <v-btn
            variant="flat"
            color="white"
            size="small"
            block
            border
            class="text-none justify-start"
            :href="link.url"
            target="_blank"
            rounded="md"
          >
            <v-icon start size="16" color="secondary">mdi-open-in-new</v-icon>
            <span class="text-truncate text-grey-darken-3">{{ link.uri }}</span>
          </v-btn>
        </div>
        <div
          v-if="testLinks.length === 0"
          class="text-caption text-grey text-center py-2"
        >
          {{ $t('clientDetails.noRedirectUris') }}
        </div>
      </v-card>
    </v-col>

    <v-col cols="12" lg="8" md="7" class="pa-0">
      <div class="d-flex justify-space-between align-center mb-4">
        <div>
          <h2 class="text-h6 font-weight-bold d-flex align-center">
            {{ $t('clientDetails.registerScreen') }}
            <v-chip
              size="x-small"
              color="secondary"
              variant="tonal"
              class="ml-2"
              >Preview Vivo</v-chip
            >
          </h2>
        </div>
        <div class="d-flex">
          <v-btn-toggle
            v-model="previewMode"
            mandatory
            density="compact"
            color="secondary"
            rounded="lg"
          >
            <v-btn value="desktop" size="small">
              <v-icon>mdi-monitor</v-icon>
            </v-btn>
            <v-btn value="mobile" size="small">
              <v-icon>mdi-cellphone</v-icon>
            </v-btn>
          </v-btn-toggle>
        </div>
      </div>

      <!-- Preview Container -->
      <div
        class="preview-wrapper d-flex align-center justify-center bg-grey-lighten-4 rounded-xl border pa-8 overflow-hidden"
        style="height: 800px; position: relative"
      >
        <!-- Browser Frame Decorator for Desktop -->
        <div
          v-if="previewMode === 'desktop'"
          key="browser-frame"
          class="browser-frame-top d-flex align-center px-4"
        >
          <div class="circle red"></div>
          <div class="circle yellow"></div>
          <div class="circle green"></div>
          <div class="browser-address-bar ml-4 text-caption text-grey-darken-1">
            {{ currentOrigin }}/register...
          </div>
        </div>

        <v-card
          elevation="24"
          :class="['preview-content', previewMode]"
          class="rounded-lg overflow-hidden border-0"
          :style="{
            width: previewMode === 'mobile' ? '375px' : '100%',
            height: previewMode === 'mobile' ? '667px' : '100%',
            transition: 'all 0.4s cubic-bezier(0.4, 0, 0.2, 1)',
          }"
        >
          <!-- Custom HTML Preview -->
          <div
            v-if="controller.registerConfig.use_custom_html"
            key="custom-html-preview"
            class="fill-height overflow-auto"
            style="position: relative"
          >
            <component is="style" v-if="controller.registerConfig.custom_css">
              {{ controller.registerConfig.custom_css }}
            </component>
            <div
              v-html="
                controller.registerConfig.custom_html ||
                '<div class=\'pa-8 text-center text-grey\'>Enter custom HTML to see preview</div>'
              "
            ></div>
          </div>

          <v-row
            v-else
            key="default-preview"
            no-gutters
            class="fill-height"
            :class="{
              'justify-center align-center':
                controller.registerConfig.layout === 'centered',
            }"
            :style="
              controller.registerConfig.layout === 'centered'
                ? backgroundStyle
                : {}
            "
          >
            <!-- Branding Side -->
            <v-col
              v-if="controller.registerConfig.layout !== 'centered'"
              cols="12"
              md="6"
              class="fill-height"
              :class="previewMode === 'mobile' ? 'd-none' : 'd-none d-md-flex'"
              :order="controller.registerConfig.layout === 'split-left' ? 0 : 1"
            >
              <div
                class="bg-gradient-primary d-flex flex-column justify-center align-center pa-12 w-100 fill-height"
                :style="{
                  position: 'relative',
                  overflow: 'hidden',
                  ...backgroundStyle,
                }"
              >
                <div class="decorative-circle circle-1"></div>
                <div class="decorative-circle circle-2"></div>

                <div
                  style="position: relative; z-index: 1; width: 100%"
                  class="text-center text-white d-flex flex-column align-center"
                >
                  <v-img
                    v-if="controller.registerConfig.logo_url"
                    key="branding-logo"
                    :src="controller.registerConfig.logo_url"
                    max-height="160"
                    max-width="280"
                    contain
                    style="width: 100%"
                  />
                  <v-icon
                    v-else
                    key="branding-icon"
                    size="100"
                    color="white"
                    style="opacity: 0.9"
                  >
                    mdi-shield-lock-outline
                  </v-icon>
                </div>
              </div>
            </v-col>

            <!-- Form Side -->
            <v-col
              cols="12"
              :md="
                previewMode === 'mobile'
                  ? 12
                  : controller.registerConfig.layout === 'centered'
                    ? 5
                    : 6
              "
              class="d-flex flex-column justify-center align-center"
              :class="[
                previewMode === 'mobile' ? 'pa-6' : 'pa-8 pa-md-16',
                {
                  'rounded-xl': controller.registerConfig.layout === 'centered',
                },
              ]"
              :style="previewMode === 'mobile' ? { minHeight: '100%' } : {}"
            >
              <div class="w-100" style="max-width: 450px">
                <div
                  class="text-center mb-6"
                  v-if="controller.registerConfig.layout === 'centered'"
                >
                  <img
                    v-if="controller.registerConfig.logo_url"
                    key="centered-logo"
                    :src="controller.registerConfig.logo_url"
                    alt="Logo"
                    style="max-height: 60px; max-width: 200px"
                  />
                  <v-icon
                    v-else
                    key="centered-icon"
                    size="48"
                    :color="controller.registerConfig.primary_color || 'secondary'"
                  >
                    mdi-account-plus-outline
                  </v-icon>
                </div>

                <v-form class="register-form">
                  <v-text-field
                    :label="$t('clientDetails.fullName')"
                    :variant="controller.registerConfig.input_variant"
                    density="comfortable"
                    rounded="lg"
                    prepend-inner-icon="mdi-account-outline"
                    hide-details
                    class="mb-4"
                    placeholder="John Doe"
                    :color="controller.registerConfig.primary_color || 'secondary'"
                  ></v-text-field>

                  <v-text-field
                    :label="$t('clientDetails.email')"
                    :variant="controller.registerConfig.input_variant"
                    density="comfortable"
                    rounded="lg"
                    prepend-inner-icon="mdi-email-outline"
                    hide-details
                    class="mb-4"
                    placeholder="user@example.com"
                    :color="controller.registerConfig.primary_color || 'secondary'"
                  ></v-text-field>

                  <v-text-field
                    :label="$t('clientDetails.password')"
                    :variant="controller.registerConfig.input_variant"
                    density="comfortable"
                    rounded="lg"
                    prepend-inner-icon="mdi-lock-outline"
                    type="password"
                    hide-details
                    class="mb-6"
                    placeholder="••••••••"
                    :color="controller.registerConfig.primary_color || 'secondary'"
                  ></v-text-field>

                  <v-btn
                    :color="controller.registerConfig.primary_color || 'secondary'"
                    size="large"
                    block
                    elevation="2"
                    rounded="lg"
                    class="text-none mb-4"
                  >
                    {{ $t('clientDetails.signUp') }}
                  </v-btn>

                  <template v-if="controller.registerConfig.show_social">
                    <v-divider class="my-6">
                      <span class="text-caption text-grey-darken-1 px-3">
                        {{ $t('clientDetails.orSignUpWith') }}
                      </span>
                    </v-divider>
                    <v-row dense>
                      <v-col cols="6">
                        <v-btn
                          variant="outlined"
                          block
                          rounded="lg"
                          class="text-none px-0"
                          size="small"
                        >
                          <v-icon start size="small">mdi-google</v-icon> Google
                          Register
                        </v-btn>
                      </v-col>
                      <v-col cols="6">
                        <v-btn
                          variant="outlined"
                          block
                          rounded="lg"
                          class="text-none px-0"
                          size="small"
                        >
                          <v-icon start size="small">mdi-github</v-icon> GitHub
                          Register
                        </v-btn>
                      </v-col>
                    </v-row>
                  </template>

                  <div class="text-center mt-6">
                    <span class="text-body-2 text-grey-darken-1">
                      {{ $t('clientDetails.alreadyHaveAccount') }}
                    </span>
                    <a
                      href="#"
                      class="text-primary text-decoration-none text-body-2 font-weight-bold ml-1"
                    >
                      {{ $t('clientDetails.signIn') }}
                    </a>
                  </div>
                </v-form>
              </div>
            </v-col>
          </v-row>
        </v-card>
      </div>
    </v-col>
  </v-row>
</template>

<script setup>
import { computed, ref } from 'vue'
import { useI18n } from 'vue-i18n'

const props = defineProps({
  controller: {
    type: Object,
    required: true,
  },
})

const { locale } = useI18n()
const previewMode = ref('desktop')
const activePanels = ref([0, 1, 2])
const currentOrigin = window.location.origin
const logoFileInput = ref(null)
const faviconFileInput = ref(null)

const triggerLogoUpload = () => {
  logoFileInput.value?.click()
}

const triggerFaviconUpload = () => {
  faviconFileInput.value?.click()
}

const handleFaviconUpload = (event) => {
  const file = event.target.files?.[0]
  if (!file) return

  if (file.size > 5 * 1024 * 1024) {
    alert('Favicon must be less than 5MB')
    return
  }

  const reader = new FileReader()
  reader.onload = (e) => {
    const result = e.target?.result
    if (!result || typeof result !== 'string') return

    const img = new Image()
    img.onload = () => {
      const canvas = document.createElement('canvas')
      canvas.width = 512
      canvas.height = 512
      const ctx = canvas.getContext('2d')
      ctx.drawImage(img, 0, 0, 512, 512)
      props.controller.registerConfig.favicon_url = canvas.toDataURL('image/png')
    }
    img.src = result
  }
  reader.readAsDataURL(file)

  event.target.value = ''
}

const handleLogoUpload = (event) => {
  const file = event.target.files?.[0]
  if (!file) return

  if (file.size > 5 * 1024 * 1024) {
    alert('File size must be less than 5MB')
    return
  }

  const reader = new FileReader()
  reader.onload = (e) => {
    const result = e.target?.result
    if (result && typeof result === 'string') {
      props.controller.registerConfig.logo_url = result
    }
  }
  reader.readAsDataURL(file)
  
  event.target.value = ''
}

const testLinks = computed(() => {
  if (!props.controller.client?.redirect_uris) return []
  return props.controller.client.redirect_uris
    .split(/[\s,]+/)
    .filter((uri) => uri.trim())
    .map((uri) => ({
      uri,
      url: `${window.location.origin}/${locale.value}/register?client_id=${props.controller.client.client_id}&redirect_uri=${uri}`,
    }))
})

// Use branding from register config if available
const backgroundStyle = computed(() => {
  const { background_type, background_color, background_image } =
    props.controller.registerConfig || {}

  if (background_type === 'image' && background_image) {
    return {
      backgroundImage: `url(${background_image})`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
    }
  }

  if (background_color) {
    return {
      background: background_color,
    }
  }

  return {}
})
</script>

<style scoped>
.custom-panels :deep(.v-expansion-panel-text__wrapper) {
  padding: 16px;
}

.custom-panels :deep(.v-expansion-panel-title--active) {
  color: rgb(var(--v-theme-secondary));
}

.preview-wrapper {
  background-image: radial-gradient(#d1d1d1 1px, transparent 1px);
  background-size: 20px 20px;
}

.browser-frame-top {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 40px;
  background: #f1f3f4;
  border-bottom: 1px solid #ddd;
  z-index: 10;
}

.browser-frame-top .circle {
  width: 12px;
  height: 12px;
  border-radius: 50%;
  margin-right: 8px;
}

.circle.red {
  background: #ff5f56;
}
.circle.yellow {
  background: #ffbd2e;
}
.circle.green {
  background: #27c93f;
}

.browser-address-bar {
  background: white;
  height: 24px;
  flex-grow: 1;
  border-radius: 12px;
  display: flex;
  align-items: center;
  padding: 0 12px;
  border: 1px solid #e0e0e0;
}

.preview-content.desktop {
  margin-top: 40px;
  height: calc(100% - 40px) !important;
}

.preview-content.mobile {
  box-shadow:
    0 0 0 12px #333,
    0 0 0 15px #555,
    0 30px 60px rgba(0, 0, 0, 0.3) !important;
}

.register-form {
  width: 100%;
}

.code-textarea :deep(textarea) {
  font-family: 'Fira Code', 'Courier New', monospace;
  font-size: 13px;
  line-height: 1.5;
  background-color: #f8f9fa;
  color: #2c3e50;
}

.bg-gradient-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
}

.decorative-circle {
  position: absolute;
  border-radius: 50%;
}

.circle-1 {
  top: -50px;
  right: -50px;
  width: 300px;
  height: 300px;
  background: rgba(255, 255, 255, 0.1);
}

.circle-2 {
  bottom: -100px;
  left: -100px;
  width: 400px;
  height: 400px;
  background: rgba(255, 255, 255, 0.05);
}
</style>
