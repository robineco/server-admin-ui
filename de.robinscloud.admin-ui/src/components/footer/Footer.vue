<template>
  <v-footer color="grey grey darken-4" padless>
    <v-row justify="center" no-gutters>
      <v-col class="grey grey darken-4 py-4 text-center white--text" cols="12">
        Admin UI for <strong>robinscloud</strong> created by robineco
        <v-btn text icon @click="goToGitHub">
          <v-icon>
            mdi-github
          </v-icon>
        </v-btn>
      </v-col>
    </v-row>
  </v-footer>
</template>

<script>
export default {
  name: 'Footer',
  methods: {
    goToGitHub() {
      window.location.href = 'https://github.com/robineco';
    },
  },
};
</script>
