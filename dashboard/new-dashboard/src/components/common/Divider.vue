<template>
  <div
    v-if="label == null"
    class="relative flex py-5 items-center"
  >
    <div class="flex-grow border-t border-gray-400" />
    <span class="flex-shrink mx-4 text-gray-400 text-lg">{{ props.title }}</span>
    <div class="flex-grow border-t border-gray-400" />
  </div>
  <!-- https://tailwindui.com/components/application-ui/layout/dividers -->
  <div class="v-else relative py-5">
    <div
      class="absolute inset-0 flex items-center"
      aria-hidden="true"
    >
      <div class="w-full border-t border-gray-300" />
    </div>
    <div class="relative flex justify-center">
      <span class="bg-white px-2 text-sm text-gray-500">{{ props.label }}</span>
    </div>
  </div>
</template>

<script setup lang="ts">
export interface Props {
  title?: string
  label?: string
}

const props = withDefaults(defineProps<Props>(), {
  title: undefined,
  label: undefined,
})
</script>