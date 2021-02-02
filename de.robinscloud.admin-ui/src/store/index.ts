import Vue from 'vue';
import Vuex from 'vuex';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    loggedIn: false,
  },
  mutations: {
    loginUser(state) {
      state.loggedIn = true;
    },
    logoutUser(state) {
      state.loggedIn = false;
    },
  },
  actions: {},
  modules: {},
});
