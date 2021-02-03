<template>
  <v-app-bar app shrink-on-scroll>
    <v-app-bar-nav-icon @click="toggleDrawer"></v-app-bar-nav-icon>

    <v-toolbar-title>{{ title }}</v-toolbar-title>

    <v-spacer></v-spacer>
    <v-btn icon @click="logout">
      <v-icon large color="orange darken-2">
        mdi-logout-variant
      </v-icon>
    </v-btn>
  </v-app-bar>
</template>

<script lang="ts">
import Vue from 'vue';

export default Vue.extend({
  name: 'NavigationBar',
  props: {
    toggleDrawer: {
      type: Function,
    },
  },
  data: () => ({
    title: 'RobinsCloud AdminUI',
  }),
  methods: {
    logout() {
      this.$store.commit('logoutUser');
      this.$router.push('/');
    },
  },
});
</script>
