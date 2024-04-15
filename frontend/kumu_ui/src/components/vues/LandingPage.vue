<template>
  <div id = "WelcomeDiv">
    <div class="small-fluid-container">
      <form id="WelcomeForm" class="login-center-form" v-on:submit.prevent>
        <h1 class="card-header" style="text-align: center;">Rechnungscockpit</h1>
        <hr/>
        <div class="login-form-group">
          <label for="exampleInputEmail1">Anmeldename</label>
          <input style="font-size: 2vmin;" type="text" class="form-control" v-tooltip="{ value: 'Email oder Anmeldename eingeben', showDelay: 1000, hideDelay: 300}" aria-label="Anmeldename" v-model="username" placeholder="Anmeldename">
        </div>
        <div class="login-form-group">
          <label for="exampleInputPassword1">Passwort</label>
          <input style="font-size: 2vmin;" type="password" class="form-control" v-tooltip="{ value: 'Passwort eingeben', showDelay: 1000, hideDelay: 300}"  v-model="password" id="exampleInputPassword1" placeholder="Passwort">
          <small id="Anmeldehilfe" class="form-text text-muted"></small>
        </div>
        <div class="login-form-group">
          <input type="checkbox" class="form-check-input" id="exampleCheck1">
          <label class="form-check-label" for="exampleCheck1" style="font-size: 2vmin;">Eingeloggt bleiben</label>
        </div>
        <div class="login-form-group ">
          <button v-if="this.validateInput() == true" v-on:click="dummy1" class="btn btn-dark btn-Login" type="Submit" id="new-game" aria-expanded="false">Anmelden</button>
          <button v-else v-on:click="dummy1" v-tooltip="{ value: 'Bitte Anmeldename und Passwort eingeben',pt: {text: 'text-danger'}}" class="btn btn-dark btn-Login" type="Submit" id="new-game" aria-expanded="false">Anmelden</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script>
//import SocketioService from '../../services/socketio.service.js';

export default {
  name: 'LandingPage',

  created(){
    /*if(SocketioService.socket.connected == false){
      console.log("Hello there!");
      SocketioService.connect();
    }*/
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
    }

  }
}
</script>

<style scoped>
 
</style>