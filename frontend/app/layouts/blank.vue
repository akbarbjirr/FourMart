<template>
  <div class="blank-layout">
    <slot />

    <div class="md-snackbar-container">
      <div 
        v-for="toast in toasts" 
        :key="toast.id" 
        class="md-snackbar"
        :style="toast.type === 'error' ? 'border-left: 4px solid var(--md-sys-color-error)' : 'border-left: 4px solid var(--md-sys-color-primary)'"
      >
        <span class="md-snackbar-text">{{ toast.message }}</span>
        <span class="md-snackbar-action" @click="removeToast(toast.id)">TUTUP</span>
      </div>
    </div>
  </div>
</template>

<script setup>
import { useToast } from '~/composables/useToast'
const { toasts, removeToast } = useToast()
</script>

<style scoped>
.blank-layout {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--md-sys-color-background);
  padding: 16px;
}
</style>
