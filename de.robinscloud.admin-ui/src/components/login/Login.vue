<template>
  <v-app>
    <v-main>
      <v-container class="main-content">
        <v-form @submit.prevent.capture="login">
          <v-row>
            <v-col cols="12">
              <h1>Welcome to the Login</h1>
              <p class="secret-message">
                01100010 01111001 00111010 01100111 01101001 01110100 01101000 01110101 01100010
                00101110 01100011 01101111 01101101 00101111 01110010 01101111 01100010 01101001
                01101110 01100101 01100011 01101111
              </p>
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="email"
                label="E-mail"
                :error="inputError.isEmailError"
                required
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-text-field
                v-model="password"
                :append-icon="show ? 'mdi-eye' : 'mdi-eye-off'"
                :type="show ? 'text' : 'password'"
                label="Password"
                :error="inputError.isPasswordError"
                @click:append="show = !show"
                required
              ></v-text-field>
            </v-col>
            <v-col cols="12">
              <v-btn block color="primary" elevation="2" large type="submit">Login</v-btn>
            </v-col>
          </v-row>
        </v-form>
      </v-container>
    </v-main>
    <Footer />
  </v-app>
</template>

<script>
import Footer from '@/components/footer/Footer.vue';

export default {
  name: 'Login',
  components: { Footer },
  data: () => ({
    email: '',
    password: '',
    show: false,
    inputError: {
      isEmailError: false,
      isPasswordError: false,
    },
  }),
  methods: {
    login() {
      this.validateInputs();
      if (this.inputError.isEmailError || this.inputError.isPasswordError) {
        return;
      }
      this.$store.commit('loginUser');
      this.$router.push('/home');
    },
    validateInputs() {
      if (this.email === '') {
        this.inputError.isEmailError = true;
      } else {
        this.inputError.isEmailError = false;
      }
      if (this.password === '') {
        this.inputError.isPasswordError = true;
      } else {
        this.inputError.isPasswordError = false;
      }
    },
  },
};
</script>

<style scoped>
.main-content {
  text-align: center;
  margin-top: 10%;
  width: 30%;
}

.secret-message {
  font-size: 20%;
  font-weight: bold;
}
@media screen and (max-width: 1075px) {
  .secret-message {
    visibility: hidden;
  }
}
</style>
