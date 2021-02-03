<template>
  <v-navigation-drawer v-model="isOpen" app>
    <!--  -->
  </v-navigation-drawer>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'Drawer',
  props: {
    isOpen: {
      type: Boolean,
      default: false,
    },
  },
});
</script>
