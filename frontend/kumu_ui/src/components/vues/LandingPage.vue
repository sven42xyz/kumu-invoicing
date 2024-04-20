<template>
  <div id="WelcomeDiv">
    <div class="small-fluid-container">
      <form id="WelcomeForm" class="login-center-form" v-on:submit.prevent>
        <h1 class="card-header" style="text-align: center;">Rechnungscockpit</h1>
        <hr />
        <div class="login-form-group">
          <label for="username">Anmeldename</label>
          <input id="username" style="font-size: 2vmin;" type="text" class="form-control"
            v-tooltip="{ value: 'Email oder Anmeldename eingeben', showDelay: 1000, hideDelay: 300 }"
            aria-label="Anmeldename" v-model="username" placeholder="Anmeldename">
        </div>
        <div class="login-form-group">
          <label for="password">Passwort</label>
          <input id="password" style="font-size: 2vmin;" type="password" class="form-control"
            v-tooltip="{ value: 'Passwort eingeben', showDelay: 1000, hideDelay: 300 }" v-model="password"
            placeholder="Passwort">
          <small id="Anmeldehilfe" class="form-text text-muted"></small>
        </div>
        <div class="login-form-group">
          <input type="checkbox" class="form-check-input" id="check">
          <label class="form-check-label" for="check" style="font-size: 2vmin;">Eingeloggt bleiben</label>
        </div>
        <div class="login-form-group ">
          <button v-if="this.validateInput() == true" v-on:click="loginUser" class="btn btn-dark btn-Login"
            type="Submit" id="new-game" aria-expanded="false">Anmelden</button>
          <button v-else v-on:click="loginUser"
            v-tooltip="{ value: 'Bitte Anmeldename und Passwort eingeben', pt: { text: 'text-danger' } }"
            class="btn btn-dark btn-Login" type="Submit" id="new-game" aria-expanded="false">Anmelden</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
import { registerUser, loginUser } from "../../services/auth";

export default {
  name: 'LandingPage',

  created() {
  },

  data() {
    return {
      username: null,
      password: null,
    };
  },

  methods: {
    validateInput() {
      const regex = /^[a-zA-Z0-9]+$/;

      if (!this.username) {
        //add error handling -> user input needed
        return;
      }

      if (!regex.test(this.username)) {
        return;
      }

      if (!this.password) {
        //add error handling -> user input needed
        return;
      }

      if (!regex.test(this.password)) {
        return;
      }

      return true;
    },

    dummy1() {
      //Login erfolgreich
      this.$router.push('/dashboard');
    },

    async registerUser() {
      try {
        const data = {
          username: this.username,
          password: this.password
        }
        const response = await registerUser(JSON.parse(data));
        if (response.code === 200) {
          console.log("hello :)")
        }
      } catch (error) {
        console.log("Error registering user:", error);
      }
    },

    async loginUser() {
      try {
        const data = {
          username: this.username,
          password: this.password
        }
        const response = await loginUser(data);
        if (response.code === 200) {
          this.$router.push("/dashboard")
        }
      } catch (error) {
        console.log("Error loggin in user:", error);
      }
    }
  }
}
</script>

<style scoped></style>