<template>
  <v-app id="inspire">
    <v-navigation-drawer v-model="drawer" app>
      <!--  -->
    </v-navigation-drawer>
    <v-app-bar app shrink-on-scroll>
      <v-app-bar-nav-icon @click="drawer = !drawer"></v-app-bar-nav-icon>

      <v-toolbar-title>RobinsCloud AdminUI</v-toolbar-title>

      <v-spacer></v-spacer>
      <v-btn icon @click="logout">
        <v-icon large color="orange darken-2">
          mdi-logout-variant
        </v-icon>
      </v-btn>
    </v-app-bar>

    <v-main v-if="this.$store.state.loggedIn">
      <v-container>
        <h3 class="category">Status</h3>
        <v-row class="crow">
          <v-col cols="6" md="6" sm="12">
            <v-card outlined class="b-card">
              <v-card-title>Bitwarden</v-card-title>
              <v-card-text>
                Status: <v-badge inline color="green" dot /> {{ services.bitwarden.status.status }}
                <br />
              </v-card-text>
            </v-card>
          </v-col>
          <v-col cols="6" md="6" sm="12">
            <v-card outlined class="b-card">
              <v-card-title>Nextcloud</v-card-title>
              <v-card-text>
                Status:
                <v-badge inline color="red" dot />
                {{ services.nextcloud.status.status }} <br />
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>

        <h3 class="category">Backups</h3>
        <v-row class="crow">
          <v-col cols="6" md="6" sm="12">
            <v-card outlined class="b-card">
              <v-card-title>Bitwarden</v-card-title>
              <v-card-text>
                Last Backup: {{ services.bitwarden.backup.lastBackup }} <br />
                Backup Size {{ services.bitwarden.backup.size }}
              </v-card-text>
            </v-card>
          </v-col>
          <v-col cols="6" md="6" sm="12">
            <v-card outlined class="b-card">
              <v-card-title>Nextcloud</v-card-title>
              <v-card-text>
                Last Backup: {{ services.nextcloud.backup.lastBackup }} <br />
                Backup size: {{ services.nextcloud.backup.size }}
              </v-card-text>
            </v-card>
          </v-col>
        </v-row>
      </v-container>
    </v-main>
    <Footer />
  </v-app>
</template>

<script>
// import { checkStatus } from "../services/StatusService/StatusService";
import Footer from '@/components/footer/Footer.vue';

export default {
  name: 'Dashboard',
  components: { Footer },
  async mounted() {
    // checkStatus(this.services.bitwarden).then(
    //   (status) => (this.services.bitwarden.status.status = status)
    // );
  },
  data: () => ({
    drawer: false,
    isLoggedIn: false,
    services: {
      bitwarden: {
        hostname: 'bw.robinscloud.de',
        backup: {
          lastBackup: '27.12.2020',
          size: '5MB',
        },
        status: {
          status: 'offline',
        },
      },
      nextcloud: {
        hostname: 'nc.robinscloud.de',
        backup: {
          lastBackup: '28.12.2020',
          size: '5GB',
        },
        status: {
          status: 'offline',
        },
      },
    },
  }),
  methods: {
    logout() {
      this.$store.commit('logoutUser');
      this.$router.push('/');
    },
  },
};
</script>

<style scoped>
.crow {
  margin-top: 1rem;
  margin-bottom: 1rem;
}

.b-card {
  height: 200px;
}
</style>
